package service

import (
	"HumoAcademy/models"
	"HumoAcademy/pkg/repository"
)

type MainPageService struct {
	repo repository.MainPage
}

func NewMainPageService(repo repository.MainPage) *MainPageService {
	return &MainPageService{repo: repo}
}

func (s *MainPageService) GetAll() (models.MainPageContent, error) {
	return s.repo.GetAll()
}

func (s *MainPageService) AddUserForNews (user models.SubscribedUsers) error {
	return s.repo.AddUserForNews(user)
}