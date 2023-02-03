package repository

import (
	"CustomServerTemplate/internal/dto"
	"database/sql"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"net/http"
)

const userConfig string = "config/userConfig.toml"

var instance *userRepository

type userRepository struct {
	db           *sql.DB `toml:"database"`
	schemaName   string  `toml:"schema_name"`
	usernameName string  `toml:"username_name"`
	passwordName string  `toml:"password_name"`
}

func NewUserRepository() *userRepository {
	if instance == nil {
		instance = &userRepository{}
		userDto := &dto.UserRepositoryDto{}
		_, err := toml.DecodeFile(userConfig, userDto)
		if err != nil {
			logrus.Error(err)
			return nil
		}
		instance.constructWithDto(userDto)
	}
	return instance
}

func (repository *userRepository) Configure(db *DB) {
	repository.db = db.conDB
}

func (repository *userRepository) constructWithDto(dto *dto.UserRepositoryDto) {
	repository.schemaName = dto.SchemaName
	repository.usernameName = dto.UsernameName
	repository.passwordName = dto.PasswordName
}

func (repository *userRepository) Create(userData *dto.User) error {
	row, err := repository.db.Exec(
		"INSERT INTO " + repository.schemaName + "(" + repository.usernameName + ", " + repository.passwordName + ") VALUES (" +
			userData.Username + ", " + userData.Password + ")")
	if err != nil {
		logrus.Error(err)
	}
	rowA, _ := row.RowsAffected()
	println(rowA)
	return nil
}

func (repository *userRepository) GetUserData(r *http.Request) *dto.User {
	user := &dto.User{}
	user.Username = r.Form.Get(repository.usernameName)
	user.Password = r.Form.Get(repository.passwordName)
	return user
}
