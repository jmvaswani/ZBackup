package file

import (
	"log"
	"os"
	"path/filepath"
)

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
