package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
    Problems  []string    `firestore:"problems,omitempty"` // list of problem names
}

func ProblemHandler(ctx *fiber.Ctx) error {
	// cred_file_path := utils.GetCredentialFilePath()
	cred_file_path := "secret\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"
	fmt.Print(cred_file_path)
	pid := ctx.Params("pid")
	fmt.Print(pid)

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

	// firebase firestore query로 pid의 정보를 가져옴
	// 가져온 정보를 json으로 변환하여 response로 보냄
	return ctx.Status(http.StatusOK).SendString(string(jsonData))  
}


func GetFireStoreProblem(context context.Context, pidStr string, cred_file_path string) (interface{}, error) {
    // firebase firestore query로 pid의 정보를 가져옴
    // 가져온 정보를 json으로 변환하여 response로 보냄

    // ctx := context.Background()
    projectID := "seesay"
    var problem Problem
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return problem, err
	}

    client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
    if err != nil {
        return problem, err
    }
    defer client.Close()

    if pid == 0 {
        
		doc, err := client.Collection("problems").Doc("problem_list").Get(context)
		if err != nil {
			log.Fatalf("Failed to get problem list document: %v", err)
		}

		var data map[string]interface{}
		if err := doc.DataTo(&data); err != nil {
			return problem, fmt.Errorf("failed to convert problem list to map: %v", err)
		}

		// Extract problem names from map
		var problemList []string
		for _, value := range data {
			if strValue, ok := value.(string); ok {
				problemList = append(problemList, strValue)
			}
		}
	
		problem.Problems = problemList
		problem.Pid = pid
		return problem, nil

    } else {
        // fetch specific problem
        query := client.Collection("problems").Where("p_id", "==", pid).Limit(1)
        docs, err := query.Documents(context).GetAll()
        if err != nil {
            return problem, fmt.Errorf("failed to get documents: %v", err)
        }
        if len(docs) > 0 {
            var p Problem
            if err := docs[0].DataTo(&p); err != nil {
                return problem, fmt.Errorf("failed to convert document to Problem struct: %v", err)
            }
            problem = p
        }
    }

    return problem, nil
}



// func GetFireStoreProblem(context context.Context, pidStr string, cred_file_path string) (Problem, error)  {
// 	// firebase firestore query로 pid의 정보를 가져옴
// 	// 가져온 정보를 json으로 변환하여 response로 보냄

// 	// ctx := context.Background()
// 	projectID := "seesay"
// 	var problem Problem
// 	pid, err := strconv.Atoi(pidStr)
// 	if err != nil {
// 		return problem, err
// 	}

// 	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
// 	if err != nil {
// 		return problem, err
// 	}
// 	defer client.Close()

// 	query := client.Collection("problems").Where("p_id", "==", pid).Limit(1)

// 	docs, err := query.Documents(context).GetAll()
// 	if err != nil {
// 		return problem, fmt.Errorf("failed to get documents: %v", err)
// 	}

// 	if len(docs) > 0 {
// 		if err := docs[0].DataTo(&problem); err != nil {
// 			return problem, fmt.Errorf("failed to convert document to Problem struct: %v", err)
// 		}
// 	}

// 	return problem, nil

// }
