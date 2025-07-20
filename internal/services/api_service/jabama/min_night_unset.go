package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) UnsetMiniNight(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.UnsetMinNight)
	endpoint := h.getEndpoints().UnsetMinNight
	err = h.handleUpdateResult(log, nil, endpoint, fields)
	return log, err
}
