package main

import (
	"log"
	"net/http"

	"github.com/micahasowata/micro/metadata/internal/controller/metadata"
	httphandler "github.com/micahasowata/micro/metadata/internal/handler/http"
	"github.com/micahasowata/micro/metadata/internal/repository/memory"
)

func main() {
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("GET /metadata", http.HandlerFunc(h.GetMetadata))

	log.Println("starting metadata service")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
