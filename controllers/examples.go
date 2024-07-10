package controllers

import (
	"context"
	"main/proto"
)

type Examples struct{}

func (srv Examples) DoExample(context context.Context, request *proto.Request) (*proto.Response, error) {
	return nil, nil
}
