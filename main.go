package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"mvc/controllers"
)

func main() {
	router := httprouter.New()
    router.GET("/", controllers.Index)
    //router.GET("/hello/:name", Hello)
	router.ServeFiles("/resources/*filepath", http.Dir("resources"))
    log.Fatal(http.ListenAndServe(":8080", router))
}
