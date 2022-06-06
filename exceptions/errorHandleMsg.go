package exceptions

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const IN_VALID_PARENT_ID = "Invalid parent id"
const EMAIL_EXIST = "Email exits in DB"
const CATEGORY_EXIST = "Category exits in DB"
const DATABASE_CONNECTION_ERROR = "Database connection error"
const VALIDATION_EXCEPTION = "Validation exception"
const IN_VALID_EMAIL_PASSWORD = "User name or password is incorrect"
const UN_AUTHORIZED = "Un Authorized"
const IN_VALID_TOKEN = "Invalid token"
const INTERNAL_SERVER_ERROR = "Internal server error"
const IN_VALID_PARAMS = "Invalid params"
const PRODUCT_SIZE_NAME_EXIST = "Product size name exists in DB"
const CATEGORY_NOT_EXIST = "Category not exists in DB"
const SLUG_EXIST = "Slug exists in DB"
const Name_EXIST = "Name exists in DB"

func NameExistsDBException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": Name_EXIST,
		"success":  false,
	})

}

func SlugExistsDBException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": SLUG_EXIST,
		"success":  false,
	})

}

func CategoryNotExistsDBException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": CATEGORY_NOT_EXIST,
		"success":  false,
	})

}

func ProductSizeNameExistsDBException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": PRODUCT_SIZE_NAME_EXIST,
		"success":  false,
	})

}

func InValidParamsException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success":  false,
		"message": IN_VALID_PARAMS,
	})

}

func InternalServerError(c *fiber.Ctx) error {

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"success":  false,
		"message": INTERNAL_SERVER_ERROR,
	})

}

func InValidTokenException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success":  false,
		"message": IN_VALID_TOKEN,
	})

}

func UnauthorizedException(c *fiber.Ctx) error {

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"success":  false,
		"message": UN_AUTHORIZED,
	})

}

func IncorrectEmailPasswordException(c *fiber.Ctx) error {

	return c.Status(http.StatusForbidden).JSON(fiber.Map{
		"success":  false,
		"message": IN_VALID_EMAIL_PASSWORD,
	})

}

func EmailExistsDBException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": EMAIL_EXIST,
		"success":  false,
	})

}

func CategoryExistsDBException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": CATEGORY_EXIST,
		"success":  false,
	})

}

func InvalidParentIdException(c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": IN_VALID_PARENT_ID,
		"success":  false,
	})

}

func DatabaseConnectionException(err error, c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success":  false,
		"message": DATABASE_CONNECTION_ERROR,
		"error":   err.Error(),
	})

}

func ValidationFieldException(error []FieldError, c *fiber.Ctx) error {

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success":  false,
		"message": VALIDATION_EXCEPTION,
		"error":   error,
	})

}

func msgForTag(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%v is required", fe.Field())
	case "email":
		return "Invalid email"
	case "gte":
		return fmt.Sprintf("%v must be at least %v characters", fe.Field(), fe.Param())
	case "number":
		return "This field must be a number"
	}
	return fe.Error() // default error
}

func Validate(form interface{}) []FieldError {

	var validate = validator.New()
	err := validate.Struct(form)
	var listError []FieldError
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorMessage FieldError
			errorMessage.Field = err.Field()
			errorMessage.Error = msgForTag(err)
			listError = append(listError, errorMessage)
		}
	}
	return listError
}
