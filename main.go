package main

import "fmt"

// package main
//
// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"northpole/server"
// )
//
// func main() {
// 	port := 8080
// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	grpcServer := server.NewServer()
// 	grpcServer.Serve(listener)
// }
func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	// close(ch)

	for n := range ch {
		fmt.Println(n)
	}
	fmt.Println("done")
}
