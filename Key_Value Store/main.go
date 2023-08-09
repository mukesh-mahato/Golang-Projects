package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
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

type Server struct {
	Storage    Storer[string, string]
	ListenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		Storage:    NewKVStore[string, string](),
		ListenAddr: listenAddr,
	}
}

func (s *Server) handlePut(c echo.Context) error {
	key := c.Param("key")
	value := c.Param("value")

	s.Storage.Put(key, value)

	return c.JSON(http.StatusOK, map[string]string{"msg": "ok"})
}

func (s *Server) handleGet(c echo.Context) error {
	key := c.Param("key")

	value, err := s.Storage.Get(key)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"value": value})
}

func (s *Server) Start() {
	fmt.Printf("server is running on port %s", s.ListenAddr)

	e := echo.New()

	e.GET("/put/:key/:value", s.handlePut)
	e.GET("/get/:key", s.handleGet)

	e.Start(s.ListenAddr)

}

func main() {
	s := NewServer(":3000")

	s.Start()

}
