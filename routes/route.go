package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/controllers"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetUserByID)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)

	app.Get("/users-detail", controllers.GetUserDetailAll)
	app.Get("/users-detail/:id", controllers.GetUserDetailByID)

	app.Get("/lockers", controllers.LockerGetAll)
	app.Post("/lockers", controllers.CreateLocker)
	app.Get("/lockers/:id", controllers.GetLockerByID)
	app.Put("/lockers/:id", controllers.UpdateLocker)
	app.Delete("/lockers/:id", controllers.DeleteLocker)

	app.Get("/posts", controllers.PostGetAll)
	app.Post("/posts", controllers.CreatePost)
	app.Get("/posts/:id", controllers.GetPostById)
	app.Put("/posts/:id", controllers.UpdatePost)

	app.Get("/tags", controllers.TagGetAll)
	app.Post("/tags", controllers.CreateTag)
	app.Get("/tags/:id", controllers.GetTagById)
	app.Put("/tags/:id", controllers.UpdateTag)
	app.Delete("/tags/:id", controllers.DeleteTag)

	app.Get("/tags-detail", controllers.TagGetAllWithPost)
	app.Get("/tags-detail/:id", controllers.GetTagByIdWithPost)
}
