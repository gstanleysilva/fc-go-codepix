run:
	@go run main.go all 

run-kafka:
	@go run main.go kafka 

run-grpc:
	@go run main.go grpc -p 50051

proto:
	@protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/proto application/grpc/proto/*.proto

evans:
	@evans -r repl