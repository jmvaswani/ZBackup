package main

import (
	"fmt"
	"log"
	"net"

	uploadpb "github.com/jmvaswani/zbackup/server/pkg/proto"
	server "github.com/jmvaswani/zbackup/server/pkg/server"
	"google.golang.org/grpc"
)


func main() {
	port := 5000
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
if err != nil {
  log.Fatalf("failed to listen: %v", err)
}
var opts []grpc.ServerOption
storageLocation := "/home/jai/Desktop/Work/ZBackup/data_server"
grpcServer := grpc.NewServer(opts...)
uploadpb.RegisterFileServiceServer(grpcServer, server.NewFileUploadServer(storageLocation))
log.Println("Server starting to listen")
grpcServer.Serve(lis)
log.Println("Server done listening")
}



