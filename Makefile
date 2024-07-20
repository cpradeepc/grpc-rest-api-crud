del:
	rm -r ./protofile/*.go
gen:
	protoc --go_out=. --go-grpc_out=. protofile/*.proto

run-srv:
	go run ./server/server.go
run-clnt:
	go run ./client/client.go

