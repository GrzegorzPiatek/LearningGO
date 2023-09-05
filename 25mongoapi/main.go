package main

import (
	"25mongoapi/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")

	fmt.Println("Server start")
	r := router.Router()
	err := http.ListenAndServe(":4000", r) //	fmt.Println("Listening at port 4000.")

	log.Fatal(err)
}
