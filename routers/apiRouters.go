package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pqhuy2312/the-coffee-house/controllers"
	"github.com/pqhuy2312/the-coffee-house/middlewares"
)

func InitializeApiMapping(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/refresh-token", controllers.RefreshToken)

	user := v1.Group("/user")
	user.Get("/me", middlewares.Protected(), controllers.Me)

	category := v1.Group("/category")
	category.Post("/", middlewares.RequiredAdmin(), controllers.Store)
	category.Get("/", controllers.List)
	category.Delete("/:id", middlewares.RequiredAdmin(), controllers.Destroy)
	category.Patch("/:id", middlewares.RequiredAdmin(), controllers.Update)

	topping := v1.Group("/topping", middlewares.RequiredAdmin())
	topping.Post("/", controllers.CreateTopping)
	topping.Patch("/:id", controllers.UpdateTopping)
	topping.Delete("/:id", controllers.DeleteTopping)

	product := v1.Group("/product")
	product.Post("/", middlewares.RequiredAdmin(), controllers.CreateProduct)
	product.Patch("/:id", middlewares.RequiredAdmin(), controllers.UpdateProduct)
	product.Delete("/:id", middlewares.RequiredAdmin(), controllers.DeleteProduct)

	productSize := product.Group("/size", middlewares.RequiredAdmin())
	productSize.Post("/", controllers.CreateProductSize)

	topic := v1.Group("/topic")
	topic.Post("/", middlewares.RequiredAdmin(), controllers.CreateTopic)
	topic.Patch("/:id", middlewares.RequiredAdmin(), controllers.UpdateTopic)
	topic.Delete("/:id", middlewares.RequiredAdmin(), controllers.DeleteTopic)

	tag := v1.Group("/tag")
	tag.Post("/", controllers.CreateTag)
	tag.Patch("/:id", controllers.UpdateTag)
	tag.Delete("/:id", controllers.DeleteTag)

	post := v1.Group("/post")
	post.Post("/", middlewares.RequiredAdmin(), controllers.CreatePost)
	post.Patch("/:id", middlewares.RequiredAdmin(), controllers.UpdatePost)
	post.Delete("/:id", middlewares.RequiredAdmin(), controllers.DeletePost)
}