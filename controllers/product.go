package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/pqhuy2312/the-coffee-house/configs"
	"github.com/pqhuy2312/the-coffee-house/exceptions"
	"github.com/pqhuy2312/the-coffee-house/forms"
	"github.com/pqhuy2312/the-coffee-house/models"
	"github.com/pqhuy2312/the-coffee-house/repositories/repoimpl"
)

func CreateProductSize(c *fiber.Ctx) error {
	formData := new(forms.FSize)

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

	sizeRepo := repoimpl.NewSizeRepo(db)

	if _, err = sizeRepo.GetByName(formData.Name); err == nil {
		return exceptions.ProductSizeNameExistsDBException(c)
	}

	newSize, err := sizeRepo.Insert(formData)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": newSize,
		"message": "Create category successfully",
	})
}

func CreateProduct(c *fiber.Ctx) error {
	formData := new(forms.FProduct)

	if err := c.BodyParser(formData); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(formData)
	if err != nil {
		listError := exceptions.Validate(formData)
		return exceptions.ValidationFieldException(listError, c)
	}

	if formData.Slug == nil {
		slug := slug.Make(formData.Name)
		formData.Slug = &slug
	}

	db := configs.GetConnection()

	productRepo := repoimpl.NewProductRepo(db)
	categoryRepo := repoimpl.NewCategoryRepo(db)
	toppingRepo := repoimpl.NewToppingRepo(db)
	
	_, err = categoryRepo.GetCategoryById(formData.CategoryId)
	if err != nil {
		return exceptions.CategoryNotExistsDBException(c)
	}

	_, err = productRepo.GetBySlug(*formData.Slug)

	if err == nil {
		return exceptions.SlugExistsDBException(c)
	}
	
	var images []models.ProductImage
	var sizes []models.Size
	var toppings []models.Topping

	if formData.Images != nil {
		for _, url := range formData.Images {
			images = append(images, models.ProductImage{
				Url: url,
			})
		}
	}

	if formData.Sizes != nil {
		for _, size := range formData.Sizes {
			sizes = append(sizes, models.Size{
				Name: size.Name,
				Price: size.Price,
			})
		}
	}

	if formData.Toppings != nil {
		toppings, err = toppingRepo.GetByIds(formData.Toppings)
		if err != nil {
			return exceptions.InValidParamsException(c)
		}
	}

	product := &models.Product{
		Name: formData.Name,
		Slug: *formData.Slug,
		Info: formData.Info,
		Story: formData.Story,
		Sizes: sizes,
		Images: images,
		Toppings: toppings,
	}

	_, err = productRepo.Insert(product)
	if err != nil {
		return exceptions.InternalServerError(c)
	}
	

	return c.Status(200).JSON(fiber.Map{
		"success":  true,
		"message": "Create product successfully",
		"data": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	productRepo := repoimpl.NewProductRepo(db)
	toppingRepo := repoimpl.NewToppingRepo(db)
	prod, err := productRepo.GetById(id)

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	formData := new(forms.FProduct)
	
	if err := c.BodyParser(formData); err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(formData)
	if err != nil {
		listError := exceptions.Validate(formData)
		return exceptions.ValidationFieldException(listError, c)
	}

	if formData.Slug == nil {
		slug := slug.Make(formData.Name)
		formData.Slug = &slug
	}else if *formData.Slug != prod.Slug {
		if _, err := productRepo.GetBySlug(*formData.Slug); err == nil {
			return exceptions.SlugExistsDBException(c)
		}
	}


	var images []models.ProductImage
	var sizes []models.Size
	var toppings []models.Topping

	if formData.Images != nil {
		for _, url := range formData.Images {
			images = append(images, models.ProductImage{
				Url: url,
			})
		}
	}

	if formData.Sizes != nil {
		for _, size := range formData.Sizes {
			sizes = append(sizes, models.Size{
				Name: size.Name,
				Price: size.Price,
			})
		}
	}

	if formData.Toppings != nil {
		toppings, err = toppingRepo.GetByIds(formData.Toppings)
		if err != nil {
			return exceptions.InValidParamsException(c)
		}
	}

	product := &models.Product{
		Id: id,
		Name: formData.Name,
		Slug: *formData.Slug,
		Info: formData.Info,
		Story: formData.Story,
		Sizes: sizes,
		Images: images,
		Toppings: toppings,
	}

	_, err = productRepo.Update(product)
	if err != nil {
		log.Print(err)
		return exceptions.InternalServerError(c)
	}
	

	return c.Status(200).JSON(fiber.Map{
		"success":  true,
		"message": "Update product successfully",
		"data": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return exceptions.InValidParamsException(c)
	}

	db := configs.GetConnection()

	productRepo := repoimpl.NewProductRepo(db)

	_, err = productRepo.GetById(id)
	if err != nil {
		return exceptions.InValidParamsException(c)
	}
	
	product, err := productRepo.Delete(id)

	if err != nil {
		return exceptions.InternalServerError(c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Delete product successfully",
		"data": product,
	})
}