package handler

var chosung_style = map[rune]int{
	'ㄱ': 1,
	'ㄲ': 1,
	'ㅋ': 1,
	'ㅇ': 1,
	'ㄴ': 2,
	'ㄷ': 2,
	'ㄹ': 2,
	'ㄸ': 2,
	'ㅌ': 2,
	'ㅁ': 3,
	'ㅂ': 3,
	'ㅃ': 3,
	'ㅍ': 3,
	'ㅅ': 4,
	'ㅆ': 4,
	'ㅈ': 5,
	'ㅉ': 5,
	'ㅊ': 5,
	'ㅎ': 6,
}

var chosung_description = map[int]string{
	1: "혀의 뒤쪽을 입천장 안쪽(연구개)에 대고 공기를 터뜨리며 발음합니다.",
	2: "혀 끝을 앞니 뒤의 잇몸(치조)에 대고 공기를 뿜어내며 발음합니다.",
	3: "혀를 가운데에 자연스럽게 두고 입을 닫았다가 두 입술을 떼어 공기를 뿜어내며 발음합니다.",
	4: "잇몸과 혀 사이의 좁은 틈으로 공기를 통과시켜 발음합니다. (혓바닥이 입 천장에 닿을랑 말랑한 공간을 만들어줘야 하는데 혀는 아랫니 뒤쪽에 위치하고 턱을 이용해 공간을 만들어준 후 입을 벌리며 소리를 내야 함)",
	5: "혀를 입 천장(경구개)에 접촉해 공기를 뿜어내며 발음합니다.",
	6: "혀를 가운데에 자연스럽게 두고 폐에서 나오는 공기를 목구멍에 마찰시켜 발음합니다.",
}

var chongsung_style = map[rune]int{
	'ㄱ': 1,
	'ㄴ': 2,
	'ㄷ': 3,
	'ㄹ': 4,
	'ㅁ': 5,
	'ㅂ': 6,
	'ㅇ': 7,
}

var chongsung_description = map[int]string{
	1: "혀의 뒤쪽을 입천장 안쪽(연구개)에 대고 공기를 막으며 발음합니다.",
	2: "혀 끝을 앞니 뒤의 잇몸에 대고 소리를 길게 내며 발음합니다.",
	3: "혀 끝을 앞니 뒤의 잇몸(치조)에 대고 공기를 막으며 발음합니다.",
	4: "혀 끝을 앞니 뒤의 잇몸에 대고 소리를 길게 내며 발음합니다.",
	5: "혀를 자연스럽게 두고 입을 닫으면서 소리를 길게 내며 발음합니다.",
	6: "혀를 자연스럽게 두고 입을 닫으면서 공기를 막아 발음합니다.",
	7: "혀의 뒤쪽을 입천장 안쪽(연구개)에 대고 공기를 터뜨리며 발음합니다.",
}

var chungsung_style = map[rune]int{
	'ㅏ': 1,
	'ㅐ': 2,
	'ㅔ': 2,
	'ㅡ': 3,
	'ㅣ': 3,
	'ㅓ': 4,
	'ㅗ': 5,
	'ㅜ': 6,
	'ㅑ': 130,
	'ㅕ': 430,
	'ㅛ': 530,
	'ㅠ': 630,
	'ㅒ': 230,
	'ㅖ': 230,
	'ㅘ': 15,
	'ㅙ': 25,
	'ㅚ': 35,
	'ㅞ': 26,
	'ㅟ': 36,
	'ㅝ': 46,
	'ㅢ': 3,
}

var chungsung_description = map[int]string{
	1: "검지, 중지, 엄지손가락을 삼각형 모양으로 만든 다음 살짝 입에 넣어보도록 한다. 이 세 개의 손가락이 들어갈 정도로 입술에 힘을 주고 아래턱을 떨어뜨려줘서 발음. 이때 입모양이 동그라미가 되도록.",
	2: "입술을 평평하게 만든다는 생각으로 입꼬리를 옆으로 쭉 늘린 다음 검지와 중지를 겹쳐 무는 느낌으로 입을 좀 더 크게 벌려서 발음.",
	3: "입술을 평평하게 만든 후 입술에 힘을 주고 확실히 입꼬리를 양쪽으로 당겨 미소 짓는듯한 느낌에서 입술만 살짝 떨어지는 느낌의 입모양으로 발음. 이때 양쪽 입꼬리가 내려오지 않아야 하며 입술 전체에 힘이 풀리지 않도록 주의.",
	4: "검지, 중지, 엄지손가락을 삼각형 모양으로 만든 다음 살짝 입에 넣어보도록 한다. 이 세 개의 손가락이 들어갈 정도로 입술에 힘을 주고 아래턱을 떨어뜨려줘서 발음. 이때 입모양이 타원이 되도록.",
	5: "입술에 힘을 주고 오므린 다음에 발음. 입술 주름을 다 접는다는 생각으로 입술에 힘을 주고 모은다. 입모양을 둥글에 만든 다음 소리를 입 밖으로 끌어내는 느낌으로 발음.",
	6: "입술에 힘을 주고 오므린 다음에 발음. 입술 주름을 다 접는다는 생각으로 입술에 힘을 주고 모은다. 입모양을 둥글에 만든 다음 소리를 입 밖으로 끌어내는 느낌으로 발음. 입술을 쭈욱 내밀고 발음.",
}