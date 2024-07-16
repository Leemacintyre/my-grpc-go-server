package grpc

import (
	"fmt"
	"github.com/Leemacintyre/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
	"my-grpc-go-server/port"
	"net"
)

type Adapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *Adapter {
	return &Adapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (a *Adapter) Run() {
	var err error
	log.Printf("port %d", a.grpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen to port %d : %v \n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n\n", a.grpcPort)

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server grpc on gRPC port %d : %v elloService \n", a.grpcPort, err)
	}
}

func (a *Adapter) stop() {
	a.server.Stop()
}
