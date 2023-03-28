package handler

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"
)

func TextToPronounce(text string) (string, error) {
	// pwd
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// fmt.Printf("%v", dir)
	// Use python script
	// fmt.Print(fmt.Sprintf("\n%s\\handler\\t2p.py\n", dir))
	cmd, err := exec.Command("python", fmt.Sprintf("%s\\handler\\t2p.py", dir), "--text", text).Output()
	if err != nil {
		return "", err
	}
	s := string(cmd[2 : len(string(cmd))-1])
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "x", "")
	// ㅅ 받침은 생성되지 않으므로 ㅅ 받침 단어들로 특수문자들의 예외 처리를 해준다.
	s = strings.ReplaceAll(s, "*", "ec83b7") //샷
	s = strings.ReplaceAll(s, ".", "eba79b") //맛
	s = strings.ReplaceAll(s, " ", "eb96b3") //떳
	// fmt.Printf("%v", s)
	b, _ := hex.DecodeString(s)
	// fmt.Printf("%v", b)
	ans := ""
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i : i+3])
		if string(r) == "샷" {
			ans += "*"
		} else if string(r) == "맛" {
			ans += ""
		} else if string(r) == "떳" {
			ans += ""
		} else {
			ans += string(r)
		}
		i += size

	}
	fmt.Printf("%v", ans)
	return ans, nil
}
