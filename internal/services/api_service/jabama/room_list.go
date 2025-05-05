package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) GetRoomList(fields dto.GetDetail, result any) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().GETRooms
	err = h.handleGet(log, nil, endpoint, fields, result)
	if err != nil {
		return log, err
	}
	return
}
