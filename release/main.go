package main

import (
	"api/dlna"
	"context"
	"embed"
	"fmt"
	"golib/chrome"
	"golib/server"
	"golib/server/http"
	"log"
)

//go:embed dist
var dist embed.FS

func main() {
	dlna.Initialize(-1000, -1001)
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	l1, port1, err := server.ServeLocal(http.NewServer())
	if err != nil {
		log.Fatal(err)
	}
	defer l1.Close()
	apiURL := fmt.Sprintf("http://127.0.0.1:%d", port1)

	c, err := chrome.New(context.Background(), "", "", 640, 480, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.DisableContextMenu()
	c.DisableDevTools()
	c.Bind("fetchURL", func() string {
		return apiURL
	})

	l2, port2, err := server.ServeLocalFile(dist)
	if err != nil {
		log.Fatal(err)
	}
	defer l2.Close()

	c.Load(fmt.Sprintf("http://127.0.0.1:%d/dist", port2))
	<-c.Done()
}
