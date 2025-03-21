package main

import (
	"api/dlna"
	"golib/server"
	"golib/server/http"
	"log"
	"time"
)

func main() {
	dlna.Initialize(-1000, -1001)
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	http.ReturnErr = true
	log.Fatal(server.ServeHTTP(":3000", http.NewServer(), 5*time.Second))
}
