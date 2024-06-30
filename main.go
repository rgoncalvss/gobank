package main

func main() {
	server := NewApiServer(":3000")
	server.Run()
}
