package main

import (
	"app/infra/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  	// Echo instance
	e := echo.New()
	course_controller := controller.NewCourseController()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/courses", course_controller.ListCourses)
	e.POST("/courses", course_controller.CreateCourses)

  	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
