package handler

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"
	"unicode/utf8"
)


func TextToPronounce() {
	// Use python script
	cmd, err := exec.Command("python", "C:\\workspace\\API-Server\\API-Server\\handler\\1.py", "--text", "안녕하세요").Output()
	if err != nil {
		panic(err)
	}
	s := string(cmd[2:len(string(cmd))-1])
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "x", "")

	b, _ := hex.DecodeString(s)
	// fmt.Printf("%v", b)
	ans := ""
	for i := 0; i < len(b);  {
		r, size := utf8.DecodeRune(b[i:i+3])
		ans += string(r)
		i += size
	}
	fmt.Printf("%v", ans)
}
