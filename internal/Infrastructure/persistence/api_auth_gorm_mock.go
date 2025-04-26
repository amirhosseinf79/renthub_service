package persistence

import (
	"errors"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/domain/repository"
)

type MockApiAuthRepo struct {
	data map[string]*models.ApiAuth
}

func NewMockApiAuthRepo() repository.ApiAuthRepository {
	return &MockApiAuthRepo{
		data: map[string]*models.ApiAuth{
			"1:client1:homsa": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "homsa",
				AccessToken: "token1",
			},
			"1:client1:jabama": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "jabama",
				AccessToken: "token2",
			},
			"1:client1:jajiga": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "jajiga",
				AccessToken: "token3",
			},
			"1:client1:mihmansho": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "mihmansho",
				AccessToken: "token3",
			},
			"1:client1:otaghak": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "otaghak",
				AccessToken: "token3",
			},
			"1:client1:shab": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "shab",
				AccessToken: "token3",
			},
		},
	}
}

func (r *MockApiAuthRepo) CheckExists(userID uint, clientID string, service string) bool {
	key := generateKey(userID, clientID, service)
	_, exists := r.data[key]
	return exists
}

func (r *MockApiAuthRepo) GetByUnique(userID uint, clientID string, service string) (*models.ApiAuth, error) {
	key := generateKey(userID, clientID, service)
	if auth, exists := r.data[key]; exists {
		return auth, nil
	}
	return nil, errors.New("record not found")
}

func (r *MockApiAuthRepo) GetAll(userID uint, clientID string) ([]*models.ApiAuth, error) {
	var list []*models.ApiAuth
	for _, auth := range r.data {
		if auth.UserID == userID && auth.ClientID == clientID {
			list = append(list, auth)
		}
	}
	return list, nil
}

func (r *MockApiAuthRepo) Create(token *models.ApiAuth) error {
	key := generateKey(token.UserID, token.ClientID, token.Service)
	r.data[key] = token
	return nil
}

func (r *MockApiAuthRepo) Update(token *models.ApiAuth) error {
	key := generateKey(token.UserID, token.ClientID, token.Service)
	if _, exists := r.data[key]; !exists {
		return errors.New("record not found")
	}
	r.data[key] = token
	return nil
}

func generateKey(userID uint, clientID, service string) string {
	return fmt.Sprintf("%d:%s:%s", userID, clientID, service)
}
