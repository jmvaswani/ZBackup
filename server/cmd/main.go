package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"

	"github.com/jmvaswani/zbackup/common/constants"
	uploadpb "github.com/jmvaswani/zbackup/server/pkg/proto"
	server "github.com/jmvaswani/zbackup/server/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	port := 5000
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	storageLocation := os.Getenv(constants.ServerDirectoryEnvVariable)
	certificateLocation := os.Getenv(constants.CertificateDirectoryEnvVariable)
	serverCrtLocation := path.Join(certificateLocation, "server.crt")
	serverKeyLocation := path.Join(certificateLocation, "server.key")

	creds, err := credentials.NewServerTLSFromFile(serverCrtLocation, serverKeyLocation)
	if err != nil {
		log.Fatalf("Failed to register server certificates -> %s", err.Error())
	}

	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)
	uploadpb.RegisterFileServiceServer(grpcServer, server.NewFileUploadServer(storageLocation))
	log.Println("Server starting to listen")
	grpcServer.Serve(lis)
	log.Println("Server done listening")
}
