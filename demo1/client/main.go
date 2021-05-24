package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"learngrpc/demo1/services"
	"log"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		log.Fatalf("加载客户端证书失败, err: %v\n", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		log.Fatalf("读取公钥文件失败: %v\n", err)
	}

	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})


	conn,err:=grpc.Dial(":8081",grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx:=context.Background()
	prodClient:= services.NewProdServiceClient(conn)
	prodRes,err:= prodClient.GetProdStock(ctx,&services.ProdRequest{ProdId: 12,ProdArea: services.ProdAreas_A})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
	resp,err:=prodClient.GetProdStocks(ctx,&services.QuerySize{Size: 2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Prodres)
	prod,err:=prodClient.GetProdInfo(ctx,&services.ProdRequest{
		ProdId:   101,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prod.ProdName)

	orderClient := services.NewOrderServiceClient(conn)
	orderTime := timestamp.Timestamp{Seconds: time.Now().Unix()}
	orderRes, err := orderClient.NewOrder(context.Background(),
		&services.OrderMain{
			OrderId:    100,
			OrderNo:    "酱油1001",
			UserId:     200,
			OrderMoney: 9.8,
			OrderTime:  &orderTime})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	fmt.Printf("订单状态 %v, 订单消息: %v\n", orderRes.Status, orderRes.Message)

}
