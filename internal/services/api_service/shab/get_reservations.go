package shab

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
)

// last_state[]=pended , page=100

func (h *service) GetReservations(fields dto.RecieveFields) (*models.Log, interfaces.ReservationResponseInterface, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	fields.Filters["type"] = "host"
	fields.Filters["is_expired"] = 0

	var response shab_dto.ReservationResponse
	err := h.handleGetResult(log, endpoint, fields, &response)
	return log, &response, err
}
