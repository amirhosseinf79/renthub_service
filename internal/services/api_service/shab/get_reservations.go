package shab

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

// last_state[]=pended , page=100

func (h *service) GetReservations(fields dto.RecieveFields, response any) (*models.Log, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	fields.Filters["type"] = "host"
	fields.Filters["is_expired"] = 0

	err := h.handleGetResult(log, endpoint, fields, response)
	return log, err
}
