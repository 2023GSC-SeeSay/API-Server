// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"fmt"

	// "log"
	"io/ioutil"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"

	"google.golang.org/api/option"
)

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
