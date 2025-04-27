package cloner_test

import (
	"testing"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	cloner "github.com/amirhosseinf79/renthub_service/internal/services/web_api_service"
	"github.com/stretchr/testify/assert"
)

func TestOTPVerify(t *testing.T) {
	mockRepo := persistence.NewMockApiAuthRepo()
	tests := []struct {
		name    string
		fields  dto.RequiredFields
		otp     string
		wantErr bool
	}{
		{
			name: "valid request",
			fields: dto.RequiredFields{
				UserID:   1,
				ClientID: "client1",
			},
			otp:     "244413",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := cloner.NewHomsaService(mockRepo)
			err := service.Set("homsa").VerifyOtp(tt.fields, tt.otp)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
