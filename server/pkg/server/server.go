package server

import (
	// "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jmvaswani/zbackup/common/constants"
	file "github.com/jmvaswani/zbackup/common/file"
	"github.com/jmvaswani/zbackup/common/utils"
	uploadpb "github.com/jmvaswani/zbackup/server/pkg/proto"
	"google.golang.org/grpc"
)

func NewFileUploadServer(storageLocation string) FileServiceServer {
	metaDataCache, err := utils.PrepareMetaDataMap(storageLocation)

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

	return FileServiceServer{
		storageLocation: storageLocation,
		metaDataCache:   metaDataCache,
	}
}

type FileServiceServer struct {
	uploadpb.UnimplementedFileServiceServer
	// l   *logger.Logger
	// cfg *config.Config
	storageLocation string
	metaDataCache   *sync.Map
}

// mustEmbedUnimplementedFileServiceServer implements uploadpb.FileServiceServer.
// Subtle: this method shadows the method (UnimplementedFileServiceServer).mustEmbedUnimplementedFileServiceServer of FileServiceServer.UnimplementedFileServiceServer.
func (g FileServiceServer) mustEmbedUnimplementedFileServiceServer() {
	panic("unimplemented")
}

func (g FileServiceServer) Upload(stream uploadpb.FileService_UploadServer) error {
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
			agnosticPath := req.GetAgnosticFilePath()
			directoryPath := filepath.Dir(agnosticPath)
			fileName := filepath.Base(agnosticPath)
			file.SetFile(fileName, path.Join(g.storageLocation, directoryPath))
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
	// g.l.Debug("saved file: %s, size: %d", fileName, fileSize)
	return stream.SendAndClose(&uploadpb.FileUploadResponse{AgnosticFilePath: file.FilePath, Size: fileSize})
}

func (g FileServiceServer) Download(req *uploadpb.FileDownloadRequest, stream grpc.ServerStreamingServer[uploadpb.FileDownloadResponse]) error {
	agnosticFilePath := req.GetAgnosticFilePath()
	realFilePath := path.Join(g.storageLocation, agnosticFilePath)

	if _, err := os.Stat(realFilePath); err == nil {
		// File exists
		file, err := os.Open(realFilePath)
		if err != nil {
			return errors.New(fmt.Sprintf("Error opening file %s on server : %s", agnosticFilePath, err))
		}
		//TODO Jai fix this constant batchsize value
		buf := make([]byte, 10)
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
			if err := stream.Send(&uploadpb.FileDownloadResponse{Chunk: chunk}); err != nil {
				return err
			}
			log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
			batchNumber += 1

		}

	} else if errors.Is(err, os.ErrNotExist) {
		return errors.New(fmt.Sprintf("File %s not found on server", agnosticFilePath))
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		fmt.Printf("Unexpected error occureddd while looking for file %s : %s", realFilePath, err)

	}
	return nil
}

func (g FileServiceServer) GetMetaDataMap(context.Context, *uploadpb.GetMetaDataMapRequest) (*uploadpb.GetMetaDataMapResponse, error) {

	responseMap := map[string]*uploadpb.FileMetaData{}

	g.metaDataCache.Range(func(key any, value any) bool {
		metadata, ok := value.(utils.FileMetaData)
		if !ok {
			// Jai fix this error
			log.Fatalf("Error occured while preparin checksum map")
		}
		fileLocation, ok := key.(string)
		if !ok {
			// Jai fix this error
			log.Fatalf("Error occured while preparin checksum map")
		}

		// While checking the server location, we need to make it "machine agnostic", so we should trim the server base path from the map.
		serverBaseDirectory := os.Getenv(constants.ServerDirectoryEnvVariable)
		agnosticLocation := strings.Replace(fileLocation, serverBaseDirectory+"/", "", 1)
		responseMap[agnosticLocation] = &uploadpb.FileMetaData{
			FileCheckSum: metadata.CheckSum,
			LastModified: metadata.LastModifiedTime.Format(time.UnixDate),
		}

		return true
	})
	return &uploadpb.GetMetaDataMapResponse{MetaDataMap: responseMap}, nil
}
