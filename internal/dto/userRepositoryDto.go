package dto

type UserRepositoryDto struct {
	SchemaName   string `toml:"schema_name"`
	UsernameName string `toml:"username_name"`
	PasswordName string `toml:"password_name"`
}
