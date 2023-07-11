package main

import (
	"context"
	"fmt"
	v1 "github.com/dlfdyd96/proto-optional-test/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

type server struct {
	port string
	v1.UnimplementedYourServiceServer
}

func (s *server) mustEmbedUnimplementedYourServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *server) Echo(ctx context.Context, message *v1.TestMessage) (*v1.TestMessage, error) {
	fmt.Printf("%+v\n", message)

	fmt.Printf("string_value: %s\n", message.StringValue)
	if message.OptionalStringValue != nil {
		fmt.Printf("optional_string_value: %s\n", *message.OptionalStringValue)
	}
	fmt.Printf("int32_value: %d\n", message.Int32Value)
	if message.OptionalInt32Value != nil {
		fmt.Printf("optional_int32_value: %d\n", *message.OptionalInt32Value)
	}
	fmt.Printf("bool_value: %t\n", message.BoolValue)
	if message.OptionalBoolValue != nil {
		fmt.Printf("optional_bool_value: %t\n", *message.OptionalBoolValue)
	}

	return message, nil
}

func runServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Second * 10,
			Timeout:           time.Second * 20,
		}),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Second,
				PermitWithoutStream: true,
			}),
		grpc.MaxConcurrentStreams(5),
	)
	v1.RegisterYourServiceServer(s, &server{port, v1.UnimplementedYourServiceServer{}})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	fmt.Println("run server")
	go runServer("9090")

	<-make(chan int)
}
