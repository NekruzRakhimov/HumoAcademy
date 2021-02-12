package repository

import (
	"HumoAcademy/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (r *NewsPostgres) CreateNews(news models.News) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO news (title, short_desc, expire_at, img, full_desc, status) VALUES($1, $2, $3, $4, $5, $6) RETURNING id")
	row := r.db.QueryRow(query, news.Title, news.ShortDesc, news.ExpireAt, news.Img, news.FullDesc, news.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NewsPostgres) GetNewsByID (id int) (models.News, error) {
	var news models.News
	query := "SELECT * FROM news WHERE id = $1"
	err := r.db.Get(&news, query, id)
	return news, err
}

func (r *NewsPostgres) GetAllNews() ([]models.News, error) {
	var courses []models.News
	query := fmt.Sprintf("SELECT id, title, short_desc, expire_at, img, full_desc, status FROM news")
	err := r.db.Select(&courses, query)
	if err != nil {
		return []models.News{}, err
	}
	return courses, err
}