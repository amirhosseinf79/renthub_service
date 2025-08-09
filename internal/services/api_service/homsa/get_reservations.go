package homsa

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

// filter=pending , page=1

func (h *service) GetReservations(fields dto.RecieveFields, response any) (*models.Log, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	err := h.handleGetResult(log, endpoint, fields, response)
	return log, err
}
