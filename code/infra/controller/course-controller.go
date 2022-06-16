package controller

import (
	"app/domain"
	"app/infra/repo"
	"net/http"

	"github.com/labstack/echo/v4"
)

var course_repository = repo.NewCourseRepository()

type CourseController struct {
	ListCourses   echo.HandlerFunc
	CreateCourses echo.HandlerFunc
}

func NewCourseController() CourseController {
	controller := CourseController{
		listCourses,
		createCourses,
	}

	return controller
}

func listCourses(c echo.Context) error {
	courses := course_repository.ListCourses()
	return c.JSON(http.StatusOK, courses)
}

func createCourses(c echo.Context) error {
	course := domain.Course{}
	c.Bind(&course)

	c.JSON(http.StatusCreated, course_repository.CreateCourses(course))

	return nil
}
