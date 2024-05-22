package main

func main() {
	server := NewAPIServer("localhost:8000")

	server.Run()

}
