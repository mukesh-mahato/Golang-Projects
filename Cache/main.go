package main

import cache "github.com/mukeshmahato18/cache/Cache"

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	server := NewServer(opts, cache.NewCacher())
	server.Start()
}
