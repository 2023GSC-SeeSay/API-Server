
// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
        "context"
        "fmt"
        // "log"
		"io/ioutil"

    	"cloud.google.com/go/speech/apiv1"	
        speechpb "cloud.google.com/go/speech/apiv1/speechpb"
		
		"google.golang.org/api/option"
)

func main() {
		credFilePath := "C:\\workspace\\API-Server\\auth\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"
        // Creates a client.
		// Create a new context and load credentials from file
		ctx := context.Background()
		clientOptions := []option.ClientOption{
			option.WithCredentialsFile(credFilePath),
		}
		speechClient, err := speech.NewClient(ctx, clientOptions...)

		if err != nil {	
			fmt.Printf("Error creating SpeechClient: %v\n", err)
			return
		}
		defer speechClient.Close()

        // The path to the remote audio file to transcribe.
        // fileURI := "gs://cloud-samples-data/speech/brooklyn_bridge.raw"
		filePath :="C:\\Users\\aa\\OneDrive - 고려대학교\\문서\\소리 녹음\\녹음 (4).m4a"
		audioData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading audio file: %v\n", err)
			return
		}
		config := &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
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
		if err != nil {
			fmt.Printf("Error transcribing audio: %v\n", err)
			return
		}
	
		// Print out the transcribed text
		for _, result := range resp.Results {
			for _, alt := range result.Alternatives {
				fmt.Printf("Transcript: %v\n", alt.Transcript)
			}
		}
}