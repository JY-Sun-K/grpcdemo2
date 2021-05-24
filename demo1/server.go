package main

import (
	"google.golang.org/grpc"
	"learngrpc/demo1/helper"
	"learngrpc/demo1/services"
	"net"
)

func main() {







	rpcServer:=grpc.NewServer(grpc.Creds(helper.GetServerCredentials()))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	services.RegisterOrderServiceServer(rpcServer,new(services.OrderService))
	lis,_:=net.Listen("tcp",":8081")
	rpcServer.Serve(lis)
	//mux:=http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request)
	//
	//	rpcServer.ServeHTTP(writer,request)
	//})
	//httpServer:=&http.Server{
	//	Addr: ":8081",
	//	Handler: mux,
	//}
	//httpServer.ListenAndServeTLS()
}
