package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"service-1/stub_proto"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//type Address struct {
//	Line        string `json:"line"`
//	AnotherLine string `json:"another_line"`
//}
//type Person struct {
//	Id      int32   `json:"id"`
//	Name    string  `json:"person"`
//	Age     int32   `json:"line"`
//	Address Address `json:"address"`
//}

// implement PersonService
type Server struct {
	//DB    *sql.DB
	stub_proto.UnimplementedPersonServiceServer
}

func (s *Server) GetAllPerson(ctx context.Context, in *stub_proto.ListRequest) (stub_proto.ListResponse, error) {
	//s.DB.Ping()
	var people []*stub_proto.Person
	for i := 1; i <= 10; i++ {
		var person stub_proto.Person
		person.Id = int32(i)
		person.Name = fmt.Sprintf("Aji %d", i)
		person.Age = 20 + int32(i)
		person.Address = &stub_proto.Address{Line: fmt.Sprintf("Line %d", i), AnotherLine: fmt.Sprintf("Another Line %d", i)}
		people = append(people, &person)
	}

	response := stub_proto.ListResponse{People: people, TotalResutls: 1}
	return response, nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	service := Server{}
	ctx := context.Background()

	r.Get("/person", func(w http.ResponseWriter, r *http.Request) {
		people, err := service.GetAllPerson(ctx, &stub_proto.ListRequest{Sender: "Wahidin"})
		if err != nil {
			msgErr := map[string]string{"messgae": err.Error()}
			responseBody, e := json.Marshal(msgErr)
			if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseBody)
			return
		}

		//var people []stub_proto.Person
		//for i := 1; i <= 10; i++ {
		//	var person stub_proto.Person
		//	person.Id = int32(i)
		//	person.Name = fmt.Sprintf("Aji %d", i)
		//	person.Age = 20 + int32(i)
		//	person.Address = &stub_proto.Address{Line: fmt.Sprintf("Line %d", i), AnotherLine: fmt.Sprintf("Another Line %d", i)}
		//	people = append(people, person)
		//}
		//
		responseBody, e := json.Marshal(people)
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
		return
	})
	http.ListenAndServe(":3000", r)
}
