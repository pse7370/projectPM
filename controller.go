package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/urfave/negroni" // http 미들웨어 negroni
)

// DB에 접속하기 위한 정보를 담을 구조체
type Ms_sqlConfig struct {
	DatabaseType string `json:"databaseType"`
	DatabaseIP   string `json:"databaseIP"`
	DatabasePort int    `json:"databasePort"`
	UserID       string `json:"userID"`
	UserPW       string `json:"userPW"`
	DataBaseName string `json:"dataBaseName"`
}

var dbConfig *Ms_sqlConfig = nil
var mux *http.ServeMux = nil
var db *sql.DB = nil

// DB설정 json 파일(dbSetting.json) 읽어 Ms_sqlConfig 구조체의 포인터를 반환하는 함수
func ReadDBSetting() (*Ms_sqlConfig, error) {

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// 실행 중인 파일을 현재 디렉토리 위치
	// "c:\GoWorkspace\src\projectPM"
	// os.Getwd()를 사용해도 같은 결과값 반환

	if err != nil {
		log.Println("-----can not find currentDir------")
		panic(err)
	}

	dbSettingFile := currentDir + "/dbSetting.json"
	// json 파일 경로 변수 설정

	file, err := os.Open(dbSettingFile)

	if err != nil {
		log.Println("------can not open DB setting json file------")
		panic(err)
	}

	jsonDecoder := json.NewDecoder(file)
	// json문자열을 struct로 변경하기 위한 디코더 생성

	err = jsonDecoder.Decode(&dbConfig) // 디코딩
	if err != nil {
		log.Println("=======Decoding Fail=======")
		panic(err)
	}

	log.Println("Success Read json file and decoding struct")

	return dbConfig, err

}

// DB 연결 함수
func DBConnect() {
	//var dbConfig *Ms_sqlConfig = new(Ms_sqlConfig)
	// new()를 사용해 Ms_sqlConfigfmf 제로 값으로 초기화하고, 포인터를 dbConfig에 할당
	// var dbConfig *Ms_sqlConfig = nil

	dbConfig, _ = ReadDBSetting()

	//fmt.Println(dbConfig.DataBaseName)
	//fmt.Println(dbConfig.DatabaseIP)

	dbConnectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		dbConfig.DatabaseIP,
		dbConfig.UserID,
		dbConfig.UserPW,
		dbConfig.DataBaseName)
	// %s는 문자열 출력표현
	// fmt.Sprintf는 문자열 값 반환, 원하는 문자열 형식으로 만들 때 사용. 화면에 출력되지 않음.

	var err error
	db, err = sql.Open(dbConfig.DatabaseType, dbConnectionString)
	// := 는 생성자를 새로 만들어 주는 것이기 때문에 전역 변수로 생성한 db와는 다른 객체 생성

	if err != nil {
		log.Println("**********Fail to open DB***********")
		panic(err)
	}

	defer db.Close() // DB 지연 종료

}

// 서버 연결 함수
func StartServer() {
	// DB 연결
	//var dbConfig *Ms_sqlConfig = new(Ms_sqlConfig)
	// new()를 사용해 Ms_sqlConfigfmf 제로 값으로 초기화하고, 포인터를 dbConfig에 할당
	// var dbConfig *Ms_sqlConfig = nil

	dbConfig, _ = ReadDBSetting()

	//fmt.Println(dbConfig.DataBaseName)
	//fmt.Println(dbConfig.DatabaseIP)

	dbConnectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		dbConfig.DatabaseIP,
		dbConfig.UserID,
		dbConfig.UserPW,
		dbConfig.DataBaseName)
	// %s는 문자열 출력표현
	// fmt.Sprintf는 문자열 값 반환, 원하는 문자열 형식으로 만들 때 사용. 화면에 출력되지 않음.

	var err error
	db, err = sql.Open(dbConfig.DatabaseType, dbConnectionString)
	// := 는 생성자를 새로 만들어 주는 것이기 때문에 전역 변수로 생성한 db와는 다른 객체 생성

	if err != nil {
		log.Println("**********Fail to open DB***********")
		panic(err)
	}

	defer db.Close() // DB 지연 종료

	// port 번호 지정
	const portNumber int = 5000
	mainServerAddress := fmt.Sprintf(":%d", portNumber)
	// :5000 형태의 문자열로 만들기

	mux = http.NewServeMux()
	// ServeMux 객체 생성
	negroniObject := negroni.Classic()
	// negroni 사용을 위한 객체 생성
	// 복구, 로그 기능을 사용자가 만든 서버 객체와 연동해 쉽게 사용 가능

	negroniObject.UseHandler(mux)

	basePath, _ := os.Getwd()
	fmt.Println(basePath)

	// 저장한 제품 이미지 경로
	imageFileServer := http.FileServer(http.Dir(basePath))
	mux.Handle("/static/", http.StripPrefix("/static/", imageFileServer))

	// 페이지 경로
	fileServer := http.FileServer(http.Dir(basePath + "/webRoot"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	fmt.Println("====================Start Server======================")

	//CustomHandleFunc("/productMangement/sideMenu", getSideMenuContent)

	// 사이드 트리 메뉴 구성 값 가져오기
	mux.HandleFunc("/productMangement/sideMenu", productMangementHandler)

	// 제품명 중복 체크
	mux.HandleFunc("/productMangement/checkProductName", checkProductName)

	// 출입통제기 등록하기
	mux.HandleFunc("/productMangement/addDevice", addDevice)
	// SW 등록하기
	mux.HandleFunc("/productMangement/addSW", addSW)
	// 출입통제기 상세 내역 불러오기
	mux.HandleFunc("/productMangement/deviceContent", getDeviceContent)
	// SW 상세 내역 불러오기
	mux.HandleFunc("/productMangement/SWcontent", getSWcontent)
	// 출입통제기 삭제하기
	mux.HandleFunc("/productMangement/deleteDevice", deleteDevice)
	// SW 삭제하기
	mux.HandleFunc("/productMangement/deleteSW", deleteSW)
	// 출입통제기 수정
	mux.HandleFunc("/productMangement/modifyDevice", modifyDevice)
	// SW 수정
	mux.HandleFunc("/productMangement/modifySW", modifySW)

	// 커스터마이징 이력 목록 불러오기
	mux.HandleFunc("/productMangement/getCustomizingList", getCustomizingList)
	// 커스터마이징 이력 삭제
	mux.HandleFunc("/productMangement/deleteCustomizing", deleteCustomizing)
	// 커스터마이징 이력 추가/수정하기
	mux.HandleFunc("/productMangement/modifyCustomizing", modifyCustomizing)

	// 산출물 목록 불러오기
	mux.HandleFunc("/productMangement/getOutputList", getOutputList)
	// 산출물 등록하기
	mux.HandleFunc("/productMangement/addOutput", addOutput)
	// 산출물 상세 조회
	mux.HandleFunc("/productMangement/getOutputContent", getOutputContent)
	// 산출물 삭제
	mux.HandleFunc("/productMangement/deleteOutput", deleteOutput)
	// 산출물 검색
	mux.HandleFunc("/productMangement/getSearchOutputList", getSearchOutputList)
	// 산출물 하나의 첨부 파일 다운로드
	mux.HandleFunc("/productMangement/downloadAttachment", downloadAttachment)
	// 산출물 복수의 첨부 파일 다운로드
	mux.HandleFunc("/productMangement/downloadAttachmentList", downloadAttachmentList)
	// 산출물 수정
	mux.HandleFunc("/productMangement/modifyOutput", modifyOutput)
	// 압축 파일 삭제
	mux.HandleFunc("/productMangement/deleteZipFile", deleteZipFile)

	// http.ListenAndServe(mainServerAddress, mux)
	// 웹서버를 실제로 동작시키기 위한 함수, 서버가 동작할 포트 번호 지정
	http.ListenAndServe(mainServerAddress, negroniObject)

}

func main() {
	StartServer()

}
