package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/AbhijithKumble/EduShare/backend/services/courses"
	"github.com/AbhijithKumble/EduShare/backend/services/dept"
	"github.com/AbhijithKumble/EduShare/backend/services/serverhealth"
	"github.com/AbhijithKumble/EduShare/backend/services/user"
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


// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Frontend origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *ApiServer) Run() error {
	//use functions to register the routes to subrouter
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	health := serverhealth.NewHandler()
	health.RegisterRoutes(subRouter)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	deptStore := dept.NewStore(s.db)
	deptHandler := dept.NewHandler(deptStore)
	deptHandler.RegisterRoutes(subRouter)

	courseStore := courses.NewStore(s.db)
	courseHandler := courses.NewHandler(courseStore)
	courseHandler.RegisterRoutes(subRouter)

  // Wrap the router with CORS middleware
	corsHandler := corsMiddleware(router)

	// add routing config above
	srv := &http.Server{
		Addr:         s.addr,
		Handler:      corsHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server started at port", srv.Addr)

	return srv.ListenAndServe()
}
