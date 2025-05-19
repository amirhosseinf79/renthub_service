package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log, err = h.AutoLogin(fields.RequiredFields)
	if err != nil {
		return log, err
	}
	endpoint := h.getEndpoints().RemoveDiscount
	body := h.generateDiscountBody(fields.RoomID, fields.Dates, 0)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
