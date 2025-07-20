package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) AddDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().AddDiscount
	log = h.initLog(fields.UserID, fields.ClientID, dto.AddDiscount)
	err = h.handleUpdateResult(log, nil, endpoint, fields)
	return log, err
}
