package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	uploadpb "github.com/jmvaswani/zbackup/client/pkg/proto"
	"github.com/jmvaswani/zbackup/common/constants"
	file "github.com/jmvaswani/zbackup/common/file"
)

func (s *ClientService) DownloadFile(ctx context.Context, cancel context.CancelFunc, agnosticFilePath string) error {

	fmt.Printf("About to download file %s", agnosticFilePath)
	//TODO Jai Remove this:
	clientBaseDirectory := os.Getenv(constants.ClientDirectoryEnvVariable)
	clientFilePath := path.Join(clientBaseDirectory, agnosticFilePath)

	stream, err := s.client.Download(ctx, &uploadpb.FileDownloadRequest{AgnosticFilePath: agnosticFilePath})
	if err != nil {
		return err
	}

	file := file.NewFile()
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			log.Printf("Error -> %s \n", err.Error())

			// g.l.Error(err)
		}
	}()

	for {
		req, err := stream.Recv()

		if file.FilePath == "" {
			directoryPath := filepath.Dir(clientFilePath)
			fileName := filepath.Base(clientFilePath)
			file.SetFile(fileName, directoryPath)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error -> %s \n", err.Error())
			return err
			// return g.logError(status.Error(codes.Internal, err.Error()))
		}
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		log.Printf("received a chunk with size: %d", fileSize)
		// g.l.Debug("received a chunk with size: %d", fileSize)
		if err := file.Write(chunk); err != nil {

			log.Printf("Error -> %s \n", err.Error())
			return err
			// return g.logError(status.Error(codes.Internal, err.Error()))
		}
	}
	// fileName := filepath.Base(file.FilePath)
	log.Printf("saved file: %s, size: %d", file.FilePath, fileSize)
	stream.CloseSend()
	return nil
}
