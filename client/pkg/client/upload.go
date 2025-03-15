package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	uploadpb "github.com/jmvaswani/zbackup/client/pkg/proto"
	"github.com/jmvaswani/zbackup/common/constants"
	"github.com/jmvaswani/zbackup/common/utils"
)

func (s *ClientService) InitiateFileUpload(filePath string) error {
	return s.UploadFile(context.Background(), func() { fmt.Printf("Cancel called for %s", filePath) }, filePath)
}

func (s *ClientService) UploadFile(ctx context.Context, cancel context.CancelFunc, filePath string) error {
	s.uploadInProgress.Lock()
	defer s.uploadInProgress.Unlock()
	fileMetadata := utils.GetFileMetadata(filePath)
	storedFileMetadata, ok := s.metaDataCache.Load(filePath)
	if ok && fileMetadata == storedFileMetadata {
		log.Printf("Duplicate call made for file : %s , Checksum : %s . Ignoring.. \n", filePath, fileMetadata)
		return nil
	}
	log.Printf("Beginning upload of file %s to GRPC server \n", filePath)
	stream, err := s.client.Upload(ctx)

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
		clientBaseDirectory := os.Getenv(constants.ClientDirectoryEnvVariable)
		agnosticFilePath := strings.Replace(filePath, clientBaseDirectory+"/", "", 1)
		if err := stream.Send(&uploadpb.FileUploadRequest{AgnosticFilePath: agnosticFilePath, Chunk: chunk}); err != nil {
			return err
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Sent - %v bytes - %s\n", res.GetSize(), res.GetAgnosticFilePath())

	s.metaDataCache.Store(filePath, fileMetadata)
	cancel()
	return nil
}
