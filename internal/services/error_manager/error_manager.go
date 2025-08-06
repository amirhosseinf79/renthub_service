package error_manager

import (
	"strings"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
)

func ErrorLocalization(i string) string {
	switch i {
	case "You do not have permission to access.":
		return "شما اجازه دسترسی به این بخش را ندارید"
	}
	if strings.Contains(i, "core-api") {
		return dto.ErrTimeOut.Error()
	}
	return i
}
