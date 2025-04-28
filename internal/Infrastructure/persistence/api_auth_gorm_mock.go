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
				ClientID:    "client1",
				Service:     "jabama",
				AccessToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkE2MDk5NjE0MUU5MTJDMDhBRjQyMEFGMjUyNjI2N0Q5NkNGRjUyRjZSUzI1NiIsInR5cCI6ImF0K2p3dCIsIng1dCI6InBnbVdGQjZSTEFpdlFncnlVbUpuMld6X1V2WSJ9.eyJuYmYiOjE3NDUzOTcyNTgsImV4cCI6MTc0NjAwMjA1OCwiaXNzIjoiaHR0cDovL2lkZW50aXR5LXNlcnZpY2UudGFyYWF6IiwiYXVkIjoiYXBpMSIsImNsaWVudF9pZCI6InJvLmNsaWVudCIsInN1YiI6IjFkNTU4YTk5LTk5OTgtNGMxMy1iMTE2LWI5ZjViODBjYTZmYyIsImF1dGhfdGltZSI6MTc0NTM5NzI1OCwiaWRwIjoibG9jYWwiLCJpcCI6IiIsImFiLWNoYW5uZWwiOiIiLCJyb2xlIjoiSG9zdCIsImlzX2hvc3QiOnRydWUsInVzZXJfdW5pcXVlX251bWJlciI6IjE0MTgzNTEiLCJqdGkiOiIxNEMyRkJDNjEyNTVDMDRGQzM4QzAyQkFDNTJBOThCNyIsImlhdCI6MTc0NTM5NzI1OCwic2NvcGUiOlsiYXBpMSIsIm9wZW5pZCIsIm9mZmxpbmVfYWNjZXNzIl0sImFtciI6WyJvdHAtdG9rZW4iXX0.YfuVN2dKiSOt5z3d8oF6rbGW3bEye3DuQX1JSQrnLImI2swyNCajvU3Mq0UCxlDNYBwVlrYtrw03HDnxTGUS5NSs7L7-mepleQnovxCPykCxGQIj7P1tpdTUbE3XszV8kLCvctBnI4stTg2UWYvoJ6EscHbwL2VGC-5-tvAogGiNr1i9-R3Xx0b5bHpD0etkjBYgzOZ2AL5t2LIvqD73Yf1V4wc3wwRFRK-KUHOJdmJB72BmzfLu3fXAnaBv3nhmOY5k2B42V_c_vDbXtpM6QG7IeLeuRM7ZKFyESZfR3fTfS-y76IkVuHu9idq4o_VEnsgRknUWgscYcO1IFtP_YA",
				Username:    "09334429096",
			},
			"1:client1:jajiga": {
				UserID:      1,
				ClientID:    "client1",
				Service:     "jajiga",
				AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI5Mjc0OTkxNC0wYjE3LTRlNmItYjk0MC04Y2Y5ODU4MzgyN2UiLCJqdGkiOiI0NDM1OGU4MDgwY2EyNjQ0ODZiMzgxMjgwOTliMjg1OTU5Mzk1NDUyZTMyMWI2YTVhZDU1ZDgwZmE4N2U2OThkYWNlZmU5ODIwMGFiZTVjYyIsImlhdCI6MTc0NTMzNzUzMy40NDk3MTYsIm5iZiI6MTc0NTMzNzUzMy40NDk3MjEsImV4cCI6MTc3Njg3MzUzMy40Mjg5NDcsInN1YiI6IjEzOTczNDgiLCJzY29wZXMiOltdfQ.tbtzrfz1MZFqT2yg3gWiiBj6tRh-qwGLrl8hKz2ayoPQwOBl45NAa_q1vR5oFYDV25YSQyCYch6KjxWXU6I_8XwMqfPJ6iBlPf0OcmaKv0h7QNO6HC4M-9K9bpIrgnYR3FHhx63O3z8JIsEkGqcuLFL2BcyAt0GJkHFxsfbaGv9y13cdCu4bI5L5DpyQ9Ej6SKUb1qWbGHxzUHkeR4RptNALJnTeWkC3T5R2yxRcVYOfSP8Uex4uHzI56IJTFsHPuBx7oUIERYLlhSa-ecQLPAZvLx1F9b2hoJT5o8Se9Rpt9W2qlfuAF-ZU_Ua-3P8P8tmYwXfG9BaeGdKRt6_uHbvTLt0Es9cbjJE7ChrDTYchHkFIEycDI3W7Z4EHpxglId4FmWX8nyE2ZGFzuN5pTtkCGlCC4SvgM76isH25jsttcmSjKV5KvTN3F0aHnaRPABpETerxnwGquX5xQ6ThqaeMOLSP3sYlY4FOX2jMPjxAAVIeiTyrVc9-beoJsl6blm7eQ6vHD-7gJbULqctOtNu6vHAFUwLRShe0tx8zgTeevtkM-Ty-rSLnq5SqayOphYFr-qR83fEBiqQFTKhd68HPJZNLItdevJbmJu94SGiAdhFa7zwAH2bdfQnqqYbDkaK1Us1SGfIi1uXxeFj8swFzxV7V1Saku5fcWrgWxFI",
				Username:    "09334429096",
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
