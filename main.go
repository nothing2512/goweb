package main

import (
	"fmt"
	"main/controllers"
	"main/db"
	"main/proto"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db.Connect()

	s := grpc.NewServer()
	proto.RegisterExamplesServer(s, &controllers.Examples{})

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	ln, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening Grpc On " + host)
	if err = http.Serve(ln, nil); err != nil {
		panic(err)
	}
}
