package service

import (
	"HumoAcademy/models"
	"HumoAcademy/pkg/repository"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) CreateNews(news models.News) (int, error) {
	return s.repo.CreateNews(news)
}

func (s *NewsService) EditNews(id int, news models.News) error {
	return s.repo.EditNews(id, news)
}

func (s *NewsService) GetNewsByID (id int) (models.News, error) {
	return s.repo.GetNewsByID(id)
}

func (s *NewsService) GetAllMiniNews () ([]models.MiniNews, error) {
	return s.repo.GetAllMiniNews()
}

func (s *NewsService) ChangeNewsStatus (id int, status bool) error {
	return s.repo.ChangeNewsStatus(id, status)
}

func (s *NewsService) ChangeNewsImg(id int, img string) error {
	return s.repo.ChangeNewsImg(id, img)
}

func (s *NewsService) GetNewsImgSrc (id int) (string, error) {
	return s.repo.GetNewsImgSrc(id)
}