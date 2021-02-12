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

func (s *NewsService) GetNewsByID (id int) (models.News, error) {
	return s.repo.GetNewsByID(id)
}

func (s *NewsService) GetAllNews () ([]models.News, error) {
	return s.repo.GetAllNews()
}