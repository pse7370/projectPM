package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	var err error // 에러를 담기 위한 변수

	var getRowCount string = "SELECT COUNT(product_id) * 3 + COUNT(DISTINCT product_type) FROM product"
	var rowCount int = 0
	err = db.QueryRow(getRowCount).Scan(&rowCount)

	fmt.Printf("트리 메뉴를 위한 총 데이터 row 수 : %d\n", rowCount)
	if err != nil {
		log.Println("rowCount 값 가져오기 오류", err)
	}

	//var getSideMenuContentQuery string = "WITH TREE (label, value, parent, product_id) AS (SELECT DISTINCT product_type AS label, product_type AS value, '' AS parent, 0 AS product_id FROM product	UNION ALL SELECT p.product_name AS label, p.product_name AS value, p.product_type AS parent, p.product_id FROM product AS p) SELECT label, value, parent, product_id FROM TREE ORDER BY product_id DESC, parent ASC"
	//var getSideMenuContentQuery string = "WITH TREE (label, value, parent, product_id) AS (SELECT DISTINCT product_type AS label, product_type AS value, '' AS parent, 0 AS product_id FROM product UNION ALL SELECT p.product_name AS label, p.product_name AS value, p.product_type AS parent, p.product_id FROM product AS p UNION ALL SELECT '커스터마이징' AS label, '커스터마이징' + CAST(product_id AS varchar) AS value, product_name AS parent, product_id FROM product UNION ALL SELECT '산출물' AS label, '산출물' + CAST(product_id AS varchar) AS value, product_name AS parent, product_id FROM product) SELECT label, value, parent, product_id FROM TREE ORDER BY product_id DESC, parent ASC"
	var getSideMenuContentQuery string = `WITH TREE (label, value, parent, product_id) AS
	(SELECT DISTINCT product_type AS label, product_type AS value, '' AS parent, 0 AS product_id
		FROM product	
		UNION ALL
		SELECT p.product_name AS label, p.product_name AS value, p.product_type AS parent, p.product_id
		FROM product AS p
		UNION ALL
		SELECT '커스터마이징' AS label, '커스터마이징' + CAST(product_id AS varchar) AS value, product_name AS parent, product_id
		FROM product
		UNION ALL
		SELECT '산출물' AS label, '산출물' + CAST(product_id AS varchar) AS value, product_name AS parent, product_id
		FROM product)
	SELECT label, value, parent, product_id
	FROM TREE ORDER BY product_id DESC, parent ASC`

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
		// DB에서 가져온 데이터 확인용
		// fmt.Printf("\nlable: %s / value: %s / parnet: %s / product_id: %d", label, value, parent, product_id)

		sideMenu.SideMenuList[count] = SideMenuContent{}
		sideMenu.SideMenuList[count].Label = label
		sideMenu.SideMenuList[count].Value = value
		sideMenu.SideMenuList[count].Parent = parent
		sideMenu.SideMenuList[count].Product_id = product_id

		count++

	}

	prettyJsonSideMenu, _ := PrettyJson(sideMenu)

	fmt.Println("\n", prettyJsonSideMenu)

	renderObj.JSON(writer, http.StatusOK, sideMenu)
	// 아래와 같은 과정을 한번에 진행할 수 있는 render 패키지
	// struct를 json으로 변환해 전달해준다.

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

func addDevice(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(100)
	log.Println("addDevice()........")

	// 출입통제기 이미지 저장 위치
	const deviceImageSaveDir string = "C:/deviceImage"

	// 해당 경로에 폴더가 있는지 확인하고 없으면 생성하기
	if _, err := os.Stat(deviceImageSaveDir); os.IsNotExist(err) {
		err := os.Mkdir(deviceImageSaveDir, os.ModeDir)
		if err != nil {
			log.Println("------------폴더 생성 오류-------------")
		}
		fmt.Println("==========해당 경로에 폴더가 없어 새로 생성 : C:/deviceImage")
	}

	multipartForm := request.MultipartForm

	var product Product

	// 파일 저장하기
	for key, _ := range multipartForm.File {
		file, fileHeader, err := request.FormFile(key)
		if err != nil {
			fmt.Println("FormFile ERROR : ", err)
			return
		}

		defer file.Close()
		fmt.Printf("upload file : name [%s], size [%d], header [%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		var imageSavePath string = deviceImageSaveDir + "/" + fileHeader.Filename

		fileUpLoad, err := os.Create(imageSavePath)
		if err != nil {
			fmt.Println("파일 열기 실패 : ", err, "\n", imageSavePath)
			return
		}
		defer fileUpLoad.Close()

		_, err = io.Copy(fileUpLoad, file)
		if err != nil {
			fmt.Println("파일 복사 실패 : ", err)
			return
		}

		fmt.Println("파일 저장 성공!", fileHeader.Filename)

		product.Real_image_name = fileHeader.Filename
		product.Save_image_name = fileHeader.Filename
		product.Save_path = imageSavePath

	} // end for

	// 전달된 데이터 맵과 셋의 데이터들을 변수에 할당
	formData := multipartForm.Value
	// map 형태로 저장되어 있어, map[key]로 접근 가능

	/*
		fmt.Println(formData)
		fmt.Println(formData["@d#"])
		fmt.Println(formData["@d1#"+"product_name"])
	*/

	var product_device ProductDevice
	//var authentication_detailsList Authentication_detailsList
	var product_developerList Product_developerList
	var authentication_detailsList Authentication_detailsList
	//var product_developerList []Product_developer

	//var count int = 0
	for key, value := range formData {
		//fmt.Println(key, "/", value)

		splitRealKey := strings.Split(key, "#")

		if len(splitRealKey) >= 2 {
			//fmt.Println(splitRealKey)

			switch splitRealKey[1] {
			case "product_type":
				product.Product_type = value[0]

			case "product_name":
				product.Product_name = value[0]

			case "product_version":
				product.Product_version = value[0]

			case "explanation":
				product.Explanation = value[0]

			case "width":
				product_device.Width, _ = strconv.ParseFloat(value[0], 64)

			case "height":
				product_device.Height, _ = strconv.ParseFloat(value[0], 64)

			case "depth":
				product_device.Depth, _ = strconv.ParseFloat(value[0], 64)

			case "ip_ratings":
				product_device.Ip_ratings = value[0]

			case "server":
				product_device.Server = value[0]

			case "wi_fi":
				product_device.Wi_fi = value[0]

			case "other":
				product_device.Other = value[0]

			case "auth_type":
				authentication_detailsList.Auth_type = value

			case "one_to_one_max_user":
				authentication_detailsList.One_to_one_max_user = value

			case "one_to_many_max_user":
				authentication_detailsList.One_to_many_max_user = value

			case "one_to_one_max_template":
				authentication_detailsList.One_to_one_max_template = value

			case "one_to_many_max_template":
				authentication_detailsList.One_to_many_max_template = value

			/*
				case "auth_method":
					authentication_detailsList.Auth_method = value

				case "max_users":
					authentication_detailsList.Max_users = value

				case "max_templates":
					authentication_detailsList.Max_templates = value
			*/

			case "department":
				product_developerList.DepartmentList = value

			case "employees_number":
				product_developerList.Employees_numberList = value

			case "employees_name":
				product_developerList.Employees_nameList = value

			case "start_date":
				product_developerList.Start_dateList = value

			case "end_date":
				product_developerList.End_dateList = value

			}

		} // end if

	} // end for

	fmt.Println("product : ", product)
	fmt.Println("product_device : ", product_device)
	fmt.Println("authentication_detailsList : ", authentication_detailsList)
	fmt.Println("product_developerList : ", product_developerList)

	sliceLength_auth := len(authentication_detailsList.Auth_type)
	var authentication_details = make([]Authentication_details, sliceLength_auth)
	for i := 0; i < sliceLength_auth; i++ {
		authentication_details[i].Auth_type = authentication_detailsList.Auth_type[i]

		temp, _ := strconv.ParseInt(authentication_detailsList.One_to_one_max_user[i], 10, 32)
		authentication_details[i].One_to_one_max_user = temp

		temp2, _ := strconv.ParseInt(authentication_detailsList.One_to_many_max_user[i], 10, 32)
		authentication_details[1].One_to_many_max_user = temp2

		temp3, _ := strconv.ParseInt(authentication_detailsList.One_to_one_max_template[i], 10, 32)
		authentication_details[i].One_to_one_max_template = temp3

		temp4, _ := strconv.ParseInt(authentication_detailsList.One_to_many_max_template[i], 10, 32)
		authentication_details[1].One_to_many_max_template = temp4

	}

	fmt.Println("authentication_details : ", authentication_details)

	/*
		sliceLength_auth := len(authentication_detailsList.Auth_type)
		var authentication_details = make([]Authentication_details, sliceLength_auth)
		for i := 0; i < sliceLength_auth; i++ {
			authentication_details[i].Auth_type = authentication_detailsList.Auth_type[i]
			authentication_details[i].Auth_method = authentication_detailsList.Auth_method[i]

			temp, _ := strconv.ParseInt(authentication_detailsList.Max_users[i], 10, 32)
			authentication_details[i].Max_users = temp

			temp2, _ := strconv.ParseInt(authentication_detailsList.Max_templates[i], 10, 32)
			authentication_details[1].Max_templates = temp2

		}


		fmt.Println("authentication_details : ", authentication_details)
	*/

	sliceLength_developer := len(product_developerList.Employees_numberList)
	var product_developer = make([]Product_developer, sliceLength_developer)
	for i := 0; i < sliceLength_developer; i++ {
		product_developer[i].Department = product_developerList.DepartmentList[i]

		temp, _ := strconv.ParseInt(product_developerList.Employees_numberList[i], 10, 32)
		product_developer[i].Employees_number = temp

		product_developer[i].Employees_name = product_developerList.Employees_nameList[i]
		product_developer[i].Start_date = product_developerList.Start_dateList[i]
		product_developer[i].End_date = product_developerList.End_dateList[i]
	}

	fmt.Println("product_developer : ", product_developer)

	// 전달된 출입통제기 스펙을 DB에 insert하기

	// 여러 테이블에 insert하는 과정을 하나의 트랜잭션으로 묶기
	transaction, err := db.Begin()
	if err != nil {
		fmt.Println("--------트랜잭션 생성 오류---------")
		log.Fatal(err)
	}

	// 에러 발생시 rollback 처리
	defer transaction.Rollback()

	_, err = db.Exec(`INSERT INTO product(product_type, product_name, product_version, real_image_name, save_image_name, save_path, explanation)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
		product.Product_type, product.Product_name, product.Product_version, product.Real_image_name, product.Save_image_name, product.Save_path, product.Explanation)
	if err != nil {
		fmt.Println("===========product 테이블 insert 실패===========")
		log.Fatal(err)
	}

	var product_id int
	err = db.QueryRow("SELECT product_id FROM product WHERE product_name = ?", product.Product_name).Scan(&product_id)
	if err != nil {
		fmt.Println("===========product 테이블 porduct_id 가져오기 실패===========")
		log.Fatal(err)
	}
	fmt.Println("product_id : ", product_id)

	/*
		for i := 0; i < sliceLength_auth; i++ {
			_, err = db.Exec(`INSERT INTO authentication_details(product_id, auth_type, auth_method, max_users, max_templates)
							VALUES (?, ?, ?, ?, ?)`,
				product_id, authentication_details[i].Auth_type, authentication_details[i].Auth_method, authentication_details[i].Max_users, authentication_details[i].Max_templates)
			if err != nil {
				fmt.Printf("===========authentication_details 테이블 insert 실패 '%d'================\n", i)
				log.Fatal(err)
			}
		}
	*/

	for i := 0; i < sliceLength_auth; i++ {
		_, err = db.Exec(`INSERT INTO authentication_details(product_id, auth_type, one_to_one_max_user, one_to_many_max_user, one_to_one_max_template, one_to_many_max_template) 
						VALUES (?, ?, ?, ?, ?)`,
			product_id, authentication_details[i].Auth_type, authentication_details[i].One_to_one_max_user, authentication_details[i].One_to_many_max_user, authentication_details[i].One_to_one_max_template, authentication_details[i].One_to_one_max_template)
		if err != nil {
			fmt.Printf("===========authentication_details 테이블 insert 실패 '%d'================\n", i)
			log.Fatal(err)
		}
	}

	_, err = db.Exec(`INSERT INTO product_device(product_id, width, height, depth, ip_ratings, server, wi_fi, other) 
					VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		product_id, product_device.Width, product_device.Height, product_device.Depth, product_device.Ip_ratings, product_device.Server, product_device.Wi_fi, product_device.Other)
	if err != nil {
		fmt.Println("===========product_device 테이블 insert 실패===========")
		log.Fatal(err)
	}

	for i := 0; i < sliceLength_developer; i++ {
		_, err = db.Exec(`INSERT INTO product_developer(product_id, department, employees_number, employees_name, start_date, end_date) 
						VALUES (?, ?, ?, ?, ?, ?)`,
			product_id, product_developer[i].Department, product_developer[i].Employees_number, product_developer[i].Employees_name, product_developer[i].Start_date, product_developer[i].End_date)
		if err != nil {
			fmt.Printf("===========product_developer 테이블 insert 실패 '%d'================\n", i)
			log.Fatal(err)
		}
	}

	// 트랜잭션 종료
	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	result.ResultCode = 1

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, result)

}
