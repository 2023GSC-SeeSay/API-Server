package utils

import (
	"fmt"

	// "log"

	"os"

	"github.com/joho/godotenv"
)

func GetCredentialFilePath() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)	
	}
	gcp_cred_file_path := os.Getenv("GCP_CRED_FILE_PATH")
	fmt.Println(gcp_cred_file_path)
	return gcp_cred_file_path
}