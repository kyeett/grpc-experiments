proto:
	protoc -I. -Ivendor/ ./proto/game.proto \
		--gopherjs_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/johanbrandhorst/protobuf/ptypes/empty:$$GOPATH/src \
		--go_out=plugins=grpc:$$GOPATH/src
	go generate ./frontend/

clean:
	rm -f ./proto/client/* ./proto/server/* ./cert.pem ./key.pem \
		./frontend/html/frontend.js ./frontend/html/frontend.js.map


install:
	go install ./vendor/github.com/golang/protobuf/protoc-gen-go \
		./vendor/github.com/johanbrandhorst/protobuf/protoc-gen-gopherjs \
		./vendor/github.com/foobaz/go-zopfli \
		./vendor/github.com/gopherjs/gopherjs

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	proto \
	test \
	testrace \
	clean \
	coverage