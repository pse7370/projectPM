package main

import (
	"encoding/json"
	"net/http"
)

// struct를 json으로 변환시 정렬해서 변환해주는 함수
func PrettyJson(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// url 경로마다 실행할 함수 매칭을 위한 커스텀 HandleFunc
func CustomHandleFunc(url string, handlefunc func(writer http.ResponseWriter, request *http.Request)) {
	if mux == nil {
		mux = http.NewServeMux()
	}

	// url 경로별 handlefunc 매칭
	mux.HandleFunc(url, handlefunc)
}
