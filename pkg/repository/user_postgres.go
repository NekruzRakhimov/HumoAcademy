package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}

}

func (r *UserPostgres) GetAllSubscribedUsers() ([]string, error) {
	var emails []string
	query := fmt.Sprintf("SELECT email FROM subscribed_users")
	err := r.db.Select(&emails, query)
	if err != nil {
		return []string{}, err
	}
	return emails, err
}


