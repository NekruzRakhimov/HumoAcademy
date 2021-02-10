package service

import (
	"HumoAcademy/models"
	"HumoAcademy/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetAdmin(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AdminService) CreateAdmin(admin models.Admin) (int, error) {
	generatePasswordHash(admin.Password)
	return s.repo.CreateAdmin(admin)
}

//func (s *AdminService) ParseToken(token string) (int, error) {
//	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("invalid signing method")
//		}
//
//		return []byte(signingKey), nil
//	})
//	if err != nil {
//		return 0, err
//	}
//
//	claims, ok := token.Claims.(*tokenClaims)
//	if !ok {
//		return 0, errors.New("token claims are not of type *tokenClaims")
//	}
//
//	return claims.UserId, nil
//}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}