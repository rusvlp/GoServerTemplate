package repository

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) Configure(db *DB) {
	ur.db = db.ConDB
}

func (ur *UserRepository) Create(data DBData) error {
	_, err := ur.db.Exec("CREATE TABLE " + data.TableName + "(    PersonID int,\n    LastName varchar(255),\n    FirstName varchar(255),\n    Address varchar(255),\n    City varchar(255))")

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
