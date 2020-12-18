package main

import (
	"fmt"
	"log"
	"net"
	"northpole/server"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := server.NewServer()
	grpcServer.Serve(listener)
}
