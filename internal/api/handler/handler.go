package handler

import (
	"gateway/config"
	pbc "gateway/internal/generated/company"
	pbd "gateway/internal/generated/debts"
	pbp "gateway/internal/generated/products"
	pbu "gateway/internal/generated/user"
	"log/slog"

	"gateway/pkg"

	logger "gateway/pkg/logs"
)

type Handler struct {
	UserClient    pbu.AuthServiceClient
	ProductClient pbp.ProductsClient
	CompanyClient pbc.CompanyServiceClient
	DebtClient    pbd.DebtsServiceClient
	PaymentClient pbd.PaymentServiceClient
	log           *slog.Logger
}

func NewHandlerRepo(cfg *config.Config) *Handler {
	return &Handler{
		UserClient:    pkg.NewUserClient(cfg),
		ProductClient: pkg.NewProductClient(cfg),
		CompanyClient: pkg.NewCompanyClient(cfg),
		DebtClient:    pkg.NewDebtClient(cfg),
		PaymentClient: pkg.NewPaymentClient(cfg),
		log:           logger.NewLogger(),
	}
}
