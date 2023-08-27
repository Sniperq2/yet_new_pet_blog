package internals

import (
	"log"
	"net"
	"net/http"
	"time"
	"yet_new_pet_blog/internals/router"
)

func NewServer(host string, port string) {

	hostPort := net.JoinHostPort(host, port)

	blogServer := &http.Server{
		Addr:         hostPort,
		Handler:      router.NewRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Printf("server start on %s", hostPort)
	err := blogServer.ListenAndServe()
	if err != nil {
		log.Println("could not start server")
	}

}
