package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dan-santos/go-grpc/client"
	"github.com/dan-santos/go-grpc/proto"
)

func main(){
	jsonAddr := flag.String("json", ":3000", "listen address of json server is running")
	grpcAddr := flag.String("grpc", ":4000", "listen address of grpc server is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	ctx := context.Background()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func(){
		time.Sleep(2 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	server := NewJSONAPIServer(*jsonAddr, svc)
	server.Run()
}