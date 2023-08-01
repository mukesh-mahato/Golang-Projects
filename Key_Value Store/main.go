package main

import "sync"

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

func NewKVStore() *KVStore[string, int] {
	return &KVStore[string, int]{
		data: make(map[string]int),
	}
}

func StoreThings(s Storer[string, int]) error {
	return s.Put("foo", 1)
}

func main() {
	_ = NewKVStore()
	// StoreThings(kv)
}
