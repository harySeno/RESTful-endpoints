package main

import (
	"RESTful-endpoints/pkg/restapi"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	// create new web service
	webService := new(restful.WebService)

	// create a new route
	webService.Path("/hello").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Route(webService.GET("").
			To(restapi.Hello))

	// create a new container
	container := restful.NewContainer()
	container.Add(webService)

	// initiate server
	log.Println(" Initiating server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", container))
}
