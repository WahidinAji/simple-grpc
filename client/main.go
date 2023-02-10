package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/WahidinAji/contract"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const sender = "wahidin"

var (
	addr = flag.String("addr", "localhost:9090", "the address to connect to")
	name = flag.String("name", sender, "Name to greet")
)

// func Handler(ctx context.Context) ([]contract.ListResponse, error) {

// }

func main() {
	flag.Parse()
	//set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to server: %v", err)
	}
	defer conn.Close()
	c := contract.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.List(ctx, &contract.ListRequest{Sender: *name})
	if err != nil {
		log.Fatalf("could not list person: %v", err)
	}
	// log.Printf("the people list: %v", res.GetPeople())

	chi := chi.NewRouter()
	chi.Use(middleware.Logger)
	chi.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responseBody, e := json.Marshal(res.GetPeople())
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
		return
	})
	http.ListenAndServe(":3000", chi)
}
