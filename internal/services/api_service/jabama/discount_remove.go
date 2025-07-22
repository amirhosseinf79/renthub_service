package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	// log = h.initLog(fields.UserID, fields.ClientID, dto.RemoveDiscount)
	// log, err = h.updateRoomID(&fields)
	// if err != nil {
	// 	return
	// }
	endpoint := h.getEndpoints().RemoveDiscount
	cPrice, log, err := h.getCurrentPrice(fields)
	if err != nil {
		return
	}
	log.Action = dto.RemoveDiscount
	body := h.generateDiscountBody(&fields, cPrice)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
