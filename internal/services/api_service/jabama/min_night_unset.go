package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) UnsetMiniNight(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().UnsetMinNight
	body := h.generateUnsetMinNightBody()
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
