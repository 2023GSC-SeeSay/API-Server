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
	Mouth     string `firestore:"mouth"`      // 입은 어떻게어떻게 벌립니다
	MouthUri  string `firestore:"mouth_uri"`  // link
	Pid       int    `firestore:"p_id"`       // 1
	Title     string `firestore:"title"`      // 기본 단어 2
	Tongue    string `firestore:"tongue"`     // 혀는 어떻게 어떻게 합니다
	TongueUri string `firestore:"tongue_uri"` // link
	Type      string `firestore:"type_"`      // 기본 발음 연습
	Uid       int    `firestore:"u_id"`       // 0
}

type ProblemList struct {
	ProblemList map[string]interface{} `firestore:"problem_list"` // list of problem names
}

type ListProblem struct {
	Pid  int    `firestore:"Pid"`  // 1
	Text string `firestore:"Text"` // 느
	Uid  int    `firestore:"Uid"`  // 0
}

func ProblemHandler(ctx *fiber.Ctx) error {
	fmt.Print("ProblemHandler called\t|")
	// cred_file_path := utils.GetCredentialFilePath()
	cred_file_path := "credentials.json"
	pidStr := ctx.Params("pid")
	pid, err := strconv.Atoi(pidStr)
	fmt.Printf("pid: %d\t|", pid)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error converting pid to int")
	}
	uidStr := ctx.Params("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error converting uid to int")
	}
	// uid == 0 이면 기본 문제 리스트를 가져옴
	if uid == 0 {
		if pid == 0 {
			fmt.Print("uid == 0, p_id == 0: get basic problem list\n\n")
			context := ctx.Context()

			problemList, err := GetFireStoreBasicProblemList(context, cred_file_path)
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
			fmt.Printf("\nuid == 0, p_id == %v: get basic problem #%v\n\n", pid, pid)
			context := ctx.Context()

			problem, err := GetFireStoreBasicProblem(context, pid, cred_file_path)
			if err != nil {
				fmt.Println(err)
				return ctx.Status(http.StatusInternalServerError).SendString("Error getting Firestore document")
			}

			jsonData, err := json.Marshal(*problem)
			if err != nil {
				return ctx.Status(http.StatusInternalServerError).SendString("Error converting Firestore data to JSON")
			}

			ctx.Set("Content-Type", "application/json") // set the content type as JSON
			ctx.Set("Access-Control-Allow-Origin", "*") // allows CORS

			// Print the data.
			fmt.Println(*problem)

			return ctx.Status(http.StatusOK).SendString(string(jsonData))
		}
	} else {
		if pid == 0 {

			fmt.Printf("\nuid == %v, p_id == %v: get problem list #%v\n\n", uid, pid, pid)
			context := ctx.Context()

			problem, err := GetFireStoreProblemList(context, uid, cred_file_path)
			if err != nil {
				fmt.Println(err)
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
		} else {
			fmt.Printf("\nuid == %v, p_id == %v: get problem #%v\n\n", uid, pid, pid)
			context := ctx.Context()

			problem, err := GetFireStoreProblem(context, pid, uid, cred_file_path)
			if err != nil {
				fmt.Println(err)
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
}

func GetFireStoreBasicProblemList(context context.Context, cred_file_path string) ([]ListProblem, error) {
	projectID := "seesay"
	// var pList ProblemList
	var prlist []ListProblem
	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("problems/basic_problem/plist").Documents(context).GetAll()
	if err != nil {
		fmt.Printf("Failed to get document: %v", err)
		return nil, err
	}

	for _, d := range doc {
		var pr ListProblem
		d.DataTo(&pr)
		prlist = append(prlist, pr)
	}

	// fmt.Printf("prlist: %v", prlist)

	// pList.ProblemList = docData["problem_list"].(map[string]interface{})

	return prlist, nil
}
func GetFireStoreProblemList(context context.Context, uid int, cred_file_path string) ([]ListProblem, error) {
	projectID := "seesay"
	// var pList ProblemList
	var prlist []ListProblem
	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return nil, err
	}
	defer client.Close()

	query := client.Collection("problems/userProblems/plist").Where("Uid", "==", uid)
	doc, err := query.Documents(context).GetAll()

	if err != nil {
		fmt.Printf("Failed to get document: %v", err)
		return nil, err
	}

	for _, d := range doc {
		var pr ListProblem
		d.DataTo(&pr)
		prlist = append(prlist, pr)
	}

	// fmt.Printf("prlist: %v", prlist)

	// pList.ProblemList = docData["problem_list"].(map[string]interface{})

	return prlist, nil
}

func GetFireStoreBasicProblem(context context.Context, pid int, cred_file_path string) (*BookSave, error) {
	projectID := "seesay"
	var problem BookSave

	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		return &problem, err
	}
	defer client.Close()

	query := client.Collection("problems/basic_problem/list").Where("Pid", "==", pid).Limit(1)

	docs, err := query.Documents(context).GetAll()
	if err != nil {
		return &problem, fmt.Errorf("failed to get documents: %v", err)
	}
	// fmt.Printf("docs: %v", docs)
	err = docs[0].DataTo(&problem)
	if err != nil {
		return &problem, fmt.Errorf("failed to convert document to Problem struct: %v", err)
	}

	return &problem, nil
}

func GetFireStoreProblem(context context.Context, pid int, uid int, cred_file_path string) (*BookSave, error) {
	projectID := "seesay"
	var problem BookSave

	client, err := firestore.NewClient(context, projectID, option.WithCredentialsFile(cred_file_path))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	query := client.Collection("problems/userProblems/list").Where("Pid", "==", pid).Limit(1)

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
