package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/kyeett/grpc-experiments/proto/backend"
	"github.com/kyeett/grpc-experiments/types"
)

func main() {

	t1 := &types.Entity{
		"Magnus", 123,
	}

	pl, err := GobMarshal(t1) // b.Bytes()
	if err != nil {
		log.Fatal(err)
	}

	a := backend.EntityResponse{
		Payload: pl,
	}

	serializedA, err := proto.Marshal(&a)
	if err != nil {
		log.Fatal("could not serialize anything")
	}
	// unmarshal to simulate coming off the wire
	var a2 backend.EntityResponse
	if err := proto.Unmarshal(serializedA, &a2); err != nil {
		log.Fatal("could not deserialize anything")
	}

	var t2 *types.Entity
	GobUnmarshal(a2.GetPayload(), &t2)

	fmt.Println(t1, t2)

}
