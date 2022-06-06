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

func Store(c *fiber.Ctx) error {
	formData := new(forms.FCategory)

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

	categoryRepo := repoimpl.NewCategoryRepo(db)

	if formData.ParentId != nil {
		_, err = categoryRepo.GetCategoryById(*formData.ParentId)

		if err != nil {
			return exceptions.InvalidParentIdException(c)
		}
	}

	cSlug := slug.Make(*formData.Title)
	formData.Slug = &cSlug
	
	_, err = categoryRepo.GetCategoryBySlug(cSlug)

	if err == nil {
		return exceptions.CategoryExistsDBException(c)
	}

	newCategory, err := categoryRepo.Insert(formData)

	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": newCategory,
		"message": "Create category successfully",
	})
} 

func List(c *fiber.Ctx) error {
	db := configs.GetConnection()

	categoryRepo := repoimpl.NewCategoryRepo(db)

	result, err := categoryRepo.List()
	if err != nil {
		return exceptions.DatabaseConnectionException(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
		"message": "Get categories successfully",
	})
}

func Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	categoryRepo := repoimpl.NewCategoryRepo(db)

	_, err = categoryRepo.GetCategoryById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	
	category, err := categoryRepo.Delete(id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Delete category successfully",
		"data": category,
	})
}

func Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	categoryRepo := repoimpl.NewCategoryRepo(db)
	_, err = categoryRepo.GetCategoryById(id)

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	formData := new(forms.FCategory)
	
	if err := c.BodyParser(formData); err != nil {
		return err
	}

	
	if formData.Slug != nil {
		_, err := categoryRepo.GetCategoryBySlug(*formData.Slug)
		if err == nil {
			return exceptions.CategoryExistsDBException(c)
		}
	}
	
	if formData.ParentId != nil {
		_, err = categoryRepo.GetCategoryById(*formData.ParentId)
		if err != nil {
			return exceptions.InValidParamsException(c)
		}
	}

	result, err := categoryRepo.Update(formData, id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message": "Update category successfully",
		"data": result,
	})
}