package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

// status=awaiting_confirmation , page=1

func (h *service) GetReservations(fields dto.RecieveFields, response any) (*models.Log, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	_, ok := fields.Filters["page"]
	if !ok {
		fields.Filters["page"] = 1
	}

	err := h.handleGetResult(log, endpoint, fields, response)
	return log, err
}
