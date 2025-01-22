package main

import (
	"gateway/config"
	api "gateway/internal/api"
	"gateway/internal/api/token"
	"gateway/internal/minio"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
)

func main() {
	cfg := config.Load()
	path, err := os.Getwd()
	if err != nil {
		log.Println("mana 3")
		log.Fatal(err)
	}
	casbinEnforcer, err := casbin.NewEnforcer(path+"/internal/casbin/model.conf", path+"/internal/casbin/policy.csv")
	if err != nil {
		log.Println("mana 2")
		log.Fatal(err)
	}

	err = minio.InitMiniOClient()
	if err != nil {
		log.Println("mana 1 ")
		log.Fatal(err)
	}

	err = token.ConfigToken(cfg)
	if err != nil {
		log.Println("mana 0 ")
		log.Fatal(err)
	}

	r := api.NewRouter(casbinEnforcer, cfg)
	log.Fatal(r.Run(cfg.API_GATEWAY))
}
