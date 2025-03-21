package main

import (
	_ "api/dlna"
	"golib/server"
	"golib/server/http"
	"log"
	"time"
)

func main() {
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	http.ReturnErr = true
	log.Fatal(server.ServeHTTP(":3000", http.NewServer(), 5*time.Second))
}
