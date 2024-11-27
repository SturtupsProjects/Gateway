package handler

import (
	"gateway/config"
	pbu "gateway/pkg/generated/user"

	"gateway/pkg"
	"log/slog"

	logger "gateway/pkg/logs"
)

type Handler struct {
	UserClient pbu.AuthServiceClient
	Logger     *slog.Logger
}

func NewHandlerRepo(cfg *config.Config) *Handler {
	return &Handler{
		UserClient: pkg.NewUserClient(cfg),
		Logger:     logger.NewLogger(),
	}
}
