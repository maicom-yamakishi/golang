package repository

import (
	"database/sql"
	"peixe-urbano/src/models"
)

// Users represents users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository create repository of users
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inster user in database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("insert into users (name, nick, email, password) values(?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return uint64(lastId), nil
}
