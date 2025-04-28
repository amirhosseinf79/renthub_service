package cloner

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *homsaService) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().RemoveDiscount
	body := h.generateRemoveDiscountBody(fields.RoomID, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields.UserID, fields.ClientID, fields.RoomID)
	return log, err
}
