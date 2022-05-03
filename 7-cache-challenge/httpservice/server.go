package httpservice

import (
	"github.com/gorilla/mux"
	"go-avanzado-concurrencia/7-cache-challenge/app"
	"io"
	"log"
	"net/http"
)

type service struct {
	app *app.Cache
}

func NewService(appBasePath string) *service {
	return &service{
		app: app.NewCache(appBasePath),
	}
}

func (service *service) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, keyGiven := vars["key"]

	if !keyGiven {
		http.Error(w, "Key not given", http.StatusBadRequest)
		return
	}

	value, keyFound := service.app.Get(key)
	if !keyFound {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(value))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func (service *service) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, keyGiven := vars["key"]

	if !keyGiven {
		http.Error(w, "Key not given", http.StatusBadRequest)
		return
	}

	service.app.Remove(key)
	w.WriteHeader(http.StatusAccepted)
}

func (service *service) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, keyGiven := vars["key"]

	if !keyGiven {
		http.Error(w, "Key not given", http.StatusBadRequest)
		return
	}
	bytes, errorReadingBody := io.ReadAll(r.Body)
	if errorReadingBody != nil {
		http.Error(w, "Cannot read body", http.StatusInternalServerError)
		return
	}
	value := string(bytes)
	service.app.Store(key, value)
	w.WriteHeader(http.StatusAccepted)
}
