package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/repositories"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService interface {
	HandleUserLogin(userInfo *models.OAuthUser) (*models.User, bool, error)
	GetGoogleLoginUrl() string
	ExchangeCodeForToken(code string) (*oauth2.Token, error)
	FetchUserInfo(token *oauth2.Token) (*models.OAuthUser, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
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

func (s *authService) GetGoogleLoginUrl() string {
	return googleOAuth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (s *authService) ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return googleOAuth2Config.Exchange(context.Background(), code)
}

func (s *authService) FetchUserInfo(token *oauth2.Token) (*models.OAuthUser, error) {
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

	fmt.Print(string(body))

	err = json.Unmarshal(body, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *authService) HandleUserLogin(userInfo *models.OAuthUser) (*models.User, bool, error) {

	existingUser, err := s.repo.FindUserByEmail(userInfo.Email)

	if err == nil {

		isProfileConpleted := existingUser.Address != "" &&
			existingUser.City != "" &&
			existingUser.State != "" &&
			existingUser.Country != "" &&
			existingUser.PinCode != ""

		return existingUser, isProfileConpleted, nil
	}

	user, err := s.repo.CreateUser(userInfo)

	if err != nil {
		return nil, false, err
	}

	return user, true, nil
}
