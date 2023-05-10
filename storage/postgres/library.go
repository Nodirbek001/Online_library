package postgres

import (
	"Library/platforma/postgres"
	"Library/storage/repo"
	"context"

	"gorm.io/gorm"
)

type libraryRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *libraryRepo {
	return &libraryRepo{
		db: db,
	}
}

func NewLibraryRepo() repo.LibraryStorageI {
	return &libraryRepo{
		db: postgres.DB(),
	}
}
func (l *libraryRepo) SignUpUser(ctx context.Context)(string, error){
	return "", nil
}
