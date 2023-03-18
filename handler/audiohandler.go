package handler

import (
	"context"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	// "log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"

	utils "api-server/utils"

	"google.golang.org/api/option"
)

const MAX_AUDIO_FILE_SIZE = 1024 * 1024 // 1MB
type Audio struct {
	AudioUrl string `json:"audio_url"`
	Uid string `json:"uid"`
}

func AudioHandler(c *fiber.Ctx) error {
	// flutter 앱에서 오디오파일을 gcp storage에 업로드하고, 그 url을 받아서, stt를 통해 텍스트를 받아온다.

	cred_file_path := utils.GetCredentialFilePath()

	// a := new(Audio)
	// if err := c.BodyParser(a); err != nil {
	// 	return err
	// }
	// TODO : user valid check

	pid := c.Params("pid") // 특정 문제에 대한 오디오인지 확인하기 위해 pid를 받아온다.
	fmt.Print(pid)
	uri := "https://firebasestorage.googleapis.com/v0/b/seesay.appspot.com/o/audio%2F2o2ongerwob2303fewns%2Faudio_1_1.wav?alt=media&token=ec2bbb31-0e57-4de3-9f9d-9a107a815e9f"
	// TODO : check if uid is valid (uid matches with audio url permission)
	script := stt(uri, cred_file_path)
	c.SendString(script)

	return nil
}
func stt(audio_uri string, cred_file_path string) string {
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

		if err != nil {
			fmt.Printf("Error reading audio file: %v\n", err)
			return ""
		}
		// fmt.Printf("Audio data: %v\n", audioData)
		config := &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 44100,
			LanguageCode:    "en-US",
			}
		
			// Create a new RecognitionAudio from the audio data
			audio := &speechpb.RecognitionAudio{
				AudioSource: &speechpb.RecognitionAudio_Uri{Uri: audio_uri},
			}
		
			// Call the SpeechClient's Recognize method to transcribe the audio
		resp, err := speechClient.Recognize(ctx, &speechpb.RecognizeRequest{
			Config: config,
			Audio:  audio,
		})
		fmt.Printf("Response: %v\n", resp)
		// fmt.Printf("Error: %v\n", err)
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
