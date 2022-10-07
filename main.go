package main

import (
	"log"
	"net/http"
	"time"

	internal "github.com/bassamtiano/gotravel/internal"

	destinationAPI "github.com/bassamtiano/gotravel/api/destination"
	tagsTypeAPI "github.com/bassamtiano/gotravel/api/tags_type"
	"github.com/gorilla/mux"
)

func main() {

	DB := internal.InitDatabase()
	r := mux.NewRouter()

	destinationAPI.AddDestinationRouterHandler(r)
	tagsTypeAPI.AddTagsTypeRouterHandler(r)
	tagsTypeAPI.InitDB(DB)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
