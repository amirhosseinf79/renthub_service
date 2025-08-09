package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) GetRoomDetails(fields dto.RecieveFields, result any) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.GetRoomDetails)
	endpoint := h.getEndpoints().GETRoomDetails
	err = h.handleGetResult(log, endpoint, fields, result)
	if err != nil {
		return log, err
	}
	return
}
