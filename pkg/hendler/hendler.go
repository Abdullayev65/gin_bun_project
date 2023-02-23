package hendler

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/service"
	"github.com/Abdullayev65/gin_bun_project/pkg/utill"
)

type Handler struct {
	Service  *service.Service
	TokenJWT *utill.TokenJWT
}

func New(service *service.Service, TokenJWT *utill.TokenJWT) *Handler {
	return &Handler{Service: service, TokenJWT: TokenJWT}
}
