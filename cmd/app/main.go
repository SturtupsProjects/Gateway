package main

import (
	"gateway/config"
	api "gateway/internal/api"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
)

func main() {
	cfg := config.Load()
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	casbinEnforcer, err := casbin.NewEnforcer(path+"/internal/casbin/model.conf", path+"/internal/casbin/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(casbinEnforcer, cfg)
	r.Run(cfg.API_GATEWAY)
}
