package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	// "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/kyeett/grpc-experiments/proto/backend"
)

func main() {

	port := 10001
	serverAddr := fmt.Sprintf("localhost:%d", port)

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	cli := pb.NewBackendClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e := pb.Empty{}
	playerID, err := cli.NewPlayer(ctx, &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", playerID.GetID())

}
