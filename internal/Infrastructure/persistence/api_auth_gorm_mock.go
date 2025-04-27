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
				AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiOWY3MzZmNTc3YTk4ZmI4ZmU0NmEwY2M1M2VlMDk4YmUyMWFmODFlZDUxOTRiODNhZTcyMjI5N2ZiNjg0ZTFhZjQyNTNmOTM3MmIyNmYxYzAiLCJpYXQiOjE3NDUzNDgyNDAuNzIwNTk5LCJuYmYiOjE3NDUzNDgyNDAuNzIwNjAxLCJleHAiOjE3NDc5NDAyNDAuNzE3MjI4LCJzdWIiOiIxNzcxMTEiLCJzY29wZXMiOltdfQ.vLJWKTp3MCMDVCemFXTQm3jDBXsdXUfgJeQg65KSlRWSvPQi_ttJ-aAhYZg99NjSXw0pP3gbbINt9j2ac4sLTH-dSzWrdNoUrDWdFmH6XktwesOC_hKGOTZZT_ZYvKAOvMB6owHd342w5KZ_Prw1rNAyGSkc3xGsnCsKplUohLUvREJHHuUc1y_Xm7ThxBEPyJrwiPQ3DzBCTvPpacmFQtdk0hSwGuICzIUmXVYKxNvwwnD1aK5fY-nWyVw9YkNWJOIqa336XnjT5PmYLWygXDmsYClhIbibK27Okv7Pv3wa4bEw_voQhXIpaAQRwgNe7BTFJGxVqYCXjxf8QDWUhhaH5tjJAn3_BBJgvKayzzFDMvfwwjVd45p2Se6vIjFURTTpI2CoNbTrP3dmIMLGZXbwIRD2WXATfAPNtoXy5oOhnMfEyBXQO1hqND2-4wZWt5bpavy81DdG02drOZFqaiRNshCRymJtucdDhwlZBG2SZlg2OaO4h6xXNswro3uyXtYJcWq42OMlzMZ5Zl3Nc-5Fk0IbOe0DQ5aYvu7bl8umY__9YoCJH6bP8UtYUMgcUIgoCM23JLmuUwbaE4854iCpd6DK-eW19IjAGJOwPAx0YJ9q7CgozjLy_0ZI1Huxsg5iMmY2PWB4YxvF3G9k0uCIuLjbaK9gM0e_DO3DOUg",
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
