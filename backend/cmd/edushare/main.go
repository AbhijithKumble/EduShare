package main

import (
	"log"

	"github.com/AbhijithKumble/EduShare/backend/api"
	"github.com/AbhijithKumble/EduShare/backend/configs"
	"github.com/AbhijithKumble/EduShare/backend/db"
)

func main() {

	dbString := configs.Envs.Db
	portString := configs.Envs.Port

	// connect to the database
	connPool, err := db.ConnectDb(dbString)
	if err != nil {
		log.Fatalf("Could not connect to the database : %v", err)
	}

	defer connPool.Close()
    
	server := api.NewApiServer(portString, connPool.DB)

	if err = server.Run(); err != nil {
		log.Fatal("Server unable to Start", err)
	}
}
