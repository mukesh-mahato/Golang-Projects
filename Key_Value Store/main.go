package main

import (
	"fmt"
	"log"
	"sync"
)

type Storer[K comparable, V any] interface {
	Put(K, V) error
	Get(K) ([]byte, error)
	Update(K, []byte) error
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

func NewKVStore[K comparable, V any]() *KVStore[K, V] {
	return &KVStore[K, V]{
		data: make(map[K]V),
	}
}

func StoreThings(s Storer[string, int]) error {
	return s.Put("foo", 1)
}

func main() {
	store := NewKVStore[string, string]()

	if err := store.Put("foo", "bar"); err != nil {
		log.Fatal(err)
	}

	value, err := store.Get("foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	if err := store.Put("foo", "oof"); err != nil {
		log.Fatal(err)
	}

	value, err = store.Get("foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
