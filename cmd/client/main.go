package main

import (
	"context"
	"log"
	"net/http"

	greetv1 "connectgo/gen/todo/v1"
	"connectgo/gen/todo/v1/todov1connect"

	"connectrpc.com/connect"
)

func main() {
	client := todov1connect.NewToDoServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.CreateToDo(
		context.Background(),
		connect.NewRequest(&greetv1.CreateToDoRequest{ToDo: "Get it done"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.ToDos)
}
