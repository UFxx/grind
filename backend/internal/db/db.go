package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func New(dsn string) (*DB, error) {

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection with dsn: %v, err: %v", dsn, err)
	}

	return &DB{
		Client: client,
	}, nil
}
