package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func (h *service) GetRoomList(fields dto.RecieveFields, result any) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.GetRoomList)
	endpoint := h.getEndpoints().GETRooms
	err = h.handleGetResult(log, endpoint, fields, result)
	return
}
