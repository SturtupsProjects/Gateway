package pkg

import (
	"gateway/config"

	pbc "gateway/internal/generated/company"
	pbd "gateway/internal/generated/debts"
	pbp "gateway/internal/generated/products"
	pbu "gateway/internal/generated/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient(cfg *config.Config) pbu.AuthServiceClient {
	conn, err := grpc.NewClient("crm-admin_auth"+cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to User Service: %v", err)
	}
	return pbu.NewAuthServiceClient(conn)
}

func NewProductClient(cfg *config.Config) pbp.ProductsClient {
	conn, err := grpc.NewClient("localhost"+cfg.PRODUCT_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Product Service: %v", err)
	}
	return pbp.NewProductsClient(conn)
}

func NewCompanyClient(cfg *config.Config) pbc.CompanyServiceClient {
	conn, err := grpc.NewClient("crm-admin_auth"+cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Client Service: %v", err)
	}
	return pbc.NewCompanyServiceClient(conn)
}

func NewDebtClient(cfg *config.Config) pbd.DebtsServiceClient {
	conn, err := grpc.NewClient("debts-service"+cfg.DEBT_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Debt Service: %v", err)
	}
	return pbd.NewDebtsServiceClient(conn)
}
