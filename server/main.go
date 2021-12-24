package main

import (
	"fmt"
	"github.com/MlsterMass/golang-react-todo/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 9000...")

	log.Fatal(http.ListenAndServe(":8888", r))
}
