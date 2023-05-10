package storage

import (
	"Library/storage/postgres"
	"Library/storage/repo"
)

type storagePg struct {
	libraryRepo repo.LibraryStorageI
}

func (s storagePg) User() repo.LibraryStorageI {
	return s.libraryRepo
}

type IStorage interface {
	User() repo.LibraryStorageI
}

func NewStoragePg() IStorage {
	return &storagePg{
		libraryRepo: postgres.NewLibraryRepo(),
	}
}
