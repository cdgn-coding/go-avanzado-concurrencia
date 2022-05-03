package main

import (
	"github.com/gorilla/mux"
	"go-avanzado-concurrencia/7-cache-challenge/httpservice"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := mux.NewRouter()
	executable, errorIdentifyingExecutable := os.Executable()
	if errorIdentifyingExecutable != nil {
		log.Fatalln("Error identifying executable")
	}
	executablePath := filepath.Dir(executable)
	dataBasePath := filepath.Join(executablePath, "data")
	service := httpservice.NewService(dataBasePath)
	router.HandleFunc("/{key}", service.Get).Methods("GET")
	router.HandleFunc("/{key}", service.Delete).Methods("DELETE")
	router.HandleFunc("/{key}", service.Put).Methods("PUT")
	errStartingServer := http.ListenAndServe(":8080", router)
	if errStartingServer != nil {
		log.Fatalln("Failed to start up the HTTP httpservice.")
	}

}
