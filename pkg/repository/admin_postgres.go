package repository

import (
	"HumoAcademy/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) CreateAdmin(admin models.Admin) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO admins (name, username, password_hash) values ($1, $2, $3) RETURNING id")

	row := r.db.QueryRow(query, admin.Name, admin.Username, admin.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AdminPostgres) GetAdmin(username, password string) (models.Admin, error) {
	var user models.Admin
	query := fmt.Sprintf("SELECT id FROM admins WHERE username=$1 AND password_hash=$2")
	err := r.db.Get(&user, query, username, password)
	return user, err
}
