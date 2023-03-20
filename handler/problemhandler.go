package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type Problem struct {
	Mouth		string 	`firestore:"mouth"`		 	// 입은 어떻게어떻게 벌립니다
	MouthUri	string 	`firestore:"mouth_uri"`		// link
	Pid		int 	`firestore:"p_id"` 			// 1
	Title		string 	`firestore:"title"` 		// 기본 단어 2
	Tongue		string 	`firestore:"tongue"` 		// 혀는 어떻게 어떻게 합니다
	TongueUri	string 	`firestore:"tongue_uri"` 	// link
	Type		string  `firestore:"type_"`			// 기본 발음 연습
	Uid		int 	`firestore:"u_id"`			// 0
}

type ProblemList struct {
	ProblemList map[string]interface{} `firestore:"problem_list"` // list of problem names
}

func ProblemHandler(ctx *fiber.Ctx) error {
	fmt.Print("ProblemHandler called\t")
	// cred_file_path := utils.GetCredentialFilePath()
	cred_file_path := "secret\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"
	pidStr := ctx.Params("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error converting pid to int")
	}

	if pid == 0 {
		fmt.Print("p_id == 0: get problem list\n\n")
		context := ctx.Context()

		problemList, err := GetFireStoreProblemList(context, cred_file_path)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Error getting Firestore document")
		}

		jsonData, err := json.Marshal(problemList)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Error converting Firestore data to JSON")
		}


		ctx.Set("Content-Type", "application/json") // set the content type as JSON
		ctx.Set("Access-Control-Allow-Origin", "*") // allows CORS


		// Print the data.
		fmt.Println(problemList)

		return ctx.Status(http.StatusOK).SendString(string(jsonData)) 

	} else {
		fmt.Printf("p_id == %v: get problem #%v\n\n", pid, pid)
		context := ctx.Context()

		problem, err := GetFireStoreProblem(context, pid, cred_file_path)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Error getting Firestore document")
		}

		jsonData, err := json.Marshal(problem)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Error converting Firestore data to JSON")
		}


		ctx.Set("Content-Type", "application/json") // set the content type as JSON
		ctx.Set("Access-Control-Allow-Origin", "*") // allows CORS


		// Print the data.
		fmt.Println(problem)

		return ctx.Status(http.StatusOK).SendString(string(jsonData))  
	}

}

func GetFireStoreProblemList(context context.Context, cred_file_path string) (*ProblemList, error)  {
	projectID := "seesay"
	var pList ProblemList

	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("problems").Doc("problem_list").Get(context)
	if err != nil {
		fmt.Printf("Failed to get document: %v", err)
		return nil, err
	}

	docData := doc.Data()

	pList.ProblemList = docData["problem_list"].(map[string]interface{})

	return &pList, nil
}


func GetFireStoreProblem(context context.Context, pid int, cred_file_path string) (*Problem, error)  {
	projectID := "seesay"
	var problem Problem

	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	query := client.Collection("problems").Where("p_id", "==", pid).Limit(1)

	docs, err := query.Documents(context).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get documents: %v", err)
	}

	if len(docs) > 0 {
		if err := docs[0].DataTo(&problem); err != nil {
			return nil, fmt.Errorf("failed to convert document to Problem struct: %v", err)
		}
	}

	return &problem, nil
}

