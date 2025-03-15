package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
	client "github.com/jmvaswani/zbackup/client/pkg/client"
	"github.com/jmvaswani/zbackup/client/pkg/timedqueue"
	"github.com/jmvaswani/zbackup/common/constants"
)

func main() {
	// Create connection to GRPC server
	serverAddr := "127.0.0.1:5000"
	client := client.NewFileUploadClient(serverAddr, 10, "/home/jai/Desktop/Work/ZBackup/data")
	err := client.Connect()
	log.Println("Connecting to GRPC server")
	if err != nil {
		log.Fatalf("Failed to connect to GRPC server")
	}
	// Sync caches
	syncErr := client.SyncCacheWithServer()
	if syncErr != nil {
		log.Panicf("Failed to sync data with server, exiting : %s", syncErr)
	}
	queue := timedqueue.NewTimedQueue(time.Second * 3)

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					queue.AddTask(event.Name, client.InitiateFileUpload)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add paths recursively.
	clientBaseDirectory := os.Getenv(constants.ClientDirectoryEnvVariable)
	err = addWatcherPathsRecursive(watcher, clientBaseDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func addWatcherPathsRecursive(watcher *fsnotify.Watcher, baseDirectory string) error {
	log.Printf("Adding watcher for %s", baseDirectory)
	err := watcher.Add(baseDirectory)
	if err != nil {
		return err
	}
	contents, err := os.ReadDir(baseDirectory)
	if err != nil {
		return err
	}

	for _, content := range contents {
		// If file is a directory, we need to set a watcher for it, so recurse
		if content.IsDir() {
			err := addWatcherPathsRecursive(watcher, path.Join(baseDirectory, content.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
