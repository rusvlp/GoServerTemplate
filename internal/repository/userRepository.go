package repository

import (
	"CustomServerTemplate/internal/dto"
	"database/sql"
)

type UserRepository struct {
	db           *sql.DB
	schemaName   string
	usernameName string
	passwordName string
}

func (ur *UserRepository) Configure(db *DB) {
	ur.db = db.conDB
}

func (ur *UserRepository) Create(user *dto.User) (*dto.User, error) {
	err := ur.db.QueryRow(
		"INSERT INTO "+ur.schemaName+"("+ur.usernameName+", "+ur.passwordName+") VALUES ($1, $2);",
		user.Username,
		user.Password).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Initialize(data DBData) error {
	_, err := ur.db.Exec("CREATE TABLE " + data.TableName + "(    PersonID int,\n    LastName varchar(255),\n    FirstName varchar(255),\n    Address varchar(255),\n    City varchar(255))")

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
