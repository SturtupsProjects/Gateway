package pkg

import (
	"gateway/config"

	pbp "gateway/internal/generated/products"
	pbu "gateway/internal/generated/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient(cfg *config.Config) pbu.AuthServiceClient {
	conn, err := grpc.NewClient("crm-admin_auth"+cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Task Management Service: %v", err)
	}
	return pbu.NewAuthServiceClient(conn)
}
func NewProductClient(cfg *config.Config) pbp.ProductsClient {
	conn, err := grpc.NewClient("product"+cfg.PRODUCT_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Task Management Service: %v", err)
	}
	return pbp.NewProductsClient(conn)
}
