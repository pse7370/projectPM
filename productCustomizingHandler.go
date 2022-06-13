package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/unrolled/render"
)

func getCustomizingList(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............getCustomizingList()...........")
	request.ParseForm()

	formData := request.Form
	fmt.Println(formData)

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	var product Product
	var customizingList []Product_customizing = []Product_customizing{}

	var product_name string
	var getProductNameQuery string = `SELECT product_name FROM product WHERE product_id = ?`
	err := db.QueryRow(getProductNameQuery, product_id).Scan(&product_name)
	if err != nil {
		log.Println("product_name 값 가져오기 오류", err)
	}

	product.Product_name = product_name

	var getCustomizingListQuery string = `SELECT customizing_id, 
												customizing_version, 
												customized_function, 
												department, 
												employees_number, 
												employees_name, 
												start_dates, 
												end_date 
											FROM product_customizing 
											WHERE product_id = ?
											ORDER BY customizing_version ASC, customizing_id ASC`

	rows, err := db.Query(getCustomizingListQuery, product_id)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 커스터마이징 내역 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	//fmt.Println(getCustomizingListQuery)
	defer rows.Close()

	var customizing_id int32
	var customizing_version sql.NullString
	var function sql.NullString
	var department sql.NullString
	var employees_number sql.NullInt32
	var employees_name sql.NullString
	var start_date sql.NullString
	var end_date sql.NullString

	for rows.Next() {
		err := rows.Scan(&customizing_id, &customizing_version,
			&function, &department, &employees_number, &employees_name,
			&start_date, &end_date)

		customizingList = append(customizingList,
			Product_customizing{
				Customizing_id:      customizing_id,
				Customizing_version: customizing_version.String,
				Customized_function: function.String,
				Department:          department.String,
				Employees_number:    employees_number.Int32,
				Employees_name:      employees_name.String,
				Start_date:          start_date.String,
				End_date:            end_date.String,
			},
		)
		if err != nil {
			log.Fatal(err)
		}

	}

	var product_customizingList CustomizingList = CustomizingList{
		Product:                 product,
		Product_customizingList: customizingList,
	}
	fmt.Println(product_customizingList)

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, product_customizingList)

}

func modifyCustomizing(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............modifyCustomizing()...........")

	request.ParseForm()

	formData := request.Form
	//fmt.Println(formData)

	var product_id int32
	var product_customizingList Product_customizingList
	var delete_customizingIdList Delete_customizingIdList

	for key, value := range formData {
		//fmt.Println(key, "/", value)

		splitRealKey := strings.Split(key, "#")

		if len(splitRealKey) >= 2 {
			//fmt.Println(splitRealKey)

			switch splitRealKey[1] {
			case "product_id":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_id = int32(temp)

			case "customizing_id":
				product_customizingList.Customizing_id = value

			case "customizing_version":
				product_customizingList.Customizing_version = value

			case "customized_function":
				product_customizingList.Customized_function = value

			case "department":
				product_customizingList.Department = value

			case "employees_number":
				product_customizingList.Employees_number = value

			case "employees_name":
				product_customizingList.Employees_name = value

			case "start_date":
				product_customizingList.Start_date = value

			case "end_date":
				product_customizingList.End_date = value

			case "delete_customizing_id":
				delete_customizingIdList.Delete_customizing_id = value

			}

		} // end if

	} // end for

	fmt.Println("product_id : ", product_id)
	fmt.Println("product_customizingList : ", product_customizingList)
	fmt.Println("delete_customizingIdList : ", delete_customizingIdList)

	var product_customizing []Product_customizing = []Product_customizing{}
	sliceLength := len(product_customizingList.Customizing_version)
	for i := 0; i < sliceLength; i++ {
		tempId, _ := strconv.ParseInt(product_customizingList.Customizing_id[i], 10, 32)
		tmepNum, _ := strconv.ParseInt(product_customizingList.Employees_number[i], 10, 32)
		product_customizing = append(product_customizing,
			Product_customizing{
				Customizing_id:      int32(tempId),
				Customizing_version: product_customizingList.Customizing_version[i],
				Customized_function: product_customizingList.Customized_function[i],
				Department:          product_customizingList.Department[i],
				Employees_number:    int32(tmepNum),
				Employees_name:      product_customizingList.Employees_name[i],
				Start_date:          product_customizingList.Start_date[i],
				End_date:            product_customizingList.End_date[i],
			},
		)
	}

	var delete_customizing_id []Delete_customizingId = []Delete_customizingId{}
	sliceLength = len(delete_customizingIdList.Delete_customizing_id)
	for i := 0; i < sliceLength; i++ {
		tempId, _ := strconv.ParseInt(delete_customizingIdList.Delete_customizing_id[i], 10, 32)
		delete_customizing_id = append(delete_customizing_id,
			Delete_customizingId{
				Delete_customizing_id: int32(tempId),
			},
		)
	}

	transaction, err := db.Begin()
	if err != nil {
		fmt.Println("--------트랜잭션 생성 오류---------")
		log.Fatal(err)
	}

	defer transaction.Rollback()

	var insertCustomizingQuery string = `INSERT INTO product_customizing VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	var updateCustomizingQuery string = `UPDATE product_customizing 
												SET customizing_version = ?,
													customized_function = ?,
													department = ?,
													employees_number = ?,
													employees_name = ?,
													start_dates = ?,
													end_date = ?
												WHERE customizing_id = ?`

	for i := 0; i < len(product_customizing); i++ {
		switch product_customizing[i].Customizing_id {
		case 0:
			_, err := db.Exec(insertCustomizingQuery,
				product_id, product_customizing[i].Customizing_version,
				product_customizing[i].Customized_function, product_customizing[i].Department,
				product_customizing[i].Employees_number, product_customizing[i].Employees_name,
				product_customizing[i].Start_date, product_customizing[i].End_date)

			if err != nil {
				fmt.Printf("===========product_customizing 테이블 insert 실패, product_id는 %d============\n", product_id)
				log.Fatal(err)
			}

		default:
			_, err := db.Exec(updateCustomizingQuery,
				product_customizing[i].Customizing_version,
				product_customizing[i].Customized_function, product_customizing[i].Department,
				product_customizing[i].Employees_number, product_customizing[i].Employees_name,
				product_customizing[i].Start_date, product_customizing[i].End_date,
				product_customizing[i].Customizing_id)

			if err != nil {
				fmt.Printf("===========product_customizing 테이블 update 실패, customizing_id는 %d============\n", product_customizing[i].Customizing_id)
				log.Fatal(err)
			}
		}
	} // end for

	var deleteCustomizingQuery string = `DELETE FROM product_customizing WHERE customizing_id = ?`
	for i := 0; i < len(delete_customizing_id); i++ {
		_, err = db.Exec(deleteCustomizingQuery, delete_customizing_id[i].Delete_customizing_id)
		if err != nil {
			fmt.Printf("-------------customizing_id %d인 제품 product_customizing 테이블 삭제 실패--------------\n", delete_customizing_id[i].Delete_customizing_id)
			log.Fatal(err)
		}
	}

	var resultcode ResultCode
	resultcode.ResultCode = 1

	var result Result = Result{resultcode}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, result)

}

func deleteCustomizing(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............deleteCustomizing()...........")

	request.ParseForm()

	formData := request.Form
	fmt.Println(formData)
	formDataKey := "@d1#" + "customizing_id"
	data := formData[formDataKey]

	var customizing_ids []int32 = []int32{}
	if len(data) > 1 {

		for i := 1; i < len(data); i++ {
			stringCustomizingID := data[i]
			int64CustomizingID, _ := strconv.ParseInt(stringCustomizingID, 10, 32)
			customizing_ids = append(customizing_ids, int32(int64CustomizingID))
		}

	}

	fmt.Println("customizing_ids : ", customizing_ids)

	/*
		requestURL := request.RequestURI

		splitURL := strings.Split(requestURL, "?")

		var customizing_ids []int32 = []int32{}
		if len(splitURL) > 1 {

			for i := 1; i < len(splitURL); i++ {
				stringCustomizingID := splitURL[i]
				int64CustomizingID, _ := strconv.ParseInt(stringCustomizingID, 10, 32)
				customizing_ids = append(customizing_ids, int32(int64CustomizingID))
			}

		}

		fmt.Println("customizing_ids : ", customizing_ids)
	*/

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Rollback()

	for i := 0; i < len(customizing_ids); i++ {
		var deleteCustomizingQuery string = `DELETE FROM product_customizing WHERE customizing_id = ?`
		_, err = db.Exec(deleteCustomizingQuery, customizing_ids[i])
		if err != nil {
			fmt.Printf("-------------customizing_id %d인 제품 product_customizing 테이블 삭제 실패--------------\n", customizing_ids[i])
			log.Fatal(err)
		}

	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}

	var resultcode ResultCode
	resultcode.ResultCode = 1

	var result Result = Result{resultcode}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, result)

}
