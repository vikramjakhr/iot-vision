package services

import (
	"log"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"os"
	"io"
	"fmt"
)

func SaveToCloudStorage(file, object string, url chan<- string) {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	url <- write(client, "ioi-vision", file, object)
}

func write(client *storage.Client, bucket, file, object string) string {
	ctx := context.Background()
	// [START upload_file]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("##################### Error while opening ")
		fmt.Println(err)
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Println("##################### Error while copying ")
		fmt.Println(err)
	}
	if err := wc.Close(); err != nil {
		fmt.Println("##################### Error while close ")
		fmt.Println(err)
	}
	acl := client.Bucket(bucket).Object(object).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		fmt.Println("##################### Error while permission ")
		fmt.Println(err)
	}
	client.Close()
	// [END upload_file]
	return "https://storage.googleapis.com/ioi-vision/" + object
}
