package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/classes", addClass)
	router.GET("/classes_list", listClass)
	port := ":8080"
	fmt.Println("server started on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
