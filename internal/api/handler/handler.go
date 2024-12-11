package handler

import (
	"gateway/config"
	pbp "gateway/pkg/generated/products"
	pbu "gateway/pkg/generated/user"

	"gateway/pkg"
	"log/slog"

	logger "gateway/pkg/logs"
)

type Handler struct {
	UserClient    pbu.AuthServiceClient
	ProductClient pbp.ProductsClient
	log           *slog.Logger
}

func NewHandlerRepo(cfg *config.Config) *Handler {
	return &Handler{
		UserClient:    pkg.NewUserClient(cfg),
		ProductClient: pkg.NewProductClient(cfg),
		log:           logger.NewLogger(),
	}
}
