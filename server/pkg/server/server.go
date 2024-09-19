package server

import (
	// "bytes"
	"io"
	"log"
	"os"
	"path/filepath"

	// config "github.com/dimk00z/grpc-filetransfer/config/server"
	// "github.com/dimk00z/grpc-filetransfer/pkg/logger"
	uploadpb "github.com/jmvaswani/zbackup/server/pkg/proto"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

func NewFileUploadServer(storageLocation string) FileServiceServer {
	return FileServiceServer{
        storageLocation:  storageLocation,
    }
}

type FileServiceServer struct {
	uploadpb.UnimplementedFileServiceServer
	// l   *logger.Logger
	// cfg *config.Config
    storageLocation string
}


// mustEmbedUnimplementedFileServiceServer implements uploadpb.FileServiceServer.
// Subtle: this method shadows the method (UnimplementedFileServiceServer).mustEmbedUnimplementedFileServiceServer of FileServiceServer.UnimplementedFileServiceServer.
func (g FileServiceServer) mustEmbedUnimplementedFileServiceServer() {
	panic("unimplemented")
}

func (g FileServiceServer) Upload(stream uploadpb.FileService_UploadServer) error {
	file := NewFile()
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			log.Printf("Error -> %s \n",err.Error())
            
			// g.l.Error(err)
		}
        }()
        for {
            req, err := stream.Recv()
            if file.FilePath == "" {
                file.SetFile(req.GetFileName(), g.storageLocation)
            }
            if err == io.EOF {
                break
            }
            if err != nil {
            log.Printf("Error -> %s \n",err.Error())
            return err
			// return g.logError(status.Error(codes.Internal, err.Error()))
		}
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		log.Printf("received a chunk with size: %d", fileSize)
		// g.l.Debug("received a chunk with size: %d", fileSize)
		if err := file.Write(chunk); err != nil {
            
            log.Printf("Error -> %s \n",err.Error())
            return err
			// return g.logError(status.Error(codes.Internal, err.Error()))
		}
	}
	fileName := filepath.Base(file.FilePath)
	log.Printf("saved file: %s, size: %d", fileName, fileSize)
	// g.l.Debug("saved file: %s, size: %d", fileName, fileSize)
	return stream.SendAndClose(&uploadpb.FileUploadResponse{FileName: fileName, Size: fileSize})
}

// func New(l *logger.Logger, cfg *config.Config) *FileServiceServer {
// 	return &FileServiceServer{
// 		l:   l,
// 		cfg: cfg,
// 	}
// }

func NewFile() File {
	return File{}
}

type File struct {
	FilePath string
	// buffer     *bytes.Buffer
	OutputFile *os.File
}

func (f *File) SetFile(fileName, path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f.FilePath = filepath.Join(path, fileName)
	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}
	f.OutputFile = file
	return nil
}

func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil {
		return nil
	}
	_, err := f.OutputFile.Write(chunk)
	return err
}

func (f *File) Close() error {
	return f.OutputFile.Close()
}
