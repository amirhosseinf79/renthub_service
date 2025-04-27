package cloner

import "github.com/amirhosseinf79/renthub_service/internal/dto"

func (h *homsaService) VerifyOtp(fields dto.RequiredFields, otp string) (err error) {
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		return
	}
	if otp == "" {
		return dto.ErrEmptyCode
	}
	field := dto.ApiEasyLogin{
		RequiredFields: fields,
		Username:       model.Username,
		Password:       otp,
	}
	response, err := h.performLoginRequest(field, true)
	if err != nil {
		return
	}
	err = h.updateOrCreateAuthRecord(field, response.GetToken())
	if err != nil {
		return
	}
	return nil
}
