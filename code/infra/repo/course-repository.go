package repo

import "app/domain"

var courses []domain.Course

type CourseRepository struct {
	ListCourses   func() []domain.Course
	CreateCourses func(data domain.Course) domain.Course
}

func listCourses() []domain.Course {
	return courses
}

func createCourses(data domain.Course) domain.Course {
	courses = append(courses, data)

	return data
}

func NewCourseRepository() CourseRepository {
	course_repository := CourseRepository{
		listCourses,
		createCourses,
	}

	return course_repository
}
