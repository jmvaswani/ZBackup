package upload

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	// "github.com/jmvaswani/zbackup/common/constants"

	uploadpb "github.com/jmvaswani/zbackup/client/pkg/proto"
	"github.com/jmvaswani/zbackup/common/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewFileUploadClient(addr string, batchSize int, baseDirectory string) ClientService {

	checksumCache, err := prepareChecksumMap(baseDirectory)

	if err != nil {
		log.Fatalf("Failed to prepare Checksum Map")
	}
	log.Println("Finished preparing Checksum Cache")
	checksumCache.Range(func(key, value any) bool {
		log.Printf("File : %s, checksum : %s", key, value)
		return true
	})

	return ClientService{
		addr:          addr,
		batchSize:     batchSize,
		checksumCache: checksumCache,
	}

}

type ClientService struct {
	uploadInProgress sync.Mutex
	checksumCache    *sync.Map
	addr             string
	conn             *grpc.ClientConn
	batchSize        int
	client           uploadpb.FileServiceClient
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

func (s *ClientService) UploadFile(ctx context.Context, cancel context.CancelFunc, filePath string) error {
	s.uploadInProgress.Lock()
	defer s.uploadInProgress.Unlock()
	fileChecksum := GetFileChecksum(filePath)
	storedFileChecksum, ok := s.checksumCache.Load(filePath)
	if ok && fileChecksum == storedFileChecksum {
		log.Printf("Duplicate call made for file : %s , Checksum : %s . Ignoring.. \n", filePath, fileChecksum)
		return nil
	}
	log.Printf("Beginning upload of file %s to GRPC server \n", filePath)
	stream, err := s.client.Upload(ctx)

	parts := strings.Split(filePath, "/")
	fileName := parts[len(parts)-1]
	//Upload File to Server

	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	buf := make([]byte, s.batchSize)
	batchNumber := 1
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]

		if err := stream.Send(&uploadpb.FileUploadRequest{FileName: fileName, Chunk: chunk}); err != nil {
			return err
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Sent - %v bytes - %s\n", res.GetSize(), res.GetFileName())

	s.checksumCache.Store(filePath, fileChecksum)
	cancel()
	return nil
}

func prepareChecksumMap(directory string) (*sync.Map, error) {
	var checkSumMap sync.Map

	contents, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, content := range contents {
		// If file is a directory, we need to recurse and get all entries in there
		if content.IsDir() {
			recurseDir := path.Join(directory, content.Name())
			recurseContents, err := prepareChecksumMap(recurseDir)
			if err != nil {
				return nil, err
			}

			recurseContents.Range(func(item, checkSum any) bool {
				checkSumMap.Store(item, checkSum)
				return true
			})
			// for item, checkSum := range recurseContents.Range() {
			// 	checkSumMap[item] = checkSum
			// }
		} else {
			absoluteFilePath := path.Join(directory, content.Name())
			fileHash := GetFileChecksum(absoluteFilePath)
			checkSumMap.Store(absoluteFilePath, fileHash)
			// checkSumMap[absoluteFilePath] = fileHash
		}

	}
	return &checkSumMap, nil

}

func GetFileChecksum(filePath string) (checksum string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	hasher := sha256.New()
	defer f.Close()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}
	checksum = hex.EncodeToString(hasher.Sum(nil))
	return

}
