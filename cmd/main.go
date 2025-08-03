package main

import (
	"Evermos-Virtual-Intern/config"
	"Evermos-Virtual-Intern/internal/builder"
	"Evermos-Virtual-Intern/internal/common"
	"Evermos-Virtual-Intern/pkg/database"
	"Evermos-Virtual-Intern/pkg/server"
)


func main() {
	cfg, err := config.NewConfig("../.env")
	checkError(err)

	db, err := database.ConnectToMysql(cfg)
	checkError(err)

	common.AppConfig = cfg

	publicRoutes := builder.BuildPublicRoutes(db, cfg)
	privateRoute := builder.BuildPrivateRoutes(db, cfg)



	srv := server.NewServer(cfg, publicRoutes, privateRoute)
	srv.Run()
	srv.GracefulShutdown()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}