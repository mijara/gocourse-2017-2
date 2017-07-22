package miniblog

import (
	"time"
)

type Entry struct {
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Body       string    `json:"body"`
	CreateDate time.Time `json:"create_date"`
}
