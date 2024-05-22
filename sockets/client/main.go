package main

import (
	"fmt"
	"net"

)


const (
	SERVER_HOST= "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)


func main(){
	//establish connection
	connection, err := net.Dial(SERVER_TYPE , SERVER_HOST + ":" + SERVER_PORT )
	if err != nil {
		panic(err)
	}

	//send some data
	_ , err  = connection.Write([]byte("hello server!"))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	//Reading is done through buffeer
	// Each buffer is a chunk of 1024 bytes and
	// if the size of incoming messages > 1024 bytes, it gets broken into different buffer chunks
	if err != nil{
		fmt.Println("Error reading: ", err.Error() )
	}

	fmt.Println("Recieved: " , string(buffer[:mLen]))
	
	
	defer connection.Close()

}



// for clients we generally use
// net.Dial --> gives connection
// connection.Read
// connection.Write



