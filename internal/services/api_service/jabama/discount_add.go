package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) AddDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	// log = h.initLog(fields.UserID, fields.ClientID, dto.AddDiscount)
	// log, err = h.updateRoomID(&fields)
	// if err != nil {
	// 	return
	// }
	endpoint := h.getEndpoints().AddDiscount
	cPrice, log, err := h.getCurrentPrice(fields)
	if err != nil {
		return
	}
	log.Action = dto.AddDiscount
	body := h.generateDiscountBody(&fields, cPrice)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
