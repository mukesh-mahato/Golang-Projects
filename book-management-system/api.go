package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadGateway, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	books      []*Book
}

func NewAPISerer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) addBook(book *Book) {
	s.books = append(s.books, book)
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleBook))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetBook))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleCreateBook))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleDeleteBook))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleBook(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetBook(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateBook(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteBook(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)

}
func (s *APIServer) handleGetBook(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &Book{})

}
func (s *APIServer) handleCreateBook(w http.ResponseWriter, r *http.Request) error {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return err
	}
	s.addBook((&book))

	return WriteJSON(w, http.StatusCreated, &book)
}
func (s *APIServer) handleDeleteBook(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTrensfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
