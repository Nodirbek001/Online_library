package postgres

import (
	"sync"

	"gorm.io/gorm"
)

var (
	instanceG *gorm.DB
	once      sync.Once
)
