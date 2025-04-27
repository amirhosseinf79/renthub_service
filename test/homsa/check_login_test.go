package cloner_test

import (
	"testing"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	cloner "github.com/amirhosseinf79/renthub_service/internal/services/web_api_service"
	"github.com/stretchr/testify/assert"
)

func TestCheckLogin(t *testing.T) {
	mockRepo := persistence.NewMockApiAuthRepo()
	tests := []struct {
		name    string
		fields  dto.RequiredFields
		wantErr bool
	}{
		{
			name: "valid request",
			fields: dto.RequiredFields{
				UserID:   1,
				ClientID: "client1",
			},
			wantErr: false,
		},
		{
			name: "unauthorized request",
			fields: dto.RequiredFields{
				UserID:   1,
				ClientID: "client2",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := cloner.NewHomsaService(mockRepo)
			err := service.Set("homsa").CheckLogin(tt.fields)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
