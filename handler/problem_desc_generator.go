package handler

import (
	"fmt"
)

func GenerateDescMouth(sentence string) string {
	/* This functions are not allowed for exclamation mark except "*" */
	desc := ""
	count := 0
	for _, v := range sentence {
		// fmt.Printf("%d : %c", i, v)
		if count%3 == 1 {
			desc += string(v)
			desc += " 발음 :"
			if chungsung_style[v] < 10 {
				desc += chungsung_description[chungsung_style[v]]
			} else {
				if chungsung_style[v] < 100 {
					desc += chungsung_description[chungsung_style[v]%10]
					desc += " 이어서 다음과 같이 발음을 합니다."
					desc += chungsung_description[chungsung_style[v]/10]
				} else {
					c := chungsung_style[v] / 10
					desc += chungsung_description[c%10]
					desc += chungsung_description[c/10]
					desc += "위의 두 발음을 빠르게 이어합니다."
				}
			}
			desc += "\n"
		}
		count++
	}
	fmt.Printf("%s", desc)
	return desc
}

func GenerateDescTongue(sentence string) string {
	fmt.Print(sentence)
	/* This functions are not allowed for exclamation mark except "*" */
	desc := ""
	count := 0
	for _, v := range sentence {
		if count%3 == 0 {
			desc += string(v)
			desc += " 발음 :"
			desc += chosung_description[chosung_style[v]]
			desc += "\n"
		}
		if count%3 == 2 {
			if v == '*' {
				count++
				continue
			}
			desc += string(v)
			desc += "발음 :"
			desc += chongsung_description[chongsung_style[v]]
			desc += "\n"
		}
		count++
	}
	fmt.Printf("%s", desc)
	return desc
}
