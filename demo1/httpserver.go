package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"learngrpc/demo1/helper"
	"learngrpc/demo1/services"
	"log"
	"net/http"
)

func main() {
	// 使用Gateway启动HTTP Server
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		gwmux, "localhost:8081", opt)

	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}


		httpServer.ListenAndServe()



}