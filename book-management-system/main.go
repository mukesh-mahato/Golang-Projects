package main

func main() {
	server := NewAPISerer(":3000")
	server.Run()
}
