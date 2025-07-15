package homsa

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().RemoveDiscount
	bodies := h.generateRemoveDiscountBody(fields.Dates)
	for _, body := range bodies {
		err = h.handleUpdateResult(log, body, endpoint, fields)
		if err != nil {
			return log, err
		}
	}
	return log, err
}
