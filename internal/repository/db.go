package repository

import (
	"CustomServerTemplate/internal/config"
	"database/sql"
)

type Configurer interface {
	Configure(db *DB)
}

type DB struct {
	ConDB          *sql.DB
	config         *config.Config
	userRepository *UserRepository
}

func NewDB(config *config.Config) *DB {
	return &DB{
		config: config,
	}
}

func (db *DB) CreateRepositories() {
	sl := []Configurer{
		&UserRepository{},
	}
	for _, j := range sl {
		j.Configure(db)
	}
}
