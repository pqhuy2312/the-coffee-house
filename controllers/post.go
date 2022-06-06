package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gosimple/slug"
	"github.com/pqhuy2312/the-coffee-house/configs"
	"github.com/pqhuy2312/the-coffee-house/exceptions"
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories/repoimpl"
)

func CreatePost(c *fiber.Ctx) error {
	formData := new(forms.FPost)

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

	postRepo := repoimpl.NewPostRepo(db)
	tagRepo := repoimpl.NewTagRepo(db)
	userRepo := repoimpl.NewUserRepo(db)
	tag, err := tagRepo.GetById(formData.TagId)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	if formData.Slug == nil {
		cSlug := slug.Make(formData.Title)
		formData.Slug = &cSlug
	}

	if _, err := postRepo.GetBySlug(*formData.Slug); err == nil {
		return exceptions.SlugExistsDBException(c)
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	author, err := userRepo.GetUserById(int(claims["id"].(float64)))

	post, err := postRepo.Insert(&models.Post{
		Title: formData.Title,
		Slug: *formData.Slug,
		Content: formData.Content,
		Thumbnail: formData.Thumbnail,
		Tag: *tag,
		Author: *author,
	})
	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": post,
		"message": "Create post successfully",
	})
}

func UpdatePost (c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	formData := new(forms.FPost)

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

	postRepo := repoimpl.NewPostRepo(db)
	tagRepo := repoimpl.NewTagRepo(db)

	post, err := postRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	if formData.TagId != post.TagId {
		tag, err := tagRepo.GetById(formData.TagId)
		if err != nil {
			return exceptions.InValidParamsException(c)
		}
		post.Tag = *tag
	}

	if formData.Slug != nil && *formData.Slug != post.Slug {
		if _, err = postRepo.GetBySlug(*formData.Slug); err == nil {
			return exceptions.SlugExistsDBException(c)
		}
		post.Slug = *formData.Slug
	}

	post.Title = formData.Title
	post.Thumbnail = formData.Thumbnail
	post.Content = formData.Content
	updatedPost, err := postRepo.Update(post)
	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": updatedPost,
		"message": "Update post successfully",
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	postRepo := repoimpl.NewPostRepo(db)

	_, err = postRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	
	post, err := postRepo.Delete(id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Delete post successfully",
		"data": post,
	})
}