package main

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		fmt.Fprintln(w, fmt.Sprintf("Welcome user: %v \n", userID))
	})

	//this will work too
	// server := http.Server{
	// 	Addr:    s.addr,
	// 	Handler: RequestAuthenticationMiddleware(RequestLoggerMiddleWare(router)),
	// }

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}


	
	log.Printf("Server started on address: %v", s.addr)
	return server.ListenAndServe()

}
