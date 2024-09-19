proto-gen:
	protoc --go_out=./client/pkg/ --go_opt=paths=source_relative --go-grpc_out=./client/pkg --go-grpc_opt=paths=source_relative proto/grpc.proto 
	protoc --go_out=./server/pkg/ --go_opt=paths=source_relative --go-grpc_out=./server/pkg --go-grpc_opt=paths=source_relative proto/grpc.proto
