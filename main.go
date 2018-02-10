package main

import (
	"net/http"
	"../visa/routes"
	"log"
)

func main() {

	//Starting the API server
	router := routes.VisaRoutes()
	http.Handle("/api/", router)

	//Starting the FileServer
	fs := http.FileServer(http.Dir("server/webapps/visa/components/visa_data"))
	http.Handle("/", fs)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3006", nil))

}