package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"

	// "log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"

	"google.golang.org/api/option"
)

const MAX_AUDIO_FILE_SIZE = 1024 * 1024 // 1MB
type Audio struct {
	AudioUrl string `json:"audio_url"`
	Uid      string `json:"uid"`
}

func AudioHandler(c *fiber.Ctx) error {
	// flutter 앱에서 오디오파일을 gcp storage에 업로드하고, 그 url을 받아서, stt를 통해 텍스트를 받아온다.
	fmt.Print("AudioHandler called\n")
	// cred_file_path := utils.GetCredentialFilePath()
	cred_file_path := "credentials.json"
	// ctx := context.Background()

	// a := new(Audio)
	// if err := c.BodyParser(a); err != nil {
	// 	return err
	// }
	// TODO : user valid checkuidStr := context.Params("uid")

	uidStr := c.Params("uid")
	pidStr := c.Params("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error converting uid to int")
	}

	fmt.Printf("uid: %v, pid: %v\n", uidStr, pid)

	// client, err := storage.NewClient(ctx, option.WithCredentialsFile(cred_file_path))
	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).SendString("Error creating storage client")
	// }

	// // Get a handle to the bucket
	// bucket := client.Bucket("seesay.appspot.com")

	// // Get a handle to the object
	// objectName := fmt.Sprintf("audio/uid:%s/pid:%d/%s", uidStr, pid, audioName)

	// objectAttrs, err := bucket.Object(objectName).Attrs(ctx)
	// if err != nil {
	// 	log.Fatalf("Error getting object attributes: %v", err)
	// }
	// audioURL, err := storage.SignedURL(objectAttrs.Bucket, objectAttrs.Name, &storage.SignedURLOptions{
	// 	GoogleAccessID: "seesay@appspot.gserviceaccount.com",
	// 	Method:         http.MethodGet,
	// 	Expires:        time.Now().Add(time.Hour),
	// })
	// if err != nil {
	// 	log.Fatalf("Error creating signed URL: %v", err)
	// }

	audioURL := fmt.Sprintf("gs://seesay.appspot.com/audio/uid:%s/pid:%d/audio.wav", uidStr, pid)

	fmt.Printf("Audio URL: %s\n", audioURL)

	// TODO : check if uid is valid (uid matches with audio url permission)
	script := stt(audioURL, cred_file_path)
	c.SendString(script)

	return nil
}
func stt(audioURL string, cred_file_path string) string {
	/* Transcribe the given audio file using google cloud speech api.*/
	/* Requires audio file path & credential file path */
	credFilePath := cred_file_path
	// Creates a client.
	// Create a new context and load credentials from file
	ctx := context.Background()
	clientOptions := []option.ClientOption{
		option.WithCredentialsFile(credFilePath),
	}
	speechClient, err := speech.NewClient(ctx, clientOptions...)

	if err != nil {
		fmt.Printf("Error creating SpeechClient: %v\n", err)
		return ""
	}
	defer speechClient.Close()

	// fmt.Printf("Audio data: %v\n", audioData)
	config := &speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_LINEAR16,
		SampleRateHertz: 44100,
		LanguageCode:    "ko-KR",
	}

	// Create a new RecognitionAudio from the audio data
	audio := &speechpb.RecognitionAudio{
		AudioSource: &speechpb.RecognitionAudio_Uri{Uri: audioURL},
	}

	// Call the SpeechClient's Recognize method to transcribe the audio
	resp, err := speechClient.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: config,
		Audio:  audio,
	})
	fmt.Printf("Response: %v\n", resp)

	if err != nil {
		fmt.Printf("Error transcribing audio: %v\n", err)
		return ""
	}

	// Print out the transcribed text
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcript: %v\n", alt.Transcript)
		}
	}
	return resp.Results[0].Alternatives[0].Transcript
}
