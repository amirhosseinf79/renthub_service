package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) EditPricePerDays(fields dto.UpdateFields) (log *models.Log, err error) {
	// log = h.initLog(fields.UserID, fields.ClientID, dto.SetPrice)
	guestPrice, log, err := h.getAddGuestPrice(fields)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	if guestPrice < 0 {
		err = dto.ErrInvalidDate
		log.FinalResult = err.Error()
		return
	}
	log.Action = dto.SetPrice
	endpoint := h.getEndpoints().EditPricePerDay
	body := h.generatePriceBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields, guestPrice)
	return log, err
}
