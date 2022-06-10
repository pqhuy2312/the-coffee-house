package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pqhuy2312/the-coffee-house/configs"
	"github.com/pqhuy2312/the-coffee-house/exceptions"
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/utils"

	"github.com/pqhuy2312/the-coffee-house/repositories/repoimpl"
)

func Register(c *fiber.Ctx) error {
	formData := new(forms.FRegister)

	if err := c.BodyParser(formData); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(formData)
	if err != nil {
		listError := exceptions.Validate(formData)
		return exceptions.ValidationFieldException(listError, c)
	}

	db := configs.GetConnection()

	userRepo := repoimpl.NewUserRepo(db)
	
	result, _ := userRepo.GetUserByEmail(formData.Email)

	if result.UserName != "" {
		return exceptions.EmailExistsDBException(c)
	}
	
	newUser, err := userRepo.Insert(formData)

	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": newUser,
	})
}

func Login(c *fiber.Ctx) error {
	formData := new(forms.FLogin)



	if err := c.BodyParser(formData); err != nil {
		return err
	}

	validate := validator.New()
	
	if err := validate.Struct(formData); err != nil {
		listError := exceptions.Validate(formData)
		return exceptions.ValidationFieldException(listError, c)
	}

	db := configs.GetConnection()

	userRepo := repoimpl.NewUserRepo(db)
	
	result, _ := userRepo.GetUserByEmail(formData.Email)

	if formData.Email != result.Email || !utils.CheckPasswordHash(formData.Password, result.Password) {
		return exceptions.IncorrectEmailPasswordException(c)
	}
	result.TokenVersion += 1
	
	if _, err := userRepo.Update(result); err != nil {
		log.Printf("%v",err)
		return exceptions.InternalServerError(c)
	}
	accessToken := utils.GenerateToken(result, time.Now().Add(time.Minute * 5).Unix())
	refreshToken := utils.GenerateToken(result, time.Now().Add(time.Hour * 12).Unix())
	
	c.Cookie(&fiber.Cookie{
		Name: "refreshToken",
		HTTPOnly: true,
		SameSite: "lax",
		Path: "/api/v1/auth/refresh-token",
		Value: refreshToken,
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"messsage": "Login successfully",
		"data": accessToken,
	})
}

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	db := configs.GetConnection()

	userRepo := repoimpl.NewUserRepo(db)

	result, _ := userRepo.GetUserById(int(claims["id"].(float64)))
	

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	token := c.Cookies("refreshToken")

	if token == "" {
		return exceptions.UnauthorizedException(c)
	}

	claims, _ := utils.ParseToken(token, c)

	db := configs.GetConnection()

	userRepo := repoimpl.NewUserRepo(db)

	user, _ := userRepo.GetUserById(claims.Id)

	if user.UserName == "" || user.TokenVersion != claims.Version {
		return exceptions.UnauthorizedException(c)
	}
	user.TokenVersion += 1
	
	if _, err := userRepo.Update(user); err != nil {
		
		return exceptions.InternalServerError(c)
	}

	accessToken := utils.GenerateToken(user, time.Now().Add(time.Minute * 5).Unix())
	refreshToken := utils.GenerateToken(user, time.Now().Add(time.Hour * 12).Unix())

	c.Cookie(&fiber.Cookie{
		Name: "refreshToken",
		HTTPOnly: true,
		SameSite: "lax",
		Path: "/api/v1/auth/refresh-token",
		Value: refreshToken,
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"messsage": "Refresh token successfully",
		"data": accessToken,
	})
}