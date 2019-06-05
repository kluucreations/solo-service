package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/kluucreations/solo-service/gen/restapi"
	"github.com/kluucreations/solo-service/gen/restapi/operations"
	"github.com/kluucreations/solo-service/internal/datastore"
	"github.com/kluucreations/solo-service/solo"
)

func main() {
	// load swagger spec
	swaggerSpec, err := loads.Analyzed(
		restapi.FlatSwaggerJSON,
		"2.0",
	)
	if err != nil {
		log.Panicln("Unable to analyze swaggerSpec", err)
	}
	api := operations.NewSoloServiceAPI(swaggerSpec)

	// instantiate connection to datastore
	ds := datastore.NewClient()
	// create service. inject dependencies
	service := solo.NewService(ds)

	solo.Configure(api, service)
	server := restapi.NewServer(api)
	// hardcode port to 8080 for ease and demonstrative purpose
	server.Port = 8080
	// gracefully shutdown server
	defer func() {
		log.Println("shutting down...")
		server.Shutdown()
	}()
	// debug logging message
	log.Println("starting server...")

	if err := server.Serve(); err != nil {
		log.Panicln("server err:", err)
	}
}
