package handler

import (
	"context"
	"fmt"
	"io/ioutil"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

func UploadGIFHandler(c *fiber.Ctx) error {
	fmt.Print("UploadGIFHandler called\t|")
	cred_file_path := "secret\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"
	
	// get the file from the request body
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("Failed to get file: %v", err)
		return err
	}
	
	ctx := context.Background()
	config := &firebase.Config{ProjectID: "seesay"}
	app, err := firebase.NewApp(ctx, config, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		fmt.Printf("Failed to create app: %v", err)
		return err
	}

	client, err := app.Storage(ctx)
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return err
	}

	bucket, err := client.Bucket("seesay.appspot.com")
	if err != nil {
		fmt.Printf("Failed to create bucket: %v", err)
		return err
	}

	// create the file path in the bucket
	gifName := file.Filename
	gif_path := fmt.Sprintf("gif/%s", gifName)
	
	// open a write stream to the bucket
	wc := bucket.Object(gif_path).NewWriter(ctx)

	// set the content type of the file
	wc.ContentType = "image/gif"
	
	// open the file
	fileReader, err := file.Open()
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return err
	}
	defer fileReader.Close()

	// read the file data from the file reader
	fileBytes, err := ioutil.ReadAll(fileReader)
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		return err
	}

	// write the file data to the bucket
	_, err = wc.Write(fileBytes)
	if err != nil {
		fmt.Printf("Failed to write file: %v", err)
		return err
	}
	
	// close the write stream
	err = wc.Close()
	if err != nil {
		fmt.Printf("Failed to close write stream: %v", err)
		return err
	}

	// return the URL of the uploaded file
	url := fmt.Sprintf("https://storage.googleapis.com/seesay.appspot.com/%s", gif_path)
	return c.JSON(fiber.Map{"url": url})
}
