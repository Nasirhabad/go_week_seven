package routes

import (
	"week7/controllers"
	"week7/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middlewares.LoggerMiddleware())

	e.POST("/api/v1/register", controllers.Register)
	e.POST("/api/v1/login", controllers.Login)

	r := e.Group("/api/v1/packages")
	r.Use(middlewares.JWTMiddleware())

	r.GET("", controllers.GetPackages)
	r.GET("/:id", controllers.GetPackageByID)
	r.POST("", controllers.CreatePackage)
	r.PUT("/:id", controllers.UpdatePackage)
	r.DELETE("/:id", controllers.DeletePackage)

	return e
}
