package main

// 상품 테이블 구조체
type Product struct {
	Product_id      int32  `json:"product_id"`
	Product_type    string `json:"product_type"`
	Product_name    string `json:"product_name"`
	Product_version string `json:"product_version"`
	Real_image_name string `json:"real_image_name"`
	Save_image_name string `json:"save_image_name"`
	Save_path       string `json:"save_path"`
	Explanation     string `json:"explanation"`
}

// 출입통제기 테이블 구조체
type ProductDevice struct {
	Device_detail_id int32   `json:"device_detail_id"`
	Product_id       int32   `json:"product_id"`
	Width            float64 `json:"width"`
	Height           float64 `json:"height"`
	Depth            float64 `json:"depth"`
	Ip_ratings       string  `json:"ip_ratings"`
	Server           string  `json:"server"`
	Wi_fi            string  `json:"wi_fi"`
	Other            string  `json:"other"`
}

// 인증 상세 테이블 구조체
/*
type Authentication_details struct {
	Auth_detail_id int32  `json:"auth_detail_id"`
	Product_id     int32  `json:"product_id"`
	Auth_type      string `json:"auth_type"`
	Auth_method    string `json:"auth_method"`
	Max_users      int64  `json:"max_users"`
	Max_templates  int64  `json:"max_templates"`
}
*/

type Authentication_details struct {
	Auth_detail_id           int32  `json:"auth_detail_id"`
	Product_id               int32  `json:"product_id"`
	Auth_type                string `json:"auth_type"`
	One_to_one_max_user      int64  `json:"one_to_one_max_user"`
	One_to_many_max_user     int64  `json:"one_to_many_max_user"`
	One_to_one_max_template  int64  `json:"one_to_one_max_template"`
	One_to_many_max_template int64  `json:"one_to_many_max_template"`
}

// SW 테이블 구조체
type ProductSW struct {
	SW_detail_id int32 `json:"sw_detail_id"`
	Product_id   int32 `json:"product_id"`
	Simultaneous uint8 `json:"simultaneous"`
	Available_db uint8 `json:"available_db"`
	Available_os uint8 `json:"available_os"`
}

// 담당 개발자 테이블 구조체
type Product_developer struct {
	Developer_id     int32  `json:"developer_id"`
	Product_id       int32  `json:"product_id"`
	Department       string `json:"department"`
	Employees_number int64  `json:"employees_number"`
	Employees_name   string `json:"Employees_name"`
	Start_date       string `json:"start_date"`
	End_date         string `json:"end_date"`
}

// 커스터마이징 상세 테이블 구조체
type Product_customizing struct {
	Customizing_id      int32  `json:"customizing_id"`
	Product_id          int32  `json:"product_id"`
	Customizing_version string `json:"customizing_version"`
	Customized_function string `json:"customized_function"`
	Department          string `json:"department"`
	Employees_number    int32  `json:"employees_number"`
	Employees_name      string `json:"Employees_name"`
	Start_date          string `json:"start_date"`
	End_date            string `json:"end_date"`
}

// 산출물 상세 테이블 구조체
type Product_output struct {
	Output_id      int32  `json:"output_id"`
	Product_id     int32  `json:"product_id"`
	Output_type    string `json:"output_type"`
	Output_title   string `json:"output_title"`
	Output_content string `json:"output_content"`
}

// 산출물 첨부파일 테이블 구조체
type Output_attachment struct {
	Attachment_id  int32  `json:"attachment_id"`
	Output_id      int32  `json:"output_id"`
	Real_file_name string `json:"real_file_name"`
	Save_file_name string `json:"save_file_name"`
	Save_path      string `json:"save_path"`
}

//============ResultCode 전달 구조체
type Result struct {
	ResultCode int `json:"resultCode"`
}

//==================사이드 메뉴(트리) 구성을 위한 구조체===============================

type SideMenuContent struct {
	Label      string `json:"label"`
	Value      string `json:"value"`
	Parent     string `json:"parent"`
	Product_id int32  `json:"product_id"`
}

type SideMenu struct {
	SideMenuList []SideMenuContent `json:"sideMenuList"`
}

//====================출입통제기 등록을 위한 구조체==============================

// 인증 방식 표의 파라미터값 받기, 전달을 위한 구조체
/*
type Authentication_detailsList struct {
	Authentication_detailsList []Authentication_details `json:"authentication_detailsList"`
}

type Product_developerList struct {
	Product_developerList []Product_developer `json:"product_developerList"`
}
*/

/*
type Authentication_detailsList struct {
	Auth_typeList            []string `json:"auth_typeList"`
	One_to_one_max_user      []string `json:"one_to_one_max_user"`
	One_to_many_max_user     []string `json:"one_to_many_max_user"`
	One_to_one_max_template  []string `json:"one_to_one_max_template"`
	One_to_many_max_template []string `json:"one_to_many_max_template"`
}
*/

type Product_developerList struct {
	DepartmentList       []string `json:"departmentList"`
	Employees_numberList []string `json:"employees_numberList"`
	Employees_nameList   []string `json:"employees_nameList"`
	Start_dateList       []string `json:"start_dateList"`
	End_dateList         []string `json:"end_dateList"`
}

// 인증 상세 테이블 구조체
/*
type Authentication_detailsList struct {
	Auth_type     []string `json:"auth_type"`
	Auth_method   []string `json:"auth_method"`
	Max_users     []string `json:"max_users"`
	Max_templates []string `json:"max_templates"`
}
*/

type Authentication_detailsList struct {
	Auth_type                []string `json:"auth_type"`
	One_to_one_max_user      []string `json:"one_to_one_max_user"`
	One_to_many_max_user     []string `json:"one_to_many_max_user"`
	One_to_one_max_template  []string `json:"one_to_one_max_template"`
	One_to_many_max_template []string `json:"one_to_many_max_template"`
}
