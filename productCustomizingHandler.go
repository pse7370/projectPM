package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
											WHERE product_id = ?`

	rows, err := db.Query(getCustomizingListQuery, product_id)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 커스터마이징 내역 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	//fmt.Println(getCustomizingListQuery)
	defer rows.Close()

	var customizing_id int32
	var customizing_version string
	var function string
	var department string
	var employees_number int32
	var employees_name string
	var start_date sql.NullString
	var end_date sql.NullString

	for rows.Next() {
		err := rows.Scan(&customizing_id, &customizing_version,
			&function, &department, &employees_number, &employees_name,
			&start_date, &end_date)

		customizingList = append(customizingList,
			Product_customizing{
				Customizing_id:      customizing_id,
				Customizing_version: customizing_version,
				Customized_function: function,
				Department:          department,
				Employees_number:    employees_number,
				Employees_name:      employees_name,
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
