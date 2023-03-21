package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)



func GIFHandler(c *fiber.Ctx) error {
	fmt.Print("GIFHandler called\t|")
	cred_file_path := "secret\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"
	gif_path := c.Params("gif_path")
	
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

	gifRef := bucket.Object(gif_path)

	gifReader, err := gifRef.NewReader(ctx)
	if err != nil {
		fmt.Printf("Failed to read GIF: %v", err)
		return err
	}
	defer gifReader.Close()

	gifBytes, err := ioutil.ReadAll(gifReader)
	if err != nil {
		fmt.Printf("Failed to get GIF: %v", err)
		return err
	}


	c.Response().Header.Set("Content-Type", "image/gif")
	c.Response().Header.Set("Content-Length", strconv.Itoa(len(gifBytes)))
	c.Write(gifBytes)
	
	return nil
}
