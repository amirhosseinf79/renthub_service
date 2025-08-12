package homsa

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
)

// filter=pending , page=1

func (h *service) GetReservations(fields dto.RecieveFields) (*models.Log, interfaces.ReservationResponseInterface, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	var response homsa_dto.ReservationResponse
	err := h.handleGetResult(log, endpoint, fields, &response)
	return log, &response, err
}
