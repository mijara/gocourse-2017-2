package store

import (
	"minisiga"
	"errors"
)

type CourseMemoryStore struct {
	courses []minisiga.Course
	lastPk  int
}

func NewCourseDatabase() *CourseMemoryStore {
	return &CourseMemoryStore{
		courses: make([]minisiga.Course, 0),
		lastPk:  0,
	}
}

func (cd *CourseMemoryStore) GetAll() []minisiga.Course {
	return cd.courses
}

func (cd *CourseMemoryStore) Get(pk int) (minisiga.Course, error) {
	// TODO: explain this syntax.
	for _, course := range cd.courses {
		if course.PK == pk {
			return course, nil
		}
	}

	return minisiga.Course{}, errors.New("no course found")
}

func (cd *CourseMemoryStore) Create(course minisiga.Course) {
	course.PK = cd.lastPk
	cd.courses = append(cd.courses, course)
	cd.lastPk++
}

func (cd *CourseMemoryStore) Delete(course minisiga.Course) {
	for i, course := range cd.courses {
		if course.PK == course.PK {
			// TODO: explain this syntax.
			cd.courses = append(cd.courses[:i], cd.courses[i+1:]...)
		}
	}
}
