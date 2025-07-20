package shab

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.RemoveDiscount)
	endpoint := h.getEndpoints().RemoveDiscount
	body := h.generateRemoveDiscountBody(fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
