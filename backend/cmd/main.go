package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AbhijithKumble/EduShare/backend/configs"
	"github.com/gorilla/mux"
)


func main() {

	router := mux.NewRouter()
    router.Host(configs.Envs.CLIENT)

	server := &http.Server{
		Addr:         configs.Envs.PORT,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	fmt.Printf("Server running at PORT : 8080 \n")

	log.Fatal(server.ListenAndServe())

}
