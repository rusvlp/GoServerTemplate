package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) Configure(db *DB) {
	ur.db = db.conDB
}
