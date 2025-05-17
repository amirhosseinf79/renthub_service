package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) AddDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().AddDiscount
	log = h.initLog(fields.UserID, fields.ClientID)
	body := h.generateDiscountBody(fields.RoomID, fields.Dates, fields.Amount)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
