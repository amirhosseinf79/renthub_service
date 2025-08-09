package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

// status=3 , pagenumber=1

func (h *service) GetReservations(fields dto.RecieveFields, response any) (*models.Log, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	fields.Filters["pagenumber"] = 1
	fields.Filters["pagesize"] = 30
	fields.Filters["typeproduct"] = 1
	fields.Filters["type"] = 1

	err := h.handleGetResult(log, endpoint, fields, response)
	return log, err
}
