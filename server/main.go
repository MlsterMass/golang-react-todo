package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"golang-react-todo/server/router"
	"log"
)

func main() {
	r := router.InitRoutes()
	fmt.Println("starting the server on port 8888...")
	log.Fatal(r.Run(":8888"))
}
