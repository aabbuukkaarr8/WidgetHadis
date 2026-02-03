package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"widjetHadis/internal/apiserver"
	hadisHandler "widjetHadis/internal/handler/hadis"
	hadisRepo "widjetHadis/internal/repository/hadis"
	hadisService "widjetHadis/internal/service/hadis"
	"widjetHadis/internal/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	db := store.New()
	err = db.Open(config.Store.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	//repo
	hadisRepo := hadisRepo.NewRepository(db)
	//service
	hadisService := hadisService.NewService(hadisRepo)
	//handler
	hadisHandler := hadisHandler.NewHandler(hadisService)
	s := apiserver.New(config)
	s.ConfigureRouter(hadisHandler)
	if err := s.Run(); err != nil {
		panic(err)
	}



}
