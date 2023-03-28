# -*- coding: utf-8 -*-
import requests
# encoding
problem_list = ["그"
                , "느","드","르","므","브","스","으","즈","츠","크","트","프","흐","아","야","어","여","오","요","우","유","의","와","외","왜","워","웨","의","애","에","위","예","얘","기린","토끼","염소","여우","팬더","김치","김밥","떡볶이","짜장면","김밥","학교","병원","약국","은행","공항","동쪽","서쪽","남쪽","북쪽","지하철","강남역","승강장","김치","김밥","된장찌개","떡볶이","불고기","노트북","세탁기","청소기","주전자","공기청정기","좌석","가격","팝콘","장르","배우"
                ]

for i, problem in enumerate(problem_list):
    response = requests.request('POST', 'http://127.0.0.1:8080/api/bookshelf', data={'text': problem, 'pid': i+1, 'uid': 0})
    print(response.text)

