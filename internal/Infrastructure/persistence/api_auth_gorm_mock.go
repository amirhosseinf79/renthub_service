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
				AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiNmE0YTAwNWUwMzdiYjA0NjIxMjY4MjEwNDRmMDI2OTcyZjE2YTBlY2FkNmEzODc3ZmZiYzNlZjM3MjIwZTcwMDkyZmI2OTY1NWJiYjFmYTQiLCJpYXQiOjE3NDU3NDk5NzQuODM0Nzc5LCJuYmYiOjE3NDU3NDk5NzQuODM0NzgxLCJleHAiOjE3NDgzNDE5NzQuODMwNTQyLCJzdWIiOiIxNzcxMTEiLCJzY29wZXMiOltdfQ.sZI935Qqu73fxzA3tGlIwbt6p1MaZ_Gz7Pxk3bchkuA-zVh_giEio2O6xvs-uxmZnX89K9DSp3ygtscHjJu9ZEDLnyaU_RA_O6EV8bdWG6LgefJW_HG92_snnBN_2OW4UzCYDuURWrq-CIkWI4g5ps_73vYizdCQ_04XVB9wM-RrK1iLZJ3rnttS5IuzRiTrverlynx0dgB4v3Wc_PkJsohGRfr6n7Pn9dQcoeghAM5-awno16uy7aMOZ1clramWDYL3NRMzrHercoJL7k-2uWN6ZOsC6rtqF3LdAK7pIka_UagrEDlM9k2a6WiYp67HucLHyoBeicWg23Co6hm2UjnlptQron6N0T5BRjYBGKeYjqgMySP9dIFcP_QRyUCZnooOvH3_ErR0-6hYVDO1meiWWReUWyon0o4vGOZIWIcJ_g2arD9QfpvBvVk1siVVNXNCkSo6L0E1aKKb2Y2kwOMuatkNlFeG5O0bqavZLjoujPhtpNfBCQAGlGO8jPWvSVq4IiT3omtNgX_eR_5hVC-iFhxt2Bat1-dc3TAb6SQFsxxbP4CJAzebC2SWPEuax66wIewQDLUO719ICHbjD4tYBEzQxV0OPjoLjd_Fuff2AKMOPWUqjcEKtAZ0IJxYjZUNS6kXlOOyphQ8SK2jriB8dU3x5fXBX0b6tV5j0_o",
				Username:    "09334429096",
			},
			"1:client1:jabama": {
				UserID:      1,
				ClientID:    "client2",
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
