package store

import "minisiga"

// CourseStore describes a storage for courses and operations for them.
type CourseStore interface {
	// GetAll courses
	GetAll() []minisiga.Course

	// Get one course by pk.
	Get(pk int) (minisiga.Course, error)

	// Create a course, reassigning the PK.
	Create(course minisiga.Course)

	// Deletes a course.
	Delete(course minisiga.Course)
}
