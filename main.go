package main

import (
	"grpc_demo/server"
	"sync"
)

func main(){
	go server.StartGRPC()
	go server.StartHttp()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
