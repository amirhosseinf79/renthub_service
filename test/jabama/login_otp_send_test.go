package cloner_test

import (
	"testing"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	cloner "github.com/amirhosseinf79/renthub_service/internal/services/web_api_service"
	"github.com/stretchr/testify/assert"
)

func TestOTPLogin(t *testing.T) {
	mockRepo := persistence.NewMockApiAuthRepo()
	logRepo := persistence.NewLogRepository(nil)
	tests := []struct {
		name        string
		fields      dto.RequiredFields
		phoneNumber string
		wantErr     bool
	}{
		{
			name: "Valid input",
			fields: dto.RequiredFields{
				UserID:   1,
				ClientID: "1",
			},
			phoneNumber: "09334429096",
			wantErr:     false,
		},
		{
			name: "inValid input",
			fields: dto.RequiredFields{
				UserID:   1,
				ClientID: "1",
			},
			phoneNumber: "0933442909",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := cloner.Newservice(mockRepo, logRepo)
			_, err := service.Set("jabama").SendOtp(tt.fields, tt.phoneNumber)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
