package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Storer[K comparable, V any] interface {
	Put(K, V) error
	Get(K) (V, error)
	Update(K, V) error
	Delete(K) (V, error)
}

type KVStore[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// Hash checks if the given key is present in the store.
func (s *KVStore[K, V]) Hash(key K) bool {
	_, ok := s.data[key]
	return ok
}

func (s *KVStore[K, V]) Update(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.Hash(key) {
		return fmt.Errorf("the key (%v) doesn't exists.", key)
	}

	s.data[key] = value

	return nil
}

func (s *KVStore[K, V]) Put(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

func (s *KVStore[K, V]) Get(key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) doesn't exists", key)
	}
	return value, nil
}

func (s *KVStore[K, V]) Delete(key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) doesn't exists", key)
	}

	delete(s.data, key)
	return value, nil
}

func NewKVStore[K comparable, V any]() *KVStore[K, V] {
	return &KVStore[K, V]{
		data: make(map[K]V),
	}
}

type User struct {
	ID        int
	FirstName string
	Age       int
	Gender    string
}

type Server struct {
	Storage    Storer[int, *User]
	ListenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		Storage:    NewKVStore[int, *User](),
		ListenAddr: listenAddr,
	}
}

func (s *Server) handlePut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("foo"))
}

func (s *Server) Start() {
	fmt.Printf("server is running on port %s", s.ListenAddr)

	http.HandleFunc("/put", s.handlePut)

	log.Fatal(http.ListenAndServe(s.ListenAddr, nil))
}

func main() {
	s := NewServer(":3000")

	s.Start()

}
