package main

import (
	"fmt"
	"log"
	"net/http"
	"peixe-urbano/src/config"
	"peixe-urbano/src/router"
)

func main() {
	config.Load()

	fmt.Printf("Running... at port: %d", config.Port)

	r := router.Create()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
