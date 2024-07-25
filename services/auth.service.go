package services

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuth2Config *oauth2.Config

type OAuthUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitializeOAuth() {
	googleOAuth2Config = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
}

func GetGoogleLoginUrl() string {
	return googleOAuth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return googleOAuth2Config.Exchange(context.Background(), code)
}

func FetchUserInfo(token *oauth2.Token) (*OAuthUser, error) {
	client := googleOAuth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user info")
	}

	var user OAuthUser

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
