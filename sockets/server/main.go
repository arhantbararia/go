package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func processRequest(connection net.Conn ){
	buffer := make([]byte, 1024 )

	mLen, err := connection.Read(buffer)

	if(err != nil ){
		fmt.Println("Error reading!: " , err.Error())
	}

	fmt.Println("Recieved: ", string(buffer[:mLen]))

	_ ,err = connection.Write([]byte("Hi from server"))
	connection.Close()

}

func main() {
	fmt.Println("Server running......")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Printf("Error listening! Exiting %v ", err.Error())
		os.Exit(1)
	}

	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client")

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Printf("Error accepting: %v ", err.Error())
            os.Exit(1)
		}

		fmt.Println("Client Connected!")
	}

}
