package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) EditPricePerDays(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.SetPrice)
	guestPrice := h.getAddGuestPrice(fields)
	endpoint := h.getEndpoints().EditPricePerDay
	body := h.generatePriceBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields, guestPrice)
	return log, err
}
