package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/pqhuy2312/the-coffee-house/configs"
	"github.com/pqhuy2312/the-coffee-house/exceptions"
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/repositories/repoimpl"
)

func CreateTopic(c *fiber.Ctx) error {
	formData := new(forms.FTopic)

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

	topicRepo := repoimpl.NewTopicRepo(db)

	if formData.Slug == nil {
		slug := slug.Make(formData.Title)
		formData.Slug = &slug
	}

	_, err = topicRepo.GetBySlug(*formData.Slug)
	if err == nil {
		return exceptions.SlugExistsDBException(c)
	}

	topic, err := topicRepo.Insert(formData)
	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": topic,
		"message": "Create topic successfully",
	})
}

func DeleteTopic(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	topicRepo := repoimpl.NewTopicRepo(db)

	_, err = topicRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	
	topic, err := topicRepo.Delete(id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Delete topic successfully",
		"data": topic,
	})
}

func UpdateTopic(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	topicRepo := repoimpl.NewTopicRepo(db)
	_, err = topicRepo.GetById(id)

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	formData := new(forms.FTopic)
	
	if err := c.BodyParser(formData); err != nil {
		return err
	}

	
	if formData.Slug != nil {
		_, err := topicRepo.GetBySlug(*formData.Slug)
		if err == nil {
			return exceptions.CategoryExistsDBException(c)
		}
	}
	

	result, err := topicRepo.Update(formData, id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message": "Update topic successfully",
		"data": result,
	})
}