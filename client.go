package main

import (
	"context"
	"fmt"
	"main/proto"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getClient() {
	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	client, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
	}
	srv := proto.NewExamplesClient(client)
	res, _ := srv.DoExample(context.TODO(), &proto.Request{
		Id:   "",
		Name: "",
	})
	_ = res.Status
	_ = res.Data
	_ = res.Message
}
