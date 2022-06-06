package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pqhuy2312/the-coffee-house/configs"
	"github.com/pqhuy2312/the-coffee-house/exceptions"
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/repositories/repoimpl"
)

func CreateTopping(c *fiber.Ctx) error {
	formData := new(forms.FTopping)

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

	toppingRepo := repoimpl.NewToppingRepo(db)

	if _, err := toppingRepo.GetByName(formData.Name); err == nil {
		return exceptions.NameExistsDBException(c)
	}

	topping, err := toppingRepo.Insert(formData)

	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Create category successfully",
		"data": topping,
	})
}

func UpdateTopping(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	formData := new(forms.FTopping)

	if err := c.BodyParser(formData); err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(formData)
	if err != nil {
		listError := exceptions.Validate(formData)
		return exceptions.ValidationFieldException(listError, c)
	}

	db := configs.GetConnection()

	toppingRepo := repoimpl.NewToppingRepo(db)

	topping, err := toppingRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	if topping.Name != formData.Name {
		if _, err := toppingRepo.GetByName(formData.Name); err == nil {
			return exceptions.NameExistsDBException(c)
		}
	}
	
	updated, err := toppingRepo.Update(formData, id)

	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message": "Update category successfully",
		"data": updated,
	})
}

func DeleteTopping(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	toppingRepo := repoimpl.NewToppingRepo(db)

	_, err = toppingRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	
	topping, err := toppingRepo.Delete(id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Delete category successfully",
		"data": topping,
	})
}