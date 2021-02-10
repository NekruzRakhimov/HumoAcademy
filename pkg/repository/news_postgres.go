package repository

import (
	"HumoAcademy/models"
	"github.com/jmoiron/sqlx"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (r *NewsPostgres) GetNewsByID (id int) (models.News, error) {
	var news models.News
	query := "SELECT * FROM news WHERE id = $1"
	err := r.db.Get(&news, query, id)
	return news, err
}