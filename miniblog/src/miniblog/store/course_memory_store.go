package store

import (
	"errors"
	"miniblog"
)

type MemoryStore struct {
	entries []miniblog.Entry
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		entries: make([]miniblog.Entry, 0),
	}
}

func (cd *MemoryStore) GetAll() []miniblog.Entry {
	return cd.entries
}

func (cd *MemoryStore) Get(slug string) (miniblog.Entry, error) {
	// TODO: explain this syntax.
	for _, entry := range cd.entries {
		if entry.Slug == slug {
			return entry, nil
		}
	}

	return miniblog.Entry{}, errors.New("no entry found")
}

func (cd *MemoryStore) Create(entry miniblog.Entry) {
	cd.entries = append(cd.entries, entry)
}

func (cd *MemoryStore) Delete(entry miniblog.Entry) {
	for i, entry := range cd.entries {
		if entry.Slug == entry.Slug {
			// TODO: explain this syntax.
			cd.entries = append(cd.entries[:i], cd.entries[i+1:]...)
			return
		}
	}
}
