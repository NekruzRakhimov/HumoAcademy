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

func (r *NewsPostgres) ChangeNewsImg(id int, img string) error {
	query := fmt.Sprintf("UPDATE news SET img=$1 WHERE id=$2")

	_, err := r.db.Exec(query, img, id)
	if err != nil {
		return err
	}

	return nil
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
	query := fmt.Sprintf("SELECT * FROM news WHERE id = $1")
	err := r.db.Get(&news, query, id)
	return news, err
}

func (r *NewsPostgres) GetNewsImgSrc (id int) (string, error) {
	var imgSrc string
	query := fmt.Sprintf("SELECT img FROM news WHERE id=$1")
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&imgSrc); err != nil {
		return "", err
	}
	return imgSrc, nil
}

func (r *NewsPostgres) GetAllMiniNews() ([]models.MiniNews, error) {
	var courses []models.MiniNews
	query := fmt.Sprintf("SELECT id, title, short_desc, img, status FROM news")
	err := r.db.Select(&courses, query)
	if err != nil {
		return []models.MiniNews{}, err
	}
	return courses, err
}

func (r *NewsPostgres) EditNews(id int, news models.News) error {
	query := fmt.Sprintf("UPDATE news SET title=$1, short_desc=$2, expire_at=$3, full_desc=$4, Status=$5 WHERE id=$6")

	_, err := r.db.Exec(query, news.Title, news.ShortDesc, news.ExpireAt, news.Img, news.FullDesc, news.Status, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *NewsPostgres) ChangeNewsStatus (id int, status bool) error {
	query := fmt.Sprintf("UPDATE news SET status = $1 where id = $2")
	_, err := r.db.Exec(query, status, id)
	if err != nil {
		return err
	}
	return nil
}