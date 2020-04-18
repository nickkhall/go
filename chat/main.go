package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	port := "8000"
	fmt.Println("Starting server...\n")

	listener, listenerErr := net.Listen("tcp", port)

	if listenerErr != nil {
		fmt.Println("ERROR! Could not start listener on %v", port)
		os.Exit(1)
	}


	fmt.Println("Server listening...\n")

}

