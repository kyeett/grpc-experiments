package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"strings"

	// "github.com/johanbrandhorst/protobuf/grpcweb"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"github.com/gorilla/websocket"
	"github.com/kyeett/grpc-experiments/frontend/bundle"
	"github.com/kyeett/grpc-experiments/proto/backend"
	pb "github.com/kyeett/grpc-experiments/proto/backend"
	"github.com/kyeett/grpc-experiments/types"
	"github.com/lpar/gzipped"
	"google.golang.org/grpc"
)

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
}

// Ensure struct implements interface
var _ backend.BackendServer = (*Backend)(nil)

func (b *Backend) NewPlayer(ctx context.Context, _ *pb.Empty) (*pb.PlayerID, error) {
	fmt.Println("Send send :-)")
	return &pb.PlayerID{
		ID: "Magnus",
	}, nil
}

func (b *Backend) EntityStream(_ *pb.Empty, stream pb.Backend_EntityStreamServer) error {

	for _, v := range []int{10, 15, 1, 123, 10, 1, 3, 12} {

		payload, _ := types.GobMarshal(types.Entity{
			ID:   strconv.Itoa(rand.Intn(1000)),
			Type: types.EntityType(v),
		})

		stream.Send(&pb.EntityResponse{Payload: payload})
	}
	return nil
}

func main() {
	// port := 10001
	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	server := &Backend{}
	// var opts []grpc.ServerOption
	// grpcServer := grpc.NewServer()
	// pb.RegisterBackendServer(grpcServer, server)
	// grpcServer.Serve(lis)

	gs := grpc.NewServer()
	pb.RegisterBackendServer(gs, server)
	wrappedServer := grpcweb.WrapServer(gs,
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		// Redirect gRPC and gRPC-Web requests to the gRPC-Web Websocket Proxy server
		log.Println("Trolo:", req.ProtoMajor, wrappedServer.IsAcceptableGrpcCorsRequest(req), websocket.IsWebSocketUpgrade(req))
		if req.ProtoMajor == 2 && strings.Contains(req.Header.Get("Content-Type"), "application/grpc") ||
			websocket.IsWebSocketUpgrade(req) {
			log.Println("In here!")
			wrappedServer.ServeHTTP(resp, req)
		} else {
			// Serve the GopherJS client
			log.Println("Serve files!", req)
			folderReader(gzipped.FileServer(bundle.Assets)).ServeHTTP(resp, req)
		}
	}

	addr := "localhost:10001"
	httpsSrv := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handler),
	}

	log.Println("Serving on https://" + addr)
	log.Fatal(httpsSrv.ListenAndServe())
}

func folderReader(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/") {
			// Use contents of index.html for directory, if present.
			req.URL.Path = path.Join(req.URL.Path, "index.html")
		}
		fn.ServeHTTP(w, req)
	}
}
