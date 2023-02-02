package repository

import (
	"CustomServerTemplate/internal/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

type Configurer interface {
	Configure(db *DB)
}

type DB struct {
	conDB          *sql.DB
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

func (db *DB) Open() error {
	con, err := sql.Open("mysql", db.config.DBURL)
	if err != nil {
		return err
	}
	con.SetConnMaxLifetime(time.Minute * 2)
	con.SetMaxOpenConns(10)
	con.SetMaxIdleConns(10)
	if err := con.Ping(); err != nil {
		return err
	}
	db.conDB = con
	return nil
}

func (db *DB) Close() {
	err := db.conDB.Close()
	if err != nil {
		logrus.Error(err)
	}
}
