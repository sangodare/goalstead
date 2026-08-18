package main

import (
	"fmt"
	"goalstead/controllers"
	"goalstead/templates"
	"goalstead/views"
	"net/http"
)

func main() {
	err := server()

	if err != nil {
		panic(err)

	}

}

func server() error {

	//ROUTES,TEMPLATES AND CONTROLLERS

	//Static files
	cssFile := http.FileServer(http.Dir("assets"))

	//Templates
	tpl := views.Must(views.ParseFS(templates.FS, "home.html", "reusable.html"))

	//Routes
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", cssFile))
	mux.Handle("GET /", controllers.StaticHandler(tpl))

	//Start up the Server
	fmt.Println("starting up the server on port 4000")
	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	return server.ListenAndServe()

}
