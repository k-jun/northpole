package main

import (
	"fmt"
	"grpcserver/server"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	defaultPort = 8080
)

func main() {
	godotenv.Load("./.env")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("PORT env variable was invalid")
		log.Println("using default port")
		port = defaultPort
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := server.NewServer()
	grpcServer.Serve(listener)
}
