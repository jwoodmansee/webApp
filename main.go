package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
		fmt.Println("you made here", w)
	}


func main() {
	fmt.Println("Server listening on 8080")
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("root").HTTPBox()))
	r.HandleFunc("/go", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(getPort(), handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}
