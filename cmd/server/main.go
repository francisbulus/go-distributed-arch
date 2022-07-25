package main

import (
	"log"

	"github.com/francisbulus/distributed/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
