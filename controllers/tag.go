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

func CreateTag(c *fiber.Ctx) error {
	formData := new(forms.FTag)

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

	tagRepo := repoimpl.NewTagRepo(db)
	topicRepo := repoimpl.NewTopicRepo(db)

	if _, err = topicRepo.GetById(formData.TopicId); err != nil {
		return exceptions.InValidParamsException(c)
	}

	if formData.Slug == nil {
		slug := slug.Make(formData.Title)
		formData.Slug = &slug
	}

	_, err = tagRepo.GetBySlug(*formData.Slug)
	if err == nil {
		return exceptions.SlugExistsDBException(c)
	}

	tag, err := tagRepo.Insert(formData)
	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": tag,
		"message": "Create Tag successfully",
	})
}

func DeleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	tagRepo := repoimpl.NewTagRepo(db)

	_, err = tagRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	
	tag, err := tagRepo.Delete(id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Delete tag successfully",
		"data": tag,
	})
}

func UpdateTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	TagRepo := repoimpl.NewTagRepo(db)
	topicRepo := repoimpl.NewTopicRepo(db)
	tag, err := TagRepo.GetById(id)

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	formData := new(forms.FTag)
	
	if err := c.BodyParser(formData); err != nil {
		return err
	}

	if formData.TopicId != 0 && tag.TopicId != formData.TopicId {
		if _, err = topicRepo.GetById(formData.TopicId); err != nil {
			return exceptions.InValidParamsException(c)
		}
		tag.TopicId = formData.TopicId
	}
	
	if formData.Slug != nil {
		_, err := TagRepo.GetBySlug(*formData.Slug)
		if err == nil {
			return exceptions.CategoryExistsDBException(c)
		}
		tag.Slug = *formData.Slug
	}
	
	tag.Title = formData.Title

	result, err := TagRepo.Update(tag, id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message": "Update Tag successfully",
		"data": result,
	})
}