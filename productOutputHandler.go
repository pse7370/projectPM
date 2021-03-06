package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"archive/zip" // 파일을 압축할 수 있는 패키지

	"github.com/unrolled/render"
)

func getOutputList(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............getOutputList()...........")
	request.ParseForm()

	formData := request.Form
	//fmt.Println(formData)

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	formDataKey = "@d2#" + "pageNum"
	pageNumber, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("pageNumber : ", pageNumber)

	var getTotalOutputCountQuery string = `SELECT COUNT(output_id) FROM product_output WHERE product_id = ?`
	var outputCount int32
	err := db.QueryRow(getTotalOutputCountQuery, product_id).Scan(&outputCount)
	if err != nil {
		log.Println("outputCount 값 가져오기 오류", err)
	}

	var totalOutputCount OutputCount = OutputCount{TotalOutputCount: outputCount}

	var product_outputList []Product_output = []Product_output{}

	var getOutputListQuery string = `SELECT output_id, 
											output_type, 
											output_title, 
											write_date 
										FROM product_output
										WHERE product_id = ?
										ORDER BY output_id DESC
										OFFSET (?-1)*15 ROWS
										FETCH NEXT 15 ROWS ONLY`

	rows, err := db.Query(getOutputListQuery, product_id, pageNumber)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 산출물 목록 가져오기 실패===========\n", product_id)
		log.Println(err)
	}

	fmt.Println(getOutputListQuery)
	defer rows.Close()

	var output_id int32
	var output_type string
	var output_title string
	var write_date string

	for rows.Next() {
		err := rows.Scan(&output_id, &output_type, &output_title, &write_date)

		product_outputList = append(product_outputList,
			Product_output{
				Output_id:    output_id,
				Output_type:  output_type,
				Output_title: output_title,
				Write_date:   write_date,
			},
		)
		if err != nil {
			log.Fatal(err)
		}

	}

	var outputList OutputList = OutputList{
		Product_outputList: product_outputList,
		TotalOutputCount:   totalOutputCount,
	}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, outputList)

}

func addOutput(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1 << 30) // 1GiB 메모리 할당
	log.Println("...............addOutput().............")

	multipartForm := request.MultipartForm

	formData := multipartForm.Value
	//fmt.Println(formData)

	var product_output Product_output
	for key, value := range formData {

		splitRealKey := strings.Split(key, "#")

		if len(splitRealKey) >= 2 {
			switch splitRealKey[1] {
			case "product_id":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_output.Product_id = int32(temp)

			case "output_type":
				product_output.Output_type = value[0]

			case "output_title":
				product_output.Output_title = value[0]

			case "output_content":
				product_output.Output_content = value[0]

			case "write_date":
				product_output.Write_date = value[0]

			}

		} // end if

	} // end for

	fmt.Println("product_output : ", product_output)

	transaction, err := db.Begin()
	if err != nil {
		fmt.Println("--------트랜잭션 생성 오류---------")
		log.Fatal(err)
	}

	defer transaction.Rollback()

	_, err = db.Exec(`INSERT INTO product_output(product_id, output_type, output_title, output_content, write_date)
	VALUES (?, ?, ?, ?, ?)`,
		product_output.Product_id, product_output.Output_type, product_output.Output_title, product_output.Output_content, product_output.Write_date)
	if err != nil {
		fmt.Println("===========product 테이블 insert 실패===========")
		log.Fatal(err)
	}

	var output_id int32
	err = db.QueryRow("SELECT output_id FROM product_output WHERE product_id = ? AND output_title = ? AND write_date = ?",
		product_output.Product_id, product_output.Output_title, product_output.Write_date).Scan(&output_id)
	if err != nil {
		fmt.Println("===========product_output 테이블 output_id 가져오기 실패===========")
		log.Fatal(err)
	}
	fmt.Println("output_id : ", output_id)

	//var output_attachment Output_attachment
	var output_attachmentList []Output_attachment = []Output_attachment{}

	basePath, _ := os.Getwd()

	var attachmentSaveDir string = fmt.Sprintf("%s/%s/%d", basePath, product_output.Write_date, output_id)

	if _, err := os.Stat(attachmentSaveDir); os.IsNotExist(err) {
		err := os.MkdirAll(attachmentSaveDir, os.ModeDir)
		if err != nil {
			log.Println("------------폴더 생성 오류-------------")
			log.Fatalln(err)
		}
		fmt.Printf("==========해당 경로에 폴더가 없어 새로 생성 : %s", attachmentSaveDir)
	}

	fmt.Println(multipartForm.File)

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

		var attachmentSavePath string = attachmentSaveDir + "/" + fileHeader.Filename

		fileUpLoad, err := os.Create(attachmentSavePath)
		if err != nil {
			fmt.Println("파일 열기 실패 : ", err, "\n", attachmentSavePath)
			return
		}
		defer fileUpLoad.Close()

		_, err = io.Copy(fileUpLoad, file)
		if err != nil {
			fmt.Println("파일 복사 실패 : ", err)
			return
		}

		fmt.Println("파일 저장 성공!", fileHeader.Filename)
		output_attachmentList = append(output_attachmentList,
			Output_attachment{
				Output_id:      output_id,
				Real_file_name: fileHeader.Filename,
				Save_file_name: fileHeader.Filename,
				Save_path:      attachmentSavePath,
				File_size:      float64(fileHeader.Size),
			},
		)

		/*
			output_attachment.Output_id = output_id
			output_attachment.Real_file_name = fileHeader.Filename
			output_attachment.Save_file_name = fileHeader.Filename
			output_attachment.Save_path = attachmentSavePath
		*/

	} // end for

	fmt.Println("output_attachmentList : ", output_attachmentList)

	for i := 0; i < len(output_attachmentList); i++ {
		_, err = db.Exec(`INSERT INTO output_attachment(output_id, real_file_name, save_file_name, save_path, file_size)
		VALUES (?, ?, ?, ?, ?)`,
			output_attachmentList[i].Output_id,
			output_attachmentList[i].Real_file_name,
			output_attachmentList[i].Save_file_name,
			output_attachmentList[i].Save_path,
			output_attachmentList[i].File_size)
		if err != nil {
			fmt.Println("===========output_attachment 테이블 insert 실패===========")
			log.Fatal(err)
		}
	}

	// 트랜잭션 종료
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

func getOutputContent(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............getOutputContent()...........")
	request.ParseForm()

	formData := request.Form

	formDataKey := "@d1#" + "output_id"
	output_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("output_id : ", output_id)

	var product Product
	var product_output Product_output
	var output_attachmentList []Output_attachment = []Output_attachment{}

	var getOutputContentQuery string = `SELECT 
											po.output_type, 
											po.output_title, 
											po.output_content, 
											po.write_date, 
											p.product_name
										FROM product_output AS po
										LEFT JOIN product AS p
										ON po.product_id = p.product_id
										WHERE output_id = ?`

	err := db.QueryRow(getOutputContentQuery, output_id).Scan(&product_output.Output_type,
		&product_output.Output_title,
		&product_output.Output_content,
		&product_output.Write_date,
		&product.Product_name)
	if err != nil {
		fmt.Println("===========product_output 테이블 가져오기 실패===========")
		log.Fatal(err)
	}

	fmt.Println("product_name : ", product.Product_name)
	fmt.Println("product_output : ", product_output)

	var getAttachmentListQuery string = `SELECT real_file_name, 
												save_file_name, 
												save_path,
												file_size 
											FROM output_attachment
											WHERE output_id = ?`

	rows, err := db.Query(getAttachmentListQuery, output_id)
	if err != nil {
		log.Fatalf("==========output_id = %d인 산출물의 첨부파일 내역 가져오기 실패===========\n", output_id)
		log.Println(err)
	}

	//fmt.Println(getAttachmentListQuery)
	defer rows.Close()

	var real_file_name string
	var save_file_name string
	var save_path string
	var file_size float64

	for rows.Next() {
		err := rows.Scan(&real_file_name, &save_file_name, &save_path, &file_size)

		output_attachmentList = append(output_attachmentList,
			Output_attachment{
				Real_file_name: real_file_name,
				Save_file_name: save_file_name,
				Save_path:      save_path,
				File_size:      file_size,
			},
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("output_attachmentList : ", output_attachmentList)

	var outputContent OutputContent = OutputContent{
		Product:        product,
		Product_output: product_output,
		AttachmentList: output_attachmentList,
	}

	renderObj := render.New()
	renderObj.JSON(writer, http.StatusOK, outputContent)

}

func deleteOutput(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............deleteOutput()...........")

	requestURL := request.RequestURI

	splitURL := strings.Split(requestURL, "?")
	//fmt.Println(splitURL)

	var deleteOutputID []int32 = []int32{}
	if len(splitURL) > 1 {

		for i := 1; i < len(splitURL); i++ {
			temp, _ := strconv.ParseInt(splitURL[i], 10, 32)
			deleteOutputID = append(deleteOutputID, int32(temp))
		}

	}

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Rollback()

	var deleteOutputQuery string = `DELETE FROM product_output WHERE output_id = ?`
	var deleteAttachmentQuery string = `DELETE FROM output_attachment WHERE output_id = ?`
	for i := 0; i < len(deleteOutputID); i++ {

		_, err := db.Exec(deleteOutputQuery, deleteOutputID[i])
		if err != nil {
			fmt.Printf("-------------output_id가 %d인 제품 product_output 테이블 삭제 실패--------------", deleteOutputID[i])
			log.Fatal(err)
		}

		_, err = db.Exec(deleteAttachmentQuery, deleteOutputID[i])
		if err != nil {
			fmt.Printf("-------------output_id가 %d인 제품 output_attachment 테이블 삭제 실패--------------", deleteOutputID[i])
			log.Fatal(err)
		}

	}

	var resultcode ResultCode
	resultcode.ResultCode = 1

	var result Result = Result{resultcode}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, result)

}

func getSearchOutputList(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............getSearchOutputList()...........")
	request.ParseForm()

	formData := request.Form
	//fmt.Println(formData)

	formDataKey := "@d1#" + "product_id"
	product_id, _ := (strconv.ParseInt(formData[formDataKey][0], 10, 32))
	fmt.Println("product_id : ", product_id)

	formDataKey = "@d2#" + "searchCondition"
	searchCondition := formData[formDataKey][0]
	fmt.Println("searchCondition : ", searchCondition)

	formDataKey = "@d2#" + "searchText"
	searchText := "%" + formData[formDataKey][0] + "%"
	fmt.Println("searchText : ", searchText)

	var searchOutputQuery string
	switch searchCondition {
	case "산출물 종류":
		searchOutputQuery = `SELECT 
								po.output_id,
								po.output_type,
								po.output_title,
								po.write_date,
								c.*
							FROM product_output AS po,
							(SELECT COUNT(output_id) AS count FROM product_output WHERE product_id = ?
							AND output_type LIKE ?) AS c
							WHERE po.product_id = ?
							AND po.output_type LIKE ?
							ORDER BY po.output_id DESC`

	case "제목":
		searchOutputQuery = `SELECT 
								po.output_id,
								po.output_type,
								po.output_title,
								po.write_date,
								c.*
							FROM product_output AS po,
							(SELECT COUNT(output_id) AS count FROM product_output WHERE product_id = ?
							AND output_title LIKE ?) AS c
							WHERE po.product_id = ?
							AND po.output_title LIKE ?
							ORDER BY po.output_id DESC`

	case "내용":
		searchOutputQuery = `SELECT 
								po.output_id,
								po.output_type,
								po.output_title,
								po.write_date,
								c.*
							FROM product_output AS po,
							(SELECT COUNT(output_id) AS count FROM product_output WHERE product_id = ?
							AND output_content LIKE ?) AS c
							WHERE po.product_id = ?
							AND po.output_content LIKE ?
							ORDER BY po.output_id DESC`
	}

	rows, err := db.Query(searchOutputQuery, product_id, searchText, product_id, searchText)
	if err != nil {
		log.Fatalf("==========product_id = %d인 제품의 산출물 %s에서 %s 검색 실패===========\n", product_id, searchCondition, searchText)
		log.Println(err)
	}

	defer rows.Close()

	var totalOutputCount OutputCount

	var product_outputList []Product_output = []Product_output{}

	var outputCount int32
	var output_id int32
	var output_type string
	var output_title string
	var write_date string

	for rows.Next() {
		err := rows.Scan(&output_id, &output_type, &output_title, &write_date, &outputCount)

		product_outputList = append(product_outputList,
			Product_output{
				Output_id:    output_id,
				Output_type:  output_type,
				Output_title: output_title,
				Write_date:   write_date,
			},
		)

		if totalOutputCount.TotalOutputCount == 0 {
			totalOutputCount = OutputCount{TotalOutputCount: outputCount}
		}

		if err != nil {
			log.Fatal(err)
		}

	}

	var outputList OutputList = OutputList{
		Product_outputList: product_outputList,
		TotalOutputCount:   totalOutputCount,
	}

	renderObj := render.New()

	renderObj.JSON(writer, http.StatusOK, outputList)

}

func downloadAttachment(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............downloadAttachment()...........")

	request.ParseForm()

	formData := request.Form
	//fmt.Println(formData)

	formDataKey := "@d1#" + "file_name"
	fileName := formData[formDataKey][0]

	formDataKey = "@d1#" + "save_path"
	saveFilePath := formData[formDataKey][0]

	fmt.Println("fileName : ", fileName)
	fmt.Println("saveFilePath : ", saveFilePath)

	file, err := os.Open(saveFilePath)
	if err != nil {
		http.Error(writer, fileName+"의 파일을 찾을 수 없습니다.", 404)
		return
	}
	defer file.Close()

	fileHeader := make([]byte, 512)
	file.Read(fileHeader)

	fileStat, _ := file.Stat()

	writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	writer.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	writer.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

	file.Seek(0, 0)
	io.Copy(writer, file)

}

func downloadAttachmentList(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............downloadAttachment()...........")

	request.ParseForm()

	formData := request.Form
	//fmt.Println(formData)

	formDataKey := "@d1#" + "file_name"
	fileNameList := formData[formDataKey]

	formDataKey = "@d1#" + "save_path"
	saveFilePathList := formData[formDataKey]

	formDataKey = "@d2#" + "product_name"
	product_name := formData[formDataKey][0]

	formDataKey = "@d2#" + "output_title"
	output_title := formData[formDataKey][0]

	zipFileName := "[" + product_name + "]" + output_title + ".zip"

	// 산출물에서 선택한 첨부파일들을 찾아 알집으로 압축하기
	zipArchive, _ := os.Create(zipFileName)
	defer zipArchive.Close()

	zipWriter := zip.NewWriter(zipArchive)

	var file *os.File
	var fileWriter io.Writer
	for i := 0; i < len(saveFilePathList); i++ {
		file, _ = os.Open(saveFilePathList[i])
		fileWriter, _ = zipWriter.Create(fileNameList[i])
		io.Copy(fileWriter, file)
	}

	zipWriter.Close()

	// 위에서 압축한 알집을 브라우저로 전달 다운로드
	var err error
	file, err = os.Open(zipFileName)
	if err != nil {
		http.Error(writer, zipFileName+"의 파일을 찾을 수 없습니다.", 404)
		return
	}

	fileHeader := make([]byte, 512)
	file.Read(fileHeader)

	writer.Header().Set("Content-Disposition", "attachment; filename="+zipFileName)
	writer.Header().Set("Content-Type", "applicaiton/zip")
	file.Seek(0, 0)
	io.Copy(writer, file)

	defer file.Close()

}

func deleteZipFile(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............deleteZipFile()...........")

	request.ParseForm()

	formData := request.Form
	fmt.Println(formData)

	formDataKey := "@d1#" + "product_name"
	product_name := formData[formDataKey][0]

	formDataKey = "@d1#" + "output_title"
	output_title := formData[formDataKey][0]

	zipFileName := "[" + product_name + "]" + output_title + ".zip"

	basePath, _ := os.Getwd()
	zip_save_path := basePath + "/" + zipFileName

	err := os.Remove(zip_save_path)
	if err != nil {
		fmt.Println("알집 파일 삭제 실패 ", err)
	}

}

func modifyOutput(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("............modifyOutput()...........")
	request.ParseMultipartForm(1 << 30)

	multipartForm := request.MultipartForm
	fmt.Println(multipartForm)

	formData := multipartForm.Value

	var product_output Product_output
	var deleteFileList DeleteFileList
	for key, value := range formData {

		splitRealKey := strings.Split(key, "#")

		if len(splitRealKey) >= 2 {
			switch splitRealKey[1] {
			case "output_id":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_output.Output_id = int32(temp)

			case "product_id":
				temp, _ := strconv.ParseInt(value[0], 10, 32)
				product_output.Product_id = int32(temp)

			case "output_type":
				product_output.Output_type = value[0]

			case "output_title":
				product_output.Output_title = value[0]

			case "output_content":
				product_output.Output_content = value[0]

			case "write_date":
				product_output.Write_date = value[0]

			case "delete_file_name":
				deleteFileList.DeleteFileNameList = value

			}

		} // end if

	} // end for

	fmt.Println("product_output : ", product_output)
	fmt.Println("deleteFileList : ", deleteFileList)

	transaction, err := db.Begin()
	if err != nil {
		fmt.Println("--------트랜잭션 생성 오류---------")
		log.Fatal(err)
	}

	defer transaction.Rollback()

	_, err = db.Exec(`UPDATE product_output 
							SET output_type = ?,
								output_title = ?,
								output_content = ?
							WHERE output_id = ?`,
		product_output.Output_type,
		product_output.Output_title,
		product_output.Output_content,
		product_output.Output_id)
	if err != nil {
		fmt.Printf("===========product_output 테이블 output_id는 %d update 실패===========\n", product_output.Output_id)
		log.Fatal(err)
	}

	for i := 0; i < len(deleteFileList.DeleteFileNameList); i++ {
		_, err = db.Exec(`DELETE FROM output_attachment WHERE output_id = ? AND real_file_name = ?`,
			product_output.Output_id, deleteFileList.DeleteFileNameList[i])

		if err != nil {
			fmt.Printf("===========output_attachment 테이블 output_id는 %d, real_file_name은 %s  delete 실패===========\n", product_output.Output_id, deleteFileList.DeleteFileNameList[i])
			log.Fatal(err)
		}
	}

	var output_attachmentList []Output_attachment = []Output_attachment{}

	basePath, _ := os.Getwd()

	splitWrite_date := strings.Split(product_output.Write_date, "T")[0]
	var attachmentSaveDir string = fmt.Sprintf("%s/%s/%d", basePath, splitWrite_date, product_output.Output_id)

	if _, err := os.Stat(attachmentSaveDir); os.IsNotExist(err) {
		err := os.MkdirAll(attachmentSaveDir, os.ModeDir)
		if err != nil {
			log.Println("------------폴더 생성 오류-------------")
			log.Fatalln(err)
		}
		fmt.Printf("==========해당 경로에 폴더가 없어 새로 생성 : %s", attachmentSaveDir)
	}

	fmt.Println(multipartForm.File)

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

		var attachmentSavePath string = attachmentSaveDir + "/" + fileHeader.Filename

		fileUpLoad, err := os.Create(attachmentSavePath)
		if err != nil {
			fmt.Println("파일 열기 실패 : ", err, "\n", attachmentSavePath)
			return
		}
		defer fileUpLoad.Close()

		_, err = io.Copy(fileUpLoad, file)
		if err != nil {
			fmt.Println("파일 복사 실패 : ", err)
			return
		}

		fmt.Println("파일 저장 성공!", fileHeader.Filename)
		output_attachmentList = append(output_attachmentList,
			Output_attachment{
				Output_id:      product_output.Output_id,
				Real_file_name: fileHeader.Filename,
				Save_file_name: fileHeader.Filename,
				Save_path:      attachmentSavePath,
				File_size:      float64(fileHeader.Size),
			},
		)

	} // end for

	fmt.Println("output_attachmentList : ", output_attachmentList)

	for i := 0; i < len(output_attachmentList); i++ {
		_, err = db.Exec(`INSERT INTO output_attachment(output_id, real_file_name, save_file_name, save_path, file_size)
		VALUES (?, ?, ?, ?, ?)`,
			output_attachmentList[i].Output_id,
			output_attachmentList[i].Real_file_name,
			output_attachmentList[i].Save_file_name,
			output_attachmentList[i].Save_path,
			output_attachmentList[i].File_size)
		if err != nil {
			fmt.Println("===========output_attachment 테이블 insert 실패===========")
			log.Fatal(err)
		}
	}

	// 트랜잭션 종료
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
