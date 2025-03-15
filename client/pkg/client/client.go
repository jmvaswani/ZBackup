package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"path"
	"sync"

	uploadpb "github.com/jmvaswani/zbackup/client/pkg/proto"
	"github.com/jmvaswani/zbackup/common/constants"
	"github.com/jmvaswani/zbackup/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientService struct {
	uploadInProgress sync.Mutex
	metaDataCache    *sync.Map
	addr             string
	conn             *grpc.ClientConn
	batchSize        int
	client           uploadpb.FileServiceClient
}

func NewFileUploadClient(addr string, batchSize int, baseDirectory string) ClientService {

	metaDataCache, err := utils.PrepareMetaDataMap(baseDirectory)

	if err != nil {
		log.Fatalf("Failed to prepare Checksum Map")
	}
	log.Println("Finished preparing Checksum Cache")
	metaDataCache.Range(func(key any, value any) bool {
		metadata, ok := value.(utils.FileMetaData)
		if !ok {
			log.Fatalf("Error occured while preparin checksum map")
		}
		log.Printf("File : %s, checksum : %s , LastModified : %s", key, metadata.CheckSum, metadata.LastModifiedTime)
		return true
	})

	return ClientService{
		addr:          addr,
		batchSize:     batchSize,
		metaDataCache: metaDataCache,
	}

}

func (s *ClientService) Connect() error {
	var opts []grpc.DialOption

	certDir := os.Getenv(constants.CertificateDirectoryEnvVariable)
	clientCrtLocation := path.Join(certDir, "client.crt")
	clientKeyLocation := path.Join(certDir, "client.key")
	caLocation := path.Join(certDir, "CA.crt")

	clientKeyPair, err := tls.LoadX509KeyPair(clientCrtLocation, clientKeyLocation)
	if err != nil {
		log.Fatalf("Failed to read Client certs -> %s", err.Error())
	}

	caCertPool := x509.NewCertPool()
	ca, err := os.ReadFile(caLocation)
	if err != nil {
		log.Fatalf("Failed to read CA cert -> %s", err.Error())
	}

	caCertPool.AppendCertsFromPEM(ca)

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{clientKeyPair},
		RootCAs:      caCertPool,
		ServerName:   os.Getenv(constants.ServerNameEnvVariable),
	}
	creds := credentials.NewTLS(&tlsConfig)
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.NewClient(s.addr, opts...)
	s.conn = conn
	s.client = uploadpb.NewFileServiceClient(s.conn)
	return err
}

func (s *ClientService) Close() error {
	return s.conn.Close()
}
