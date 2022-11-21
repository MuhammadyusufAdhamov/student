package main

import (
	"fmt"
	"log"

	"github.com/MuhammadyusufAdhamov/student/api"
	"github.com/MuhammadyusufAdhamov/student/config"
	"github.com/MuhammadyusufAdhamov/student/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main(){
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	fmt.Println(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	apiServer := api.New(&api.RouterOptions{
		Cfg:     &cfg,
		Storage: strg,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	log.Print("Server stopped")

}