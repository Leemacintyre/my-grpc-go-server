package grpc

import (
	"context"
	"github.com/Leemacintyre/my-grpc-proto/protogen/go/hello"
)

func (a *Adapter) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	greet := a.helloService.GenerateHello(req.Name)

	return &hello.HelloResponse{
		Greet: greet,
	}, nil
}
