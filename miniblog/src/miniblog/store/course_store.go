package store

import (
	"miniblog"
)

// CourseStore describes a storage for courses and operations for them.
type Store interface {
	// GetAll courses
	GetAll() []miniblog.Entry

	// Get one course by pk.
	Get(pk string) (miniblog.Entry, error)

	// Create a course, reassigning the PK.
	Create(course miniblog.Entry)

	// Deletes a course.
	Delete(course miniblog.Entry)
}
