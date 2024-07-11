package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/AbhijithKumble/EduShare/backend/services/serverhealth"
	"github.com/AbhijithKumble/EduShare/backend/services/signup"
	"github.com/AbhijithKumble/EduShare/backend/services/verify"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(port string, db *sql.DB) (s *ApiServer) {
	return &ApiServer{
		addr: ":" + port,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	//use functions to register the routes to subrouter
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	health := serverhealth.NewHandler()
	health.RegisterRoutes(subRouter)

	signupStore := signup.NewStore(s.db)
	signupHandler := signup.NewHandler(signupStore)
	signupHandler.RegisterRoutes(subRouter)

	verifyStore := verify.NewStore(s.db)
	verifyHandler := verify.NewHandler(verifyStore)
	verifyHandler.RegisterRoutes(subRouter)

	// add routing config above
	srv := &http.Server{
		Addr:         s.addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server started at port", s.addr)

	return srv.ListenAndServe()
}
