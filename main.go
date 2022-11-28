package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeeDDHH/go-crud/data"
	"github.com/LeeDDHH/go-crud/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	data.Init()

	router.Init(r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
