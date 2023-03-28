package handler

import (
	"fmt"
)

func GenerateDescMouth(sentence string) string {
	desc := ""
	for i, v := range sentence {
		// fmt.Printf("%d : %c", i, v)
		if (i/3)%3 == 1 {
			if chungsung_style[v] < 10 {

				desc += chungsung_description[chungsung_style[v]]
			} else {
				if chungsung_style[v] < 100 {
					desc += chungsung_description[chungsung_style[v]%10]
					desc += "이어서 다음과 같이 발음을 합니다."
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
	}
	fmt.Printf("%s", desc)
	return desc
}

func GenerateDescTongue(sentence string) string {
	desc := ""
	for i, v := range sentence {
		if (i/3)%3 == 0 {
			desc += chosung_description[chosung_style[v]]
			desc += "\n"
		}
		if (i/3)%3 == 2 {
			desc += chongsung_description[chongsung_style[v]]
			desc += "\n"
		}
	}
	fmt.Printf("%s", desc)
	return desc
}