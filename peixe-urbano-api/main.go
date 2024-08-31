package main

import (
	"fmt"
	"log"
	"net/http"
	"peixe-urbano/src/router"
)

func main() {
	fmt.Println("Running...")

	r := router.Create()

	log.Fatal(http.ListenAndServe(":5000", r))
}
