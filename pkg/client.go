package pkg

import (
	"gateway/config"

	pbu "gateway/pkg/generated/user"
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