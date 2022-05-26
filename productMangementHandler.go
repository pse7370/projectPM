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
	fmt.Println(formData)
	fmt.Println(formData["@d#"])
	fmt.Println(formData["@d1#"+"product_name"])

	var product_device ProductDevice
	var authentication_detailsList Authentication_detailsList
	//var authentication_detailsList []Authentication_details
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
				authentication_detailsList.Auth_typeList = value

			case "one_to_one_max_user":
				authentication_detailsList.One_to_one_max_user = value

			case "one_to_many_max_user":
				authentication_detailsList.One_to_many_max_user = value

			case "one_to_one_max_template":
				authentication_detailsList.One_to_one_max_template = value

			case "one_to_many_max_template":
				authentication_detailsList.One_to_many_max_template = value

			}

		} // end if

	} // end for

	fmt.Println(product)
	fmt.Println(product_device)
	fmt.Println(authentication_detailsList)

	/*
		var product Product = Product{
			Product_id: nil,
			Product_type: formData[0][],
		}
		var authentication_details Authentication_details = Authentication_details{}
		var
	*/

}
