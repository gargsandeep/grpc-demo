package server

import (
	"context"
	"grpc_demo/proto/packages/examplepb"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
)
var(
	port = "9999"
)

type YourServiceServer struct {

	//Echo(context.Context, *examplepb.StringMessage) (*examplepb.StringMessage, error)
}

func (*YourServiceServer) Echo(ctx context.Context, msg *examplepb.StringMessage) (*examplepb.StringMessage, error){
	return &examplepb.StringMessage{Value: msg.GetValue()},nil
}

func StartGRPC() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s %v", port, err)
	}

		grpcServer := grpc.NewServer()

		examplepb.RegisterYourServiceServer(grpcServer, &YourServiceServer{})
		log.Printf("grpc server starting to listen on %s", port)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc server over port %s %v", port, err)
		}
	log.Printf("lisnening on %s", port)
}