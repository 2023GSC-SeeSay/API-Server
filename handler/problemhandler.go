package handler

import (
	"api-server/utils"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func ProblemHandler(c *fiber.Ctx) error {
	cred_file_path := utils.GetCredentialFilePath()
	fmt.Printf(cred_file_path)
	pid := c.Params("pid")
	fmt.Printf(pid)

	// firebase firestore query로 pid의 정보를 가져옴
	// 가져온 정보를 json으로 변환하여 response로 보냄
	return nil
}

func GetFireStoreProblem(pid string, cred_file_path string)  {
	// firebase firestore query로 pid의 정보를 가져옴
	// 가져온 정보를 json으로 변환하여 response로 보냄

	// ctx := context.Background()

}