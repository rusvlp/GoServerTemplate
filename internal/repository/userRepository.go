package repository

import (
	"CustomServerTemplate/internal/dto"
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const userConfig string = "config/userConfig.toml"

type UserRepositorySchema struct {
	SchemaName   string `toml:"schema_name"`
	UsernameName string `toml:"username_name"`
	PasswordName string `toml:"password_name"`
}

type UserRepository struct {
	db               *sql.DB `toml:"database"`
	repositorySchema *UserRepositorySchema
}

func NewUserRepository() *UserRepository {
	fmt.Println("User Repository Constructor is called")
	instance := &UserRepository{}
	userDto := &UserRepositorySchema{}
	_, err := toml.DecodeFile(userConfig, userDto)
	fmt.Println(userDto.SchemaName)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	instance.repositorySchema = userDto

	return instance
}

func (repository *UserRepository) Configure(db *DB) {
	repository.db = db.conDB

}

/*func (repository *UserRepository) constructWithDto(dto *UserRepositorySchema) {
	repository.repositorySchema = dto
} */

func (repository *UserRepository) Create(userData *dto.User) error {
	row, err := repository.db.Exec(
		"INSERT INTO " + repository.repositorySchema.SchemaName + "(" + repository.repositorySchema.UsernameName + ", " + repository.repositorySchema.PasswordName + ") VALUES (" + "'" +
			userData.Username + "', " + "'" + userData.Password + "')")
	if err != nil {
		logrus.Error(err)
	}
	rowA, _ := row.RowsAffected()
	logrus.Info("User with username " + userData.Username + " created." + strconv.Itoa(int(rowA)) + " rows affected.")
	return nil
}

func (repository *UserRepository) GetUserData(r *http.Request) *dto.User {
	user := &dto.User{}
	user.Username = r.Form.Get(repository.repositorySchema.UsernameName)
	user.Password = r.Form.Get(repository.repositorySchema.PasswordName)

	return user
}
