package route

import (
	"Praktikum/constants"
	"Praktikum/controller"
	m "Praktikum/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// Routing withouth JWT
	e.POST("/users", controller.CreateUserController)
	e.GET("/books", controller.GetBooksAllController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/login", controller.LoginUserController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	// Routing User with JWT
	eJwt.GET("/users", controller.GetUsersAllController)
	eJwt.GET("/users/:id", controller.GetUserController)
	//eJwt.POST("/users", controller.CreateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)

	// Routing Book with JWT
	//eJwt.GET("/books", controller.GetBooksAllController)
	//eJwt.GET("/books/:id", controller.GetBookController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	return e
}
