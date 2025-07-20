package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) AddDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log, err = h.AutoLogin(fields.RequiredFields)
	if err != nil {
		return log, err
	}
	log.Action = dto.AddDiscount
	endpoint := h.getEndpoints().AddDiscount
	body := h.generateDiscountBody(fields.RoomID, fields.Dates, fields.Amount)
	err = h.handleUpdateResult(log, body, endpoint, fields, 0)
	return log, err
}
