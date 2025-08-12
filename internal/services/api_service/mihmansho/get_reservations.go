package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
)

// status=3 , pagenumber=1

func (h *service) GetReservations(fields dto.RecieveFields) (*models.Log, interfaces.ReservationResponseInterface, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	fields.Filters["pagenumber"] = 1
	fields.Filters["pagesize"] = 30
	fields.Filters["typeproduct"] = 1
	fields.Filters["type"] = 1

	var response mihmansho_dto.ReservationResponse
	err := h.handleGetResult(log, endpoint, fields, &response)
	return log, &response, err
}
