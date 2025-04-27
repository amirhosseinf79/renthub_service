package cloner_test

import (
	"testing"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	cloner "github.com/amirhosseinf79/renthub_service/internal/services/web_api_service"
	"github.com/stretchr/testify/assert"
)

func TestEditPrice(t *testing.T) {
	mockRepo := persistence.NewMockApiAuthRepo()
	logRepo := persistence.NewLogRepository(nil)
	tests := []struct {
		name    string
		fields  dto.UpdateFields
		wantErr bool
	}{
		{
			name: "valid request",
			fields: dto.UpdateFields{
				RequiredFields: dto.RequiredFields{
					UserID:   1,
					ClientID: "client1",
				},
				Dates: []string{
					"2025-06-07",
				},
				RoomID: "104598",
				Amount: 6500000,
			},
			wantErr: false,
		},
		{
			name: "invalid request",
			fields: dto.UpdateFields{
				RequiredFields: dto.RequiredFields{
					UserID:   1,
					ClientID: "client1",
				},
				Dates: []string{
					"2025-0428",
					"2025-0429",
				},
				RoomID: "104598",
			},
			wantErr: true,
		},
		{
			name: "unauthorized",
			fields: dto.UpdateFields{
				RequiredFields: dto.RequiredFields{
					UserID:   1,
					ClientID: "client2",
				},
				Dates: []string{
					"2025-0428",
					"2025-04-29",
				},
				RoomID: "104598",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := cloner.NewHomsaService(mockRepo, logRepo)
			_, err := service.Set("homsa").EditPricePerDays(tt.fields)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
