package main

import (
	"log"
	"ndavis20_server/controller"
	"ndavis20_server/service"
	"ndavis20_server/utils"
	"net/http"
)

func main() {
	utils.LoadDotEnv("../.env")

	// Create a new set of routes with the Path prefix /ndavis20
	service.NewRoute().
		Prefix("/ndavis20").

		// Endpoints that accept GET requests
		Method("GET").
		CreateEndpoint("/get_all", controller.GetAll).
		CreateEndpoint("/status", controller.GetLiveStatus).
		CreateEndpoint("/search", controller.Search).
		//Endpoints that accept PUT requests
		Method("PUT").
		CreateEndpoint("/ebay/deletion_notification", controller.ReceiveEbayDeleteNotif)

	// open a https listener with selfsigned certs on port 8080
	log.Fatalln(http.ListenAndServeTLS(":8080", "../34.207.90.86.crt", "../34.207.90.86.key", nil))
	//log.Fatalln(http.ListenAndServe(":8080", nil))
}
