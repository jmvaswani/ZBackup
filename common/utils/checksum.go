package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

type FileMetaData struct {
	CheckSum         string
	LastModifiedTime time.Time
}

func GetFileMetadata(filePath string) (metadata FileMetaData) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	hasher := sha256.New()

	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}
	metadata.CheckSum = hex.EncodeToString(hasher.Sum(nil))

	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	metadata.LastModifiedTime = stat.ModTime()

	return metadata
}

func PrepareMetaDataMap(directory string) (metaDataMap *sync.Map, err error) {
	metaDataMap = &sync.Map{}
	contents, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, content := range contents {
		// If file is a directory, we need to recurse and get all entries in there
		if content.IsDir() {
			recurseDir := path.Join(directory, content.Name())
			recurseContents, err := PrepareMetaDataMap(recurseDir)
			if err != nil {
				return nil, err
			}

			recurseContents.Range(func(item, checkSum any) bool {
				metaDataMap.Store(item, checkSum)
				return true
			})
		} else {
			absoluteFilePath := path.Join(directory, content.Name())
			fileMetaData := GetFileMetadata(absoluteFilePath)
			metaDataMap.Store(absoluteFilePath, fileMetaData)
		}

	}
	return metaDataMap, nil

}
