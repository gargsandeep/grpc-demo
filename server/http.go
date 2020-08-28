package server

import (
	"context"
	"google.golang.org/grpc"
	"fmt"
	"log"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"grpc_demo/proto/packages/examplepb"
     _ "expvar"

)

var(
	httpport="8080"
)

type m struct {
	*runtime.JSONPb
}

func StartHttp(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Connect to the GRPC server
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	rmux := runtime.NewServeMux(
		runtime.WithMarshalerOption("*", &m{
			JSONPb: &runtime.JSONPb{
				EmitDefaults: true,
			},
		}),
	)


	yourServiceClient := examplepb.NewYourServiceClient(conn)
	err = examplepb.RegisterYourServiceHandlerClient(ctx, rmux, yourServiceClient)
	if err != nil{
		log.Printf("Error while registering client with grpc server, %v", err)
	}

	exposedMux := http.NewServeMux()

	swaggerUi := http.FileServer(http.Dir("swagger-ui"))
	exposedMux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", swaggerUi))
	proto := http.FileServer(http.Dir("proto"))
	exposedMux.Handle("/proto/", http.StripPrefix("/proto", proto))


	exposedMux.Handle("/", rmux)

	log.Printf("server starting to listen on %s", httpport)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", httpport), exposedMux)

	if err != nil {
		log.Fatal(err)
	}

}
