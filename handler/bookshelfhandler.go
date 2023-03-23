package handler

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type Book struct {
	Pid    int    `json:"pid"`
	GifUrl string `json:"gif_url"`
	Uid    int    `json:"uid"`
	Text   string `json:"text"`
}

func BookshelfHandler(c *fiber.Ctx) error {
	fmt.Print("BookshelfHandler called\n")
	cred_file_path := "secret\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"

	// Parse request body
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// // Generate GIF
	// gifUrl, err := gifGenerator(book.Text)
	// if err != nil {
	// 	log.Printf("Error generating GIF: %v", err)
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Internal server error",
	// 	})
	// }
	gifUrl := "https://storage.googleapis.com/seesay.appspot.com/gif/test.gif"

	// Update book with GIF URL
	book.GifUrl = gifUrl

	// Add book to Firestore
	client, err := firestore.NewClient(context.Background(), "seesay", option.WithCredentialsFile(cred_file_path))
	if err != nil {
		log.Printf("Error creating Firestore client: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	defer client.Close()

	_, _, err = client.Collection("problems").Add(context.Background(), book)
	if err != nil {
		log.Printf("Error adding book to Firestore: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book added to shelf",
		"book":    book,
	})
}
