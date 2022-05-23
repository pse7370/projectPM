package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

func productMangementHandler(writer http.ResponseWriter, request *http.Request) {

	log.Println("productMangementHandler()-----상품관리 메뉴 접속")
	switch request.Method {
	case "GET":
		getSideMenuContent(writer, request)
	}

}

func getSideMenuContent(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("getSideMenuContent()........")
	renderObj := render.New()
	var err error

	//DB연결
	dbConfig, _ = ReadDBSetting()

	dbConnectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		dbConfig.DatabaseIP,
		dbConfig.UserID,
		dbConfig.UserPW,
		dbConfig.DataBaseName)
	// %s는 문자열 출력표현
	// fmt.Sprintf는 문자열 값 반환, 원하는 문자열 형식으로 만들 때 사용. 화면에 출력되지 않음.

	db, err := sql.Open(dbConfig.DatabaseType, dbConnectionString)
	if err != nil {
		log.Println("**********Fail to open DB***********")
		panic(err)
	}

	defer db.Close()

	var getRowCount string = "SELECT COUNT(product_id) * 3 + COUNT(DISTINCT product_type) FROM product"
	var rowCount int = 0
	err = db.QueryRow(getRowCount).Scan(&rowCount)

	fmt.Printf("트리 메뉴를 위한 총 데이터 row 수 : %d\n", rowCount)
	if err != nil {
		log.Println("rowCount 값 가져오기 오류", err)
	}

	//var getSideMenuContentQuery string = "WITH TREE (label, value, parent, product_id) AS (SELECT DISTINCT product_type AS label, product_type AS value, '' AS parent, 0 AS product_id FROM product	UNION ALL SELECT p.product_name AS label, p.product_name AS value, p.product_type AS parent, p.product_id FROM product AS p) SELECT label, value, parent, product_id FROM TREE ORDER BY product_id DESC, parent ASC"
	var getSideMenuContentQuery string = "WITH TREE (label, value, parent, product_id) AS (SELECT DISTINCT product_type AS label, product_type AS value, '' AS parent, 0 AS product_id FROM product UNION ALL SELECT p.product_name AS label, p.product_name AS value, p.product_type AS parent, p.product_id FROM product AS p UNION ALL SELECT '커스터마이징' AS label, '커스터마이징' + CAST(product_id AS varchar) AS value, product_name AS parent, product_id FROM product UNION ALL SELECT '산출물' AS label, '산출물' + CAST(product_id AS varchar) AS value, product_name AS parent, product_id FROM product) SELECT label, value, parent, product_id FROM TREE ORDER BY product_id DESC, parent ASC"
	rows, err := db.Query(getSideMenuContentQuery)
	if err != nil {
		log.Println("########쿼리문 실행 오류########")
		log.Fatal(err)
	}

	fmt.Println(getSideMenuContentQuery)
	defer rows.Close()

	var count int = 0
	var sideMenu SideMenu
	// sideMenuList := make([]SideMenuContent, rowCount)
	sideMenu.SideMenuList = make([]SideMenuContent, rowCount)

	var label string
	var value string
	var parent string
	var product_id int32
	for rows.Next() {

		err := rows.Scan(&label, &value, &parent, &product_id)
		if err != nil {
			log.Println("사이드 트리메뉴 데이터 가져오기 실패 :", err)
		}
		fmt.Printf("\nlable: %s / value: %s / parnet: %s / product_id: %d", label, value, parent, product_id)

		sideMenu.SideMenuList[count] = SideMenuContent{}
		sideMenu.SideMenuList[count].Label = label
		sideMenu.SideMenuList[count].Value = value
		sideMenu.SideMenuList[count].Parent = parent
		sideMenu.SideMenuList[count].Product_id = product_id

		count++

	}

	fmt.Println("\n", sideMenu)

	renderObj.JSON(writer, http.StatusOK, sideMenu)

	/*
		responseSideMenuData, _ := json.Marshal(sideMenu)
		// 전달 데이터 타임 설정
		writer.Header().Set("contetnt-type", "application/json")
		// 응답코드 작성
		writer.WriteHeader(http.StatusOK)
		// request body
		writer.Write(responseSideMenuData)
	*/

}
