package main

import (
	"gateway/config"
	api "gateway/internal/api"
	"gateway/internal/api/token"
	"gateway/internal/minio"
	logger "gateway/pkg/logs"
	"github.com/casbin/casbin/v2"
	"io"
	"log"
	"net/http"
	"os"
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

	log1 := logger.NewLogger()

	r := api.NewRouter(casbinEnforcer, cfg, log1)

	//ips, err := get()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log1.Debug(fmt.Sprintf("[%s] Api Gateway is running at IP: %s",
	//	time.Now().Format("02-01-2006 15:04:05"), ips))

	log.Fatal(r.Run(cfg.API_GATEWAY))
}

// Getting
func get() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text") // Этот сервис возвращает IPv4
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
