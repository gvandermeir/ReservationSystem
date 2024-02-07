package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gvandermeir/ReservationSystem/Internal/Server"
)

func main() {
	flag.Parse()
	port := flag.Int("port", 8080, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := server.NewGRPCServer()
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
