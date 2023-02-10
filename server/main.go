package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	c "github.com/WahidinAji/contract"

	"google.golang.org/grpc"
)

type server struct {
	c.UnimplementedPersonServiceServer
}

func (s *server) List(ctx context.Context, in *c.ListRequest) (*c.ListResponse, error) {
	log.Printf("received service: %v", in.Sender)
	var people []*c.Person
	for i := 1; i <= 10; i++ {
		var person c.Person
		person.Id = int32(i)
		person.Name = fmt.Sprintf("Aji %d", i)
		person.Age = 20 + int32(i)
		person.Address = &c.Address{Line: fmt.Sprintf("Line %d", i), AnotherLine: fmt.Sprintf("Another Line %d", i)}
		people = append(people, &person)
	}

	return &c.ListResponse{People: people, TotalResutls: 10}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *flag.Int("port", 9090, "The server port to listen on")))
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	c.RegisterPersonServiceServer(s, &server{})
	log.Printf("server listenign at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
