package repository

import (
	"CustomServerTemplate/internal/dto"
	"database/sql"
	"github.com/BurntSushi/toml"
	"net/http"
	"os"
)

const userConfig string = "config/userConfig.toml"

type UserRepository struct {
	db           *sql.DB `toml:"database"`
	schemaName   string  `toml:"schema_name"`
	usernameName string  `toml:"username_name"`
	passwordName string  `toml:"password_name"`
}

func (repository *UserRepository) Configure(db *DB) {
	repository.db = db.conDB
	db.UserRepository = repository
	user, _ := os.ReadFile(userConfig)
	userString := string(user)
	println(userString)
	test := &UserRepository{}
	_, err := toml.Decode(userString, test) // TODO: понять почему не работает
	repository.schemaName = "user"
	repository.usernameName = "username"
	repository.passwordName = "password"
	if err != nil {
		return
	}
}

func (repository *UserRepository) Create(userData *dto.User) error {
	repository.db.QueryRow(
		"INSERT INTO " + repository.schemaName + "(" + repository.usernameName + ", " + repository.passwordName + ") VALUES (" +
			userData.Username + ", " + userData.Password + ")")
	return nil
}

func (repository *UserRepository) GetUserData(r *http.Request) *dto.User {
	user := &dto.User{}
	user.Username = r.Form.Get(repository.usernameName)
	user.Password = r.Form.Get(repository.passwordName)
	return user
}
