package handler

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type BookSave struct {
	Pid  int
	GifMouthUrl string
	GifTongueUrl string
	Uid  int
	MouthDesc  string
	TongueDesc string
}
type Book struct {
	Pid    int    `json:"pid" xml:"pid" form:"pid" query:"pid"`
	GifUrl string `json:"gif_url" xml:"gif_url" form:"gif_url" query:"gif_url"`
	Uid    int    `json:"uid" xml:"uid" form:"uid" query:"uid"`
	Text   string `json:"text" xml:"text" form:"text" query:"text"`
}

func BookshelfHandler(c *fiber.Ctx) error {

	fmt.Print("BookshelfHandler called\n")
	cred_file_path := "C:\\workspace\\API-Server\\API-Server\\secret\\creds.json"

	// Parse request body
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	// Make BookSave struct
	booksave := BookSave{
		Pid: book.Pid,
		GifMouthUrl: "",
		GifTongueUrl: "",
		Uid: book.Uid,
		MouthDesc: "",
		TongueDesc: "",
	}

	fmt.Printf("BookshelfHandler: %v", book)
	
	// Generate Mouth and Tongue Description
	for _, char := range book.Text {
		fmt.Printf("%c", char)
		PronounceChar := TextToPronounce(string(char))

		fmt.Printf("PronounceChar: %v \n",  PronounceChar)
		MouthDescription := GenerateDescMouth(PronounceChar)
		word := fmt.Sprintf("%s\n", string(char))
		booksave.MouthDesc += word
		booksave.MouthDesc += MouthDescription
		TongueDescription := GenerateDescTongue(PronounceChar)

		booksave.TongueDesc += word
		booksave.TongueDesc += TongueDescription
	}


	// Generate GIF
	PronounceText := TextToPronounce(book.Text)
	// Use only 모음
	count := 0
	MouthText := ""
	TongueText := ""
	for _, char := range PronounceText {
		if count % 3 == 1 {
			MouthText += string(char)
		} else {
			if string(char) != "*" {
				TongueText += string(char)
			}
		}
		count++
	}
	fmt.Printf("MouthText: %v \n", MouthText)
	fmt.Printf("TongueText: %v \n", TongueText)

	gif_mouth_path :=GenerateGIF(TranslateTextToMouthPath(MouthText), "mouth")
	gif_tongue_path :=GenerateGIF(TranslateTextToTonguePath(TongueText), "tongue")

	fmt.Printf("gif_mouth_path: %v \n", gif_mouth_path)
	fmt.Printf("gif_tongue_path: %v \n", gif_tongue_path)

	firebase_gif_mouth_path := UpLoadGIFHandler("mouth", gif_mouth_path)
	firebase_gif_tongue_path := UpLoadGIFHandler("tongue", gif_tongue_path)

	booksave.GifMouthUrl = firebase_gif_mouth_path
	booksave.GifTongueUrl = firebase_gif_tongue_path


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

	_, _, err = client.Collection("problems").Add(context.Background(), booksave)
	if err != nil {
		log.Printf("Error adding book to Firestore: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book added to shelf",
		"book":    booksave,
	})
}
