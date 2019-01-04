package main

import (
	"context"
	"log"
	"strings"

	// "github.com/golang/protobuf/ptypes/empty"
	"github.com/johanbrandhorst/protobuf/ptypes/empty"
	"github.com/kyeett/grpc-experiments/proto/client"
	"honnef.co/go/js/dom"
)

// Build this snippet with GopherJS, minimize the output and
// write it to html/frontend.js.
//go:generate gopherjs build frontend.go -m -o html/frontend.js

// Zopfli compress static files.
// //go:generate find ./html/ -name *.gz -prune -o -type f -exec go-zopfli {} +

// Integrate generated JS into a Go file for static loading.
//go:generate bash -c "go run assets_generate.go"

// This constant is very useful for interacting with
// the DOM dynamically
var document = dom.GetWindow().Document().(dom.HTMLDocument)

// Define no-op main since it doesn't run when we want it to
func main() {}

// Ensure our setup() gets called as soon as the DOM has loaded
func init() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go setup()
	})
}

// Setup is where we do the real work.
func setup() {
	// This is the address to the server, and should be used
	// when creating clients.
	serverAddr := strings.TrimSuffix(document.BaseURI(), "/")
	log.Println(serverAddr)
	serverAddr = "http://localhost:10001"

	// TODO: Use functions exposed by generated interface
	cli := client.NewBackendClient(serverAddr)
	e := empty.Empty{}
	resp, err := cli.NewPlayer(context.Background(), &e)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ID:", resp.GetID())

	// // TODO: Use functions exposed by generated interface
	// cli := client.NewBackendClient(serverAddr)

	// button := document.GetElementByID("button").(*dom.HTMLButtonElement)
	// container := document.GetElementByID("container").(*dom.HTMLDivElement)

	// button.AddEventListener("click", false, func(_ dom.Event) {
	// 	go func() {

	// 		user, err := cli.GetUser(context.Background(), &client.UserRequest{
	// 			Id: "4321",
	// 		})

	// 		if err != nil {
	// 			container.SetTextContent("NOES!")
	// 			return
	// 		}
	// 		container.SetTextContent("Name" + user.GetName() + ", Age" + strconv.Itoa(int(user.GetAge())))
	// 	}()
	// })
	// cli := pb.NewBackendClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// e := empty.Empty{}
	// playerID, err := cli.NewPlayer(ctx, &e)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("ID:", playerID.GetID())

	// button := document.GetElementByID("button").(*dom.HTMLButtonElement)

	// button.AddEventListener("click", false, func(_ dom.Event) {
	// 	go func() {

	// 		user, _ := cli.GetMsg(context.Background(), &client.MsgRequest{Id: "Get me"})
	// 		fmt.Println(user.GetText())
	// 	}()
	// })
}
