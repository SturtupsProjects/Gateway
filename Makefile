CURRENT_DIR = $(shell pwd)

DB_URL := postgres://postgres:0101@localhost:5432/?sslmode=disable

proto-gen:
	./pkg/scripts/gen-proto.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/router.go --output api/handler/docs


mig-up:
	migrate -path migrations -database '${DB_URL}' -verbose up

mig-down:
	migrate -path migrations -database '${DB_URL}' -verbose down

mig-force:
	migrate -path migrations -database '${DB_URL}' -verbose force 1

mig-create:
	migrate create -ext sql -dir migrations -seq auth_service_table

swag-gen:
	~/go/bin/swag init -g internal/api/router.go -o internal/api/docs
#   rm -r db/migrations

run:
	docker compose up -d --build