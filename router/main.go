package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func Init() {
	router.HandleFunc("/ping", Ping).Methods("GET")
}

func StartServer(port int) {
	Init()
	log.Println("Server is listening on port " + strconv.Itoa(port))
	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
