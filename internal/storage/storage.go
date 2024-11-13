package storage

import "github.com/f0xg0sasha/storage/internal/file"

type Storage struct {}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Update(filename string, content []byte) (*file.File, error) {
	return file.NewFile(filename, content)
}
