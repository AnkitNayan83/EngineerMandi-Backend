package services

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/AnkitNayan83/SMA-backend/models"
	"github.com/AnkitNayan83/SMA-backend/repositories"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserService interface {
	HandleUserLogin(userInfo *models.OAuthUser) (*models.User, error)
	GetGoogleLoginUrl() string
	ExchangeCodeForToken(code string) (*oauth2.Token, error)
	FetchUserInfo(token *oauth2.Token) (*models.OAuthUser, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

var googleOAuth2Config *oauth2.Config

func InitializeOAuth() {
	googleOAuth2Config = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
}

func (s *userService) GetGoogleLoginUrl() string {
	return googleOAuth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (s *userService) ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return googleOAuth2Config.Exchange(context.Background(), code)
}

func (s *userService) FetchUserInfo(token *oauth2.Token) (*models.OAuthUser, error) {
	client := googleOAuth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user info")
	}

	var user models.OAuthUser

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *userService) HandleUserLogin(userInfo *models.OAuthUser) (*models.User, error) {

	existingUser, err := s.repo.FindUserByEmail(userInfo.Email)

	if err == nil {
		return existingUser, nil
	}

	user, err := s.repo.CreateUser(userInfo)

	if err != nil {
		return nil, err
	}

	return user, nil
}
