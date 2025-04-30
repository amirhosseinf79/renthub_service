package cloner_test

import (
	"testing"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	cloner "github.com/amirhosseinf79/renthub_service/internal/services/web_api_service"
	"github.com/stretchr/testify/assert"
)

func TestEasyLogin(t *testing.T) {
	mockRepo := persistence.NewMockApiAuthRepo()
	logRepo := persistence.NewLogRepository(nil)
	tests := []struct {
		name    string
		fields  dto.ApiEasyLogin
		wantErr bool
	}{
		{
			name: "Valid input",
			fields: dto.ApiEasyLogin{
				Username: "09109988333",
				Password: "mr0520691016",
			},
			wantErr: false,
		},
		{
			name: "invalid request",
			fields: dto.ApiEasyLogin{
				Username: "asdasdad",
				Password: "password123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := cloner.Newservice(mockRepo, logRepo)
			_, err := service.Set("").EasyLogin(tt.fields)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
