package handler

import (
	"context"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	// "log"
	"io/ioutil"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"

	"google.golang.org/api/option"
)

const MAX_AUDIO_FILE_SIZE = 1024 * 1024 // 1MB

func uploadHandler(audio_file) error { 
	// Post request의 auido file을 받아서, 현재 로그인된 유저 id 정보와, url을 통해 받은 pid를 통해 firestore에 오디오 파일을 저장한다.
	// 현재 로그인된 유저 id 정보와, url을 통해 받은 pid를 통해 firestore에 오디오 파일을 저장한다.

	return nil
}
func AudioHandler(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	file := form.File["audio"][0]

	fmt.Printf("File: %v", file)

	// firebase store 에 업로드

	cred_file_path := "C:\\workspace\\API-Server\\API-Server\\secret\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"
	audio_file_path := "C:\\workspace\\API-Server\\API-Server\\audio_1_1.wav"
	uid := c.Params("uid")
	pid := c.Params("pid")
	fmt.Printf(uid)
	fmt.Printf(pid)
	fmt.Printf(cred_file_path)
	fmt.Printf(audio_file_path)
	ctx := context.Background()
	clientOptions := []option.ClientOption{
		option.WithCredentialsFile(cred_file_path),
	}
	speechClient, err := speech.NewClient(ctx, clientOptions...)

	if err != nil {	
		fmt.Printf("Error creating SpeechClient: %v\n", err)
		return err
	}
	defer speechClient.Close()

	// The path to the remote audio file to transcribe.
	// fileURI := "gs://cloud-samples-data/speech/brooklyn_bridge.raw"
	// filePath :="C:\\workspace\\API-Server\\음성 002.wav"
	audioData, err := ioutil.ReadFile(audio_file_path)
	if err != nil {
		fmt.Printf("Error reading audio file: %v\n", err)
		return err
	}
	// fmt.Printf("Audio data: %v\n", audioData)
	config := &speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_LINEAR16,
		SampleRateHertz: 44100,
		LanguageCode:    "en-US",
		}
	
		// Create a new RecognitionAudio from the audio data
		audio := &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
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
		return err
	}

	// Print out the transcribed text
	// for _, result := range resp.Results {
	// 	for _, alt := range result.Alternatives {
	// 		fmt.Printf("Transcript: %v\n", alt.Transcript)
	// 	}
	// }
	c.SendString(resp.Results[0].Alternatives[0].Transcript)
	return nil
}
func stt(audio_file_path string, cred_file_path string) string {
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

        // The path to the remote audio file to transcribe.
        // fileURI := "gs://cloud-samples-data/speech/brooklyn_bridge.raw"
		// filePath :="C:\\workspace\\API-Server\\음성 002.wav"
		audioData, err := ioutil.ReadFile(audio_file_path)
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
				AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
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
