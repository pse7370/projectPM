package main

import "net/http"

// url 경로마다 실행할 함수 매칭을 위한 커스텀 HandleFunc
func CustomHandleFunc(url string, handlefunc func(writer http.ResponseWriter, request *http.Request)) {
	if mux == nil {
		mux = http.NewServeMux()
	}

	// url 경로별 handlefunc 매칭
	mux.HandleFunc(url, handlefunc)
}
