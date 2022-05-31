package main

import (
	"database/sql"
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

// 사이드 트리메뉴 구성하기
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

// 출입통제기 등록
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
			log.Fatalln(err)
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
	fmt.Println("sliceLength_auth : ", sliceLength_auth)
	var authentication_details = make([]Authentication_details, sliceLength_auth)
	for i := 0; i < sliceLength_auth; i++ {
		authentication_details[i].Auth_type = authentication_detailsList.Auth_type[i]

		temp, _ := strconv.ParseInt(authentication_detailsList.One_to_one_max_user[i], 10, 32)
		authentication_details[i].One_to_one_max_user = int32(temp)

		temp2, _ := strconv.ParseInt(authentication_detailsList.One_to_many_max_user[i], 10, 32)
		authentication_details[i].One_to_many_max_user = int32(temp2)

		temp3, _ := strconv.ParseInt(authentication_detailsList.One_to_one_max_template[i], 10, 32)
		authentication_details[i].One_to_one_max_template = int32(temp3)

		temp4, _ := strconv.ParseInt(authentication_detailsList.One_to_many_max_template[i], 10, 32)
		authentication_details[i].One_to_many_max_template = int32(temp4)

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
	fmt.Println("sliceLength_developer : ", sliceLength_developer)
	var product_developer = make([]Product_developer, sliceLength_developer)
	for i := 0; i < sliceLength_developer; i++ {
		product_developer[i].Department = product_developerList.DepartmentList[i]

		temp, _ := strconv.ParseInt(product_developerList.Employees_numberList[i], 10, 32)
		product_developer[i].Employees_number = int32(temp)

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
						VALUES (?, ?, ?, ?, ?, ?)`,
			product_id, authentication_details[i].Auth_type, authentication_details[i].One_to_one_max_user, authentication_details[i].One_to_many_max_user, authentication_details[i].One_to_one_max_template, authentication_details[i].One_to_one_max_template)
		if err != nil {
			fmt.Printf("===========authentication_details 테이블 insert 실패 '%d'================\n", i)
			log.Fatal(err)
		}
	} // end for

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

// SW 등록하기
func addSW(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(100)
	log.Println("addSW()........")

	// 출입통제기 이미지 저장 위치
	const swImageSaveDir string = "C:/SWimage"

	// 해당 경로에 폴더가 있는지 확인하고 없으면 생성하기
	if _, err := os.Stat(swImageSaveDir); os.IsNotExist(err) {
		err := os.Mkdir(swImageSaveDir, os.ModeDir)
		if err != nil {
			log.Println("------------폴더 생성 오류-------------")
			log.Fatalln(err)
		}
		fmt.Println("==========해당 경로에 폴더가 없어 새로 생성 : C:/SWimage")
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

		var imageSavePath string = swImageSaveDir + "/" + fileHeader.Filename

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

	//fmt.Println(formData)

	var product_sw ProductSW
	var product_developerList Product_developerList

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

			case "simultaneous_connection":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_sw.Simultaneous = int32(temp)

			case "available_db":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_sw.Available_db = int32(temp)

			case "available_os":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_sw.Available_os = int32(temp)

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
	fmt.Println("product_sw : ", product_sw)
	fmt.Println("product_developerList : ", product_developerList)

	sliceLength_developer := len(product_developerList.Employees_numberList)
	fmt.Println("sliceLength_developer : ", sliceLength_developer)
	var product_developer = make([]Product_developer, sliceLength_developer)
	for i := 0; i < sliceLength_developer; i++ {
		product_developer[i].Department = product_developerList.DepartmentList[i]

		temp, _ := strconv.ParseInt(product_developerList.Employees_numberList[i], 10, 32)
		product_developer[i].Employees_number = int32(temp)

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

	_, err = db.Exec(`INSERT INTO product_sw(product_id, simultaneous_connection, available_db, available_os) VALUES(?, ?, ?, ?)`,
		product_id, product_sw.Simultaneous, product_sw.Available_db, product_sw.Available_os)
	if err != nil {
		fmt.Println("===========product_sw 테이블 insert 실패===========")
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
	} // end for

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

// 출입통제기 상세 내역 가져오기
func getDeviceContent(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("getDeviceContent()........")
	request.ParseForm()

	formData := request.Form
	//fmt.Println(formData)

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	var product Product
	var product_device ProductDevice
	var authenticationList []Authentication_details = []Authentication_details{}
	var developerList []Product_developer = []Product_developer{}

	/*
		// product, authentication_details, product_device, product_developer 4개 테이블 조인
		// 쿼리문 출력시 product_developer 중복값 제거가 잘 되지 않아 방식 변경
		var getDeviceDetailsQuery string = `SELECT DISTINCT
												p.product_id, p.product_type, p.product_name,
												p.product_version, p.save_path, p.explanation,
												ad.auth_type, ad.one_to_one_max_user, ad.one_to_many_max_user,
												ad.one_to_one_max_template, ad.one_to_many_max_template,
												pd.width, pd.height, pd.depth, pd.ip_ratings,
												pd.server, pd.wi_fi, pd.other,
												dev.developer_id, dev.department, dev.employees_number,
												dev.employees_name, convert(varchar(10), dev.start_date, 20) AS start_date,
												convert(varchar(10), dev.end_date, 20) AS end_date
											FROM product p
											LEFT JOIN authentication_details ad
											ON p.product_id = ad.product_id
											LEFT JOIN product_device pd
											ON p.product_id = pd.product_id
											LEFT JOIN product_developer dev
											ON p.product_id = dev.product_id
											WHERE p.product_id = ?`
		rows, err := db.Query(getDeviceDetailsQuery, product_id)
		if err != nil {
			log.Fatalf("==========product_id = %d인 제품의 상세 내역 DB에서 가져오기 실패===========\n", product_id)
			log.Println(err)
		}
	*/

	var getDeviceDetailsQuery string = `SELECT
											p.product_id, 
											p.product_type, 
											p.product_name, 
											p.product_version, 
											p.save_path, 
											p.explanation,
											ad.auth_type, 
											ad.one_to_one_max_user, 
											ad.one_to_many_max_user,
											ad.one_to_one_max_template, 
											ad.one_to_many_max_template,
											pd.width, 
											pd.height, 
											pd.depth, 
											pd.ip_ratings, 
											pd.server, 
											pd.wi_fi, 
											pd.other
										FROM product p
										LEFT JOIN authentication_details ad
										ON p.product_id = ad.product_id
										LEFT JOIN product_device pd
										ON p.product_id = pd.product_id
										WHERE p.product_id = ?`

	rows, err := db.Query(getDeviceDetailsQuery, product_id)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 상세 내역 DB에서 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	fmt.Println(getDeviceDetailsQuery)
	defer rows.Close()

	var auth_type string
	var one_to_one_max_user sql.NullInt32
	var one_to_many_max_user sql.NullInt32
	var one_to_one_max_template sql.NullInt32
	var one_to_many_max_template sql.NullInt32

	var developer_id int32
	var department string
	var employees_number int32
	var employees_name string
	var start_date sql.NullString
	var end_date sql.NullString

	/*
		for rows.Next() {
			err := rows.Scan(&product.Product_id, &product.Product_type,
				&product.Product_name, &product.Product_version,
				&product.Save_path, &product.Explanation,
				&auth_type, &one_to_one_max_user,
				&one_to_many_max_user, &one_to_one_max_template,
				&one_to_many_max_template, &product_device.Width,
				&product_device.Height, &product_device.Depth,
				&product_device.Ip_ratings, &product_device.Server,
				&product_device.Wi_fi, &product_device.Other,
				&developer_id, &department,
				&employees_number, &employees_name,
				&start_date, &end_date)

			// 여러 행 나올 수 있는 authentication_details, product_developer는 모든 값을
			// slice에 순차적으로 저장
			authentication_details = append(authentication_details,
				Authentication_details{
					Auth_type:                auth_type,
					One_to_one_max_user:      one_to_one_max_user.Int32,
					One_to_many_max_user:     one_to_many_max_user.Int32,
					One_to_one_max_template:  one_to_one_max_template.Int32,
					One_to_many_max_template: one_to_many_max_template.Int32,
				},
			)

			var count int = 0
			if count == 0 {
				product_developer = append(product_developer,
					Product_developer{
						Developer_id:     developer_id,
						Department:       department,
						Employees_number: employees_number,
						Employees_name:   employees_name,
						Start_date:       start_date.String,
						End_date:         end_date.String,
					},
				)
			}

			for count, value := range product_developer {
				fmt.Printf("product_developer[%d].Developer_id : %d / developer_id : %d\n", count, value.Developer_id, developer_id)
				if value.Developer_id != developer_id {
					product_developer = append(product_developer,
						Product_developer{
							Developer_id:     developer_id,
							Department:       department,
							Employees_number: employees_number,
							Employees_name:   employees_name,
							Start_date:       start_date.String,
							End_date:         end_date.String,
						},
					)
				}
			}

			for i := 0; i < len(product_developer); i++ {
				fmt.Printf("product_developer[%d].Developer_id : %d / developer_id : %d\n", i, product_developer[i].Developer_id, developer_id)
				if product_developer[i].Developer_id != developer_id {
					product_developer = append(product_developer,
						Product_developer{
							Developer_id:     developer_id,
							Department:       department,
							Employees_number: employees_number,
							Employees_name:   employees_name,
							Start_date:       start_date.String,
							End_date:         end_date.String,
						},
					)
				}
				continue
			}
			count++

			if err != nil {
				log.Fatal(err)
			}
		} // end for rows.Next
	*/

	for rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Product_type,
			&product.Product_name, &product.Product_version,
			&product.Save_path, &product.Explanation,
			&auth_type, &one_to_one_max_user,
			&one_to_many_max_user, &one_to_one_max_template,
			&one_to_many_max_template, &product_device.Width,
			&product_device.Height, &product_device.Depth,
			&product_device.Ip_ratings, &product_device.Server,
			&product_device.Wi_fi, &product_device.Other)

		// 여러 행 나올 수 있는 authentication_details, product_developer는 모든 값을
		// slice에 순차적으로 저장
		authenticationList = append(authenticationList,
			Authentication_details{
				Auth_type:                auth_type,
				One_to_one_max_user:      one_to_one_max_user.Int32,
				One_to_many_max_user:     one_to_many_max_user.Int32,
				One_to_one_max_template:  one_to_one_max_template.Int32,
				One_to_many_max_template: one_to_many_max_template.Int32,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	var getDeveloperInfoQuery string = `SELECT
											dev.developer_id, 
											dev.department, 
											dev.employees_number, 
											dev.employees_name, 
											convert(varchar(10), dev.start_date, 20) AS start_date, 
											convert(varchar(10), dev.end_date, 20) AS end_date
										FROM product_developer dev
										WHERE dev.product_id = ?`

	rows, err = db.Query(getDeveloperInfoQuery, product_id)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 담당 개발자 DB에서 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&developer_id, &department,
			&employees_number, &employees_name,
			&start_date, &end_date)

		developerList = append(developerList,
			Product_developer{
				Developer_id:     developer_id,
				Department:       department,
				Employees_number: employees_number,
				Employees_name:   employees_name,
				Start_date:       start_date.String,
				End_date:         end_date.String,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	/*
		fmt.Println(product)
		fmt.Println(product_device)
		fmt.Println(authenticationList)
		fmt.Println(developerList)
	*/

	var deviceContent DeviceContent = DeviceContent{
		Product:            product,
		AuthenticationList: authenticationList,
		Product_device:     product_device,
		DeveloperList:      developerList,
	}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, deviceContent)

}

func getSWcontent(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("getSWcontent()........")
	request.ParseForm()

	formData := request.Form

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	var product Product
	var product_sw ProductSW
	var developerList []Product_developer = []Product_developer{}

	var getSwDetailsQuery string = `SELECT
										p.product_id, 
										p.product_type, 
										p.product_name, 
										p.product_version, 
										p.save_path, 
										p.explanation,
										ps.simultaneous_connection, 
										ps.available_db,
										ps.available_os
									FROM product p
									LEFT JOIN product_sw ps
									ON p.product_id = ps.product_id
									WHERE p.product_id = ?`

	rows, err := db.Query(getSwDetailsQuery, product_id)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 상세 내역 DB에서 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	fmt.Println(getSwDetailsQuery)
	defer rows.Close()

	var developer_id int32
	var department string
	var employees_number int32
	var employees_name string
	var start_date sql.NullString
	var end_date sql.NullString

	for rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Product_type,
			&product.Product_name, &product.Product_version,
			&product.Save_path, &product.Explanation,
			&product_sw.Simultaneous, &product_sw.Available_db,
			&product_sw.Available_os)

		if err != nil {
			log.Fatal(err)
		}
	}

	var getDeveloperInfoQuery string = `SELECT
											dev.developer_id, 
											dev.department, 
											dev.employees_number, 
											dev.employees_name, 
											convert(varchar(10), dev.start_date, 20) AS start_date, 
											convert(varchar(10), dev.end_date, 20) AS end_date
										FROM product_developer dev
										WHERE dev.product_id = ?`

	rows, err = db.Query(getDeveloperInfoQuery, product_id)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 담당 개발자 DB에서 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&developer_id, &department,
			&employees_number, &employees_name,
			&start_date, &end_date)

		developerList = append(developerList,
			Product_developer{
				Developer_id:     developer_id,
				Department:       department,
				Employees_number: employees_number,
				Employees_name:   employees_name,
				Start_date:       start_date.String,
				End_date:         end_date.String,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(product)
	fmt.Println(product_sw)
	fmt.Println(developerList)

	var swContent = SwContent{
		Product:       product,
		Product_sw:    product_sw,
		DeveloperList: developerList,
	}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, swContent)

}

func deleteDevice(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("deleteDevice()........")
	request.ParseForm()

	formData := request.Form

	//fmt.Println(formData)

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Rollback()

	var deleteProductQuery string = `DELETE FROM product WHERE product_id = ?`
	fmt.Println(deleteProductQuery)
	_, err = db.Exec(deleteProductQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteDeviceQuery string = `DELETE FROM product_device WHERE product_id = ?`
	fmt.Println(deleteDeviceQuery)
	_, err = db.Exec(deleteDeviceQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteAuthenticationQuery string = `DELETE FROM authentication_details WHERE product_id = ?`
	fmt.Println(deleteAuthenticationQuery)
	_, err = db.Exec(deleteAuthenticationQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteDeveloperQuery string = `DELETE FROM product_developer WHERE product_id = ?`
	fmt.Println(deleteDeveloperQuery)
	_, err = db.Exec(deleteDeveloperQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteCustomizingQuery string = `DELETE FROM product_customizing WHERE product_id = ?`
	fmt.Println(deleteCustomizingQuery)
	_, err = db.Exec(deleteCustomizingQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var outPutID int32
	var outPutIdList = []int32{}
	var getOutPutIdQuery string = `SELECT output_id FROM product_output WHERE product_id = ?`
	rows, err := db.Query(getOutPutIdQuery, product_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&outPutID)
		if err != nil {
			log.Fatal(err)
		}

		outPutIdList = append(outPutIdList, outPutID)
	}

	fmt.Println("outPutIdList", outPutIdList)

	var deleteOutputQuery string = `DELETE FROM product_output WHERE product_id = ?`
	fmt.Println(deleteOutputQuery)
	_, err = db.Exec(deleteOutputQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	for i := 0; i < len(outPutIdList); i++ {
		var deleteAttachmentQuery string = `DELETE FROM output_attachment WHERE output_id = ?`
		fmt.Println(deleteAttachmentQuery)
		_, err = db.Exec(deleteAttachmentQuery, outPutIdList[i])
		if err != nil {
			fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
			log.Fatal(err)
		}
	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	result.ResultCode = 1

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, result)

}

func deleteSW(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("deleteSW()........")
	request.ParseForm()

	formData := request.Form

	//fmt.Println(formData)

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Rollback()

	var deleteProductQuery string = `DELETE FROM product WHERE product_id = ?`
	fmt.Println(deleteProductQuery)
	_, err = db.Exec(deleteProductQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteSwQuery string = `DELETE FROM product_sw WHERE product_id = ?`
	fmt.Println(deleteSwQuery)
	_, err = db.Exec(deleteSwQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteDeveloperQuery string = `DELETE FROM product_developer WHERE product_id = ?`
	fmt.Println(deleteDeveloperQuery)
	_, err = db.Exec(deleteDeveloperQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var deleteCustomizingQuery string = `DELETE FROM product_customizing WHERE product_id = ?`
	fmt.Println(deleteCustomizingQuery)
	_, err = db.Exec(deleteCustomizingQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	var outPutID int32
	var outPutIdList = []int32{}
	var getOutPutIdQuery string = `SELECT output_id FROM product_output WHERE product_id = ?`
	rows, err := db.Query(getOutPutIdQuery, product_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&outPutID)
		if err != nil {
			log.Fatal(err)
		}

		outPutIdList = append(outPutIdList, outPutID)
	}

	fmt.Println("outPutIdList", outPutIdList)

	var deleteOutputQuery string = `DELETE FROM product_output WHERE product_id = ?`
	fmt.Println(deleteOutputQuery)
	_, err = db.Exec(deleteOutputQuery, product_id)
	if err != nil {
		fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
		log.Fatal(err)
	}

	for i := 0; i < len(outPutIdList); i++ {
		var deleteAttachmentQuery string = `DELETE FROM output_attachment WHERE output_id = ?`
		fmt.Println(deleteAttachmentQuery)
		_, err = db.Exec(deleteAttachmentQuery, outPutIdList[i])
		if err != nil {
			fmt.Printf("-------------product_id가 %d인 제품 삭제 실패--------------", product_id)
			log.Fatal(err)
		}
	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	result.ResultCode = 1

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, result)
}
