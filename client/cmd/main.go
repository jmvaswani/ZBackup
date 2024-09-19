package main

import (
	"context"
	"log"

	"github.com/fsnotify/fsnotify"
	upload "github.com/jmvaswani/zbackup/client/pkg/upload"
)

func main() {
    // Create connection to GRPC server
    serverAddr := "127.0.0.1:5000"
    client := upload.NewFileUploadClient(serverAddr,10,"/home/jai/Desktop/Work/ZBackup/data")
    err := client.Connect()
    log.Println("Connecting to GRPC server")
    if err != nil{
    log.Fatalf("Failed to connect to GRPC server")
    }
    


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
                    err = client.UploadFile(context.Background(),func() {log.Println("Cancel event called ")},event.Name)
                    if err != nil{
                        log.Printf("Error uploading file -> %s",err.Error())
                    }
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add("/home/jai/Desktop/Work/ZBackup/data")
    if err != nil {
        log.Fatal(err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}