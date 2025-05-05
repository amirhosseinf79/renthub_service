package jabama

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
)

func (h *service) EditPricePerDays(fields dto.UpdateFields) (log *models.Log, err error) {
	var result jabama_dto.RoomListResponse
	getFields := dto.GetDetail{
		RequiredFields: fields.RequiredFields,
	}
	log, err = h.GetRoomList(getFields, &result)
	if err != nil {
		return
	}

	for _, room := range result.Result.Items {
		if fields.RoomID == fmt.Sprintf("%v", room.Code) {
			fields.RoomID = room.ID
			break
		}
	}

	endpoint := h.getEndpoints().EditPricePerDay
	body := h.generatePriceBody(fields.Amount, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
