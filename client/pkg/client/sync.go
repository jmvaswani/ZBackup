package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	uploadpb "github.com/jmvaswani/zbackup/client/pkg/proto"
	"github.com/jmvaswani/zbackup/common/constants"
	"github.com/jmvaswani/zbackup/common/utils"
)

func (s *ClientService) SyncCacheWithServer() error {
	serverResponse, err := s.client.GetMetaDataMap(context.Background(), &uploadpb.GetMetaDataMapRequest{})
	if err != nil {
		log.Fatalf("Failed to get metaData Map from server -> %s", err)
	}
	// TODO Jai: Check if this needs to be a sync.map or just a map
	serverMetaDataMap := map[string]utils.FileMetaData{}
	for filePath, metaData := range serverResponse.MetaDataMap {
		time, err := time.Parse(time.UnixDate, metaData.LastModified)
		if err != nil {
			log.Fatalf("Failed to parse server LastModified time -> %s", err)
		}
		metaData := utils.FileMetaData{
			CheckSum:         metaData.FileCheckSum,
			LastModifiedTime: time,
		}
		serverMetaDataMap[filePath] = metaData
	}

	for fileName, metadata := range serverMetaDataMap {
		log.Printf("Recieved following file from server : %s : Checksum: %s : LastModified %s \n", fileName, metadata.CheckSum, metadata.LastModifiedTime)
	}

	//Loop over client list to see matching files and sync them accordingly
	s.metaDataCache.Range(func(fileLocation, metaData any) bool {
		metadata, ok := metaData.(utils.FileMetaData)
		if !ok {
			log.Fatalf("Error occured while parsing file metadata : %s", fileLocation)
		}
		filelocation, ok := fileLocation.(string)
		if !ok {
			log.Fatalf("Error occured while parsing file metadata : %s", fileLocation)
		}

		// While checking the location, we need to make it "machine agnostic", so we should trim the client base path from the fileName.
		clientBaseDirectory := os.Getenv(constants.ClientDirectoryEnvVariable)
		agnosticLocation := strings.Replace(filelocation, clientBaseDirectory+"/", "", 1)
		serverMetaData, ok := serverMetaDataMap[agnosticLocation]
		if !ok {
			// File does not exist on server, need to upload it
			log.Printf("Uploading file %s to server as it does not exist \n", agnosticLocation)
		} else {
			if serverMetaData.CheckSum != metadata.CheckSum {
				log.Printf("Metadata mismatch for file %s!\n", agnosticLocation)
				// Compare Modified Timestamps to check if an upload/download needs to be done

				if metadata.LastModifiedTime.Compare(serverMetaData.LastModifiedTime) > 0 {
					// Client has a more recent file
					log.Printf("Initiating upload of file %s to server", agnosticLocation)

				} else if metadata.LastModifiedTime.Compare(serverMetaData.LastModifiedTime) < 0 {
					// Server has a more recent file
					log.Printf("Initiating download of file %s from server", agnosticLocation)
					s.DownloadFile(context.Background(), func() { fmt.Printf("Cancel called for %s", agnosticLocation) }, agnosticLocation)

				} else {
					log.Panicf("2 differnent checksums for file %s with the same timestamp, this is impossible", agnosticLocation)
				}

			} else {

				log.Printf("File checksum matched for file %s. Nothing to do", agnosticLocation)
			}
			delete(serverMetaDataMap, agnosticLocation)
		}

		// log.Printf("Local Cache : %s : Checksum: %s : LastModified %s \n", fileName, metadata.CheckSum, metadata.LastModifiedTime)
		return true
	})

	//We have removed all entries that existed on the client. Now if anything is remaining, those are the New files on the server.
	for serverAgnosticLocation, _ := range serverMetaDataMap {
		log.Printf("New file %s found on server: Need to Download", serverAgnosticLocation)
		s.DownloadFile(context.Background(), func() { fmt.Printf("Cancel called for %s", serverAgnosticLocation) }, serverAgnosticLocation)
	}

	return nil
}
