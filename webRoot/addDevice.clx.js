/*
 * App URI: addDevice
 * Source Location: addDevice.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("addDevice", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * addDevice.js
			 * Created at 2022. 5. 25. 오후 1:03:39.
			 *
			 * @author A
			 ************************************************/
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				// 제품명 inputBox에 커서 위치
				app.lookup("input_productName").focus();
				
				// 통신 방식 그리드 1행 추가
				app.lookup("communication").insertRow(1, true);
				
			}
			
			
			/*
			 * 담당개발자 "+" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick2(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				var grid_developer = app.lookup("grid_developer");
				var insertRow = grid_developer.insertRow(1, true);
				// + 버튼 클릭시 그리드 행 추가
			}
			
			/*
			 * "등록" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				var product_image = app.lookup("product_image");
				
				var addDevice = app.lookup("addDevice");
				// 업로드한 파일 정보 서브미션에 저장
				if (product_image.file != null){
					addDevice.addFileParameter("deviceImage", product_image.file);
					console.log("출입통제기 이미지 파일명 :" + product_image.file);
					console.log("출입통제기 파일 타입" + product_image.file.type);
				}
				
				var product = app.lookup("product");
				
				product.setValue("product_type", "출입통제기");
				
				var authenticationList = app.lookup("authenticationList")
				
				console.log(authenticationList.getRowCount());
				
				var authenGrid = app.lookup("authentication");
				console.log("checkedRow : " + authenGrid.getCheckRowIndices());
				
				var checkedRow = authenGrid.getCheckRowIndices();
				
				addDevice.send();
			}
			
			
			/* 출입통제기 제품 등록 성공 후
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onAddDeviceSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var addDevice = e.control;
				console.log(window.status);
				
				var resultCode = app.lookup("result").getValue("resultCode");
				console.log(resultCode);
				app.setAppProperty("resultCode", resultCode);
				
				app.getRootAppInstance().dialogManager.getDialogByName("addProduct").close(1);
				
				/*
				if(resultCode == 1){
					app.getRootAppInstance().dialogManager.getDialogByName("addProduct").close();
				}else {
					alert("상품 등록 실패");
				}
				*/
				
				//app.getRootAppInstance().dialogManager.getDialogByName("addProduct").close();
				//app.getRootAppInstance()는 최상위 앱 호출
				
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("authenticationList");
			dataSet_1.parseData({
				"columns": [
					{"name": "auth_type"},
					{
						"name": "one_to_one_max_user",
						"dataType": "number"
					},
					{
						"name": "one_to_many_max_user",
						"dataType": "number"
					},
					{
						"name": "one_to_one_max_template",
						"dataType": "number"
					},
					{
						"name": "one_to_many_max_template",
						"dataType": "number"
					}
				],
				"rows": [
					{"auth_type": "카드", "one_to_one_max_user": "", "one_to_many_max_user": "", "one_to_one_max_template": "", "one_to_many_max_template": ""},
					{"auth_type": "지문", "one_to_one_max_user": "", "one_to_many_max_user": "", "one_to_one_max_template": "", "one_to_many_max_template": ""},
					{"auth_type": "얼굴", "one_to_one_max_user": "", "one_to_many_max_user": "", "one_to_one_max_template": "", "one_to_many_max_template": ""},
					{"auth_type": "홍채", "one_to_one_max_user": "", "one_to_many_max_user": "", "one_to_one_max_template": "", "one_to_many_max_template": ""}
				]
			});
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("developerList");
			dataSet_2.parseData({
				"columns" : [
					{"name": "department"},
					{"name": "employees_number"},
					{"name": "employees_name"},
					{"name": "start_date"},
					{"name": "end_date"}
				]
			});
			app.register(dataSet_2);
			var dataMap_1 = new cpr.data.DataMap("product_device");
			dataMap_1.parseData({
				"columns" : [
					{
						"name": "width",
						"dataType": "number"
					},
					{
						"name": "height",
						"dataType": "number"
					},
					{
						"name": "depth",
						"dataType": "number"
					},
					{"name": "ip_ratings"},
					{"name": "server"},
					{"name": "wi_fi"},
					{"name": "other"}
				]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("product");
			dataMap_2.parseData({
				"columns" : [
					{"name": "product_type"},
					{"name": "product_name"},
					{"name": "product_version"},
					{"name": "explanation"}
				]
			});
			app.register(dataMap_2);
			
			var dataMap_3 = new cpr.data.DataMap("result");
			dataMap_3.parseData({
				"columns" : [{"name": "resultCode"}]
			});
			app.register(dataMap_3);
			var submission_1 = new cpr.protocols.Submission("addDevice");
			submission_1.action = "/productMangement/addDevice";
			submission_1.mediaType = "multipart/form-data";
			submission_1.addRequestData(dataMap_2);
			submission_1.addRequestData(dataMap_1);
			submission_1.addRequestData(dataSet_2);
			submission_1.addRequestData(dataSet_1);
			submission_1.addResponseData(dataMap_3, false);
			if(typeof onAddDeviceSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onAddDeviceSubmitDone);
			}
			app.register(submission_1);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 707px) and (max-width: 1023px)", "dialog");
			app.supportMedia("all and (min-width: 500px) and (max-width: 706px)", "tablet");
			app.supportMedia("all and (max-width: 499px)", "mobile");
			
			// Configure root container
			var container = app.getContainer();
			container.style.css({
				"width" : "100%",
				"top" : "0px",
				"height" : "100%",
				"left" : "0px"
			});
			
			// Layout
			var xYLayout_1 = new cpr.controls.layouts.XYLayout();
			container.setLayout(xYLayout_1);
			
			// UI Configuration
			var group_1 = new cpr.controls.Container("device_detail");
			// Layout
			var xYLayout_2 = new cpr.controls.layouts.XYLayout();
			group_1.setLayout(xYLayout_2);
			(function(container){
				var group_2 = new cpr.controls.Container();
				// Layout
				var formLayout_1 = new cpr.controls.layouts.FormLayout();
				formLayout_1.topMargin = "0px";
				formLayout_1.rightMargin = "0px";
				formLayout_1.bottomMargin = "0px";
				formLayout_1.leftMargin = "0px";
				formLayout_1.horizontalSpacing = "15px";
				formLayout_1.verticalSpacing = "0px";
				formLayout_1.setColumns(["227px", "238px", "209px"]);
				formLayout_1.setRows(["25px"]);
				group_2.setLayout(formLayout_1);
				(function(container){
					var group_3 = new cpr.controls.Container();
					// Layout
					var xYLayout_3 = new cpr.controls.layouts.XYLayout();
					group_3.setLayout(xYLayout_3);
					(function(container){
						var output_1 = new cpr.controls.Output();
						output_1.value = "제품명";
						output_1.style.css({
							"vertical-align" : "middle",
							"text-align" : "center"
						});
						container.addChild(output_1, {
							"top": "0px",
							"left": "7px",
							"width": "73px",
							"height": "25px"
						});
						var inputBox_1 = new cpr.controls.InputBox("input_productName");
						inputBox_1.bind("value").toDataMap(app.lookup("product"), "product_name");
						container.addChild(inputBox_1, {
							"top": "0px",
							"left": "79px",
							"width": "159px",
							"height": "25px"
						});
					})(group_3);
					container.addChild(group_3, {
						"colIndex": 1,
						"rowIndex": 0
					});
				})(group_2);
				container.addChild(group_2, {
					"top": "0px",
					"left": "0px",
					"width": "704px",
					"height": "27px"
				});
				var grid_1 = new cpr.controls.Grid("authentication");
				grid_1.init({
					"dataSet": app.lookup("authenticationList"),
					"columns": [
						{"width": "25px"},
						{"width": "25px"},
						{"width": "100px"},
						{"width": "100px"},
						{"width": "100px"},
						{"width": "100px"},
						{"width": "100px"}
					],
					"header": {
						"rows": [
							{"height": "24px"},
							{"height": "24px"}
						],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0, "rowSpan": 2, "colSpan": 2},
								"configurator": function(cell){
									cell.filterable = false;
									cell.sortable = false;
									cell.columnType = "checkbox";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2, "rowSpan": 2, "colSpan": 1},
								"configurator": function(cell){
									cell.targetColumnName = "auth_type";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "인증 방식";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 3},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_one_max_user";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : 1";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 4},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_many_max_user";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : N";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 5},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_one_max_template";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : 1";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 6},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_many_max_template";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : N";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3, "rowSpan": 1, "colSpan": 2},
								"configurator": function(cell){
									cell.text = "최대 등록 가능 사용자 수";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 5, "rowSpan": 1, "colSpan": 2},
								"configurator": function(cell){
									cell.text = "최대 등록 가능 템플릿 수";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							}
						]
					},
					"detail": {
						"rows": [{"height": "24px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0, "rowSpan": 1, "colSpan": 2},
								"configurator": function(cell){
									cell.columnType = "checkbox";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.columnName = "auth_type";
									cell.control = (function(){
										var inputBox_2 = new cpr.controls.InputBox("ipb1");
										inputBox_2.readOnly = true;
										inputBox_2.bind("value").toDataColumn("auth_type");
										return inputBox_2;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.columnName = "one_to_one_max_user";
									cell.control = (function(){
										var numberEditor_1 = new cpr.controls.NumberEditor("nbe1");
										numberEditor_1.style.css({
											"text-align" : "right",
											"padding-right" : "5px"
										});
										numberEditor_1.bind("value").toDataColumn("one_to_one_max_user");
										return numberEditor_1;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.columnName = "one_to_many_max_user";
									cell.control = (function(){
										var numberEditor_2 = new cpr.controls.NumberEditor("nbe2");
										numberEditor_2.style.css({
											"text-align" : "right",
											"padding-right" : "5px"
										});
										numberEditor_2.bind("value").toDataColumn("one_to_many_max_user");
										return numberEditor_2;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 5},
								"configurator": function(cell){
									cell.columnName = "one_to_one_max_template";
									cell.control = (function(){
										var numberEditor_3 = new cpr.controls.NumberEditor("nbe3");
										numberEditor_3.style.css({
											"text-align" : "right",
											"padding-right" : "5px"
										});
										numberEditor_3.bind("value").toDataColumn("one_to_one_max_template");
										return numberEditor_3;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 6},
								"configurator": function(cell){
									cell.columnName = "one_to_many_max_template";
									cell.control = (function(){
										var numberEditor_4 = new cpr.controls.NumberEditor("nbe4");
										numberEditor_4.style.css({
											"text-align" : "right",
											"padding-right" : "5px"
										});
										numberEditor_4.bind("value").toDataColumn("one_to_many_max_template");
										return numberEditor_4;
									})();
								}
							}
						]
					}
				});
				grid_1.style.css({
					"text-align" : "center"
				});
				container.addChild(grid_1, {
					"top": "45px",
					"left": "0px",
					"width": "704px",
					"height": "146px"
				});
				var group_4 = new cpr.controls.Container();
				// Layout
				var formLayout_2 = new cpr.controls.layouts.FormLayout();
				formLayout_2.topMargin = "0px";
				formLayout_2.rightMargin = "0px";
				formLayout_2.bottomMargin = "0px";
				formLayout_2.leftMargin = "0px";
				formLayout_2.horizontalSpacing = "0px";
				formLayout_2.verticalSpacing = "0px";
				formLayout_2.setColumns(["2fr", "2fr", "42px", "2fr", "42px", "2fr", "52px"]);
				formLayout_2.setRows(["22px", "25px"]);
				group_4.setLayout(formLayout_2);
				(function(container){
					var output_2 = new cpr.controls.Output();
					output_2.value = "제품 사이즈";
					output_2.style.css({
						"vertical-align" : "middle",
						"text-align" : "center"
					});
					container.addChild(output_2, {
						"colIndex": 0,
						"rowIndex": 1
					});
					var output_3 = new cpr.controls.Output();
					output_3.value = "(W) x";
					output_3.style.css({
						"vertical-align" : "middle",
						"text-align" : "center"
					});
					container.addChild(output_3, {
						"colIndex": 2,
						"rowIndex": 1
					});
					var output_4 = new cpr.controls.Output();
					output_4.value = "(H) x";
					output_4.style.css({
						"vertical-align" : "middle",
						"text-align" : "center"
					});
					container.addChild(output_4, {
						"colIndex": 4,
						"rowIndex": 1,
						"colSpan": 1,
						"rowSpan": 1
					});
					var output_5 = new cpr.controls.Output();
					output_5.value = "(D)mm";
					output_5.style.css({
						"vertical-align" : "middle",
						"text-align" : "center"
					});
					container.addChild(output_5, {
						"colIndex": 6,
						"rowIndex": 1,
						"colSpan": 1,
						"rowSpan": 1
					});
					var inputBox_3 = new cpr.controls.InputBox("ipb4");
					inputBox_3.inputFilter = "[\\d,\\.]";
					inputBox_3.bind("value").toDataMap(app.lookup("product_device"), "height");
					container.addChild(inputBox_3, {
						"colIndex": 3,
						"rowIndex": 1
					});
					var inputBox_4 = new cpr.controls.InputBox("ipb5");
					inputBox_4.inputFilter = "[\\d,\\.]";
					inputBox_4.bind("value").toDataMap(app.lookup("product_device"), "depth");
					container.addChild(inputBox_4, {
						"colIndex": 5,
						"rowIndex": 1
					});
					var inputBox_5 = new cpr.controls.InputBox("ipb3");
					inputBox_5.inputFilter = "[\\d,\\.]";
					inputBox_5.style.css({
						"background-color" : "#ffffff"
					});
					inputBox_5.bind("value").toDataMap(app.lookup("product_device"), "width");
					container.addChild(inputBox_5, {
						"colIndex": 1,
						"rowIndex": 1
					});
				})(group_4);
				container.addChild(group_4, {
					"top": "201px",
					"left": "0px",
					"width": "522px",
					"height": "54px"
				});
				var fileInput_1 = new cpr.controls.FileInput("product_image");
				fileInput_1.showClearButton = true;
				fileInput_1.placeholder = "제품 이미지 선택";
				fileInput_1.acceptFile = "image/*";
				container.addChild(fileInput_1, {
					"top": "1px",
					"left": "1px",
					"width": "227px",
					"height": "25px"
				});
				var group_5 = new cpr.controls.Container();
				// Layout
				var xYLayout_4 = new cpr.controls.layouts.XYLayout();
				group_5.setLayout(xYLayout_4);
				(function(container){
					var inputBox_6 = new cpr.controls.InputBox("ipb2");
					inputBox_6.bind("value").toDataMap(app.lookup("product"), "product_version");
					container.addChild(inputBox_6, {
						"top": "0px",
						"left": "68px",
						"width": "131px",
						"height": "25px"
					});
					var output_6 = new cpr.controls.Output();
					output_6.value = "버전";
					output_6.style.css({
						"vertical-align" : "middle",
						"text-align" : "center"
					});
					container.addChild(output_6, {
						"top": "0px",
						"left": "7px",
						"width": "62px",
						"height": "25px"
					});
				})(group_5);
				container.addChild(group_5, {
					"top": "1px",
					"left": "498px",
					"width": "209px",
					"height": "25px"
				});
				var output_7 = new cpr.controls.Output();
				output_7.value = "제품 사이즈의 경우 너비, 높이, 깊이 순으로 입력해 주세요.";
				output_7.style.css({
					"color" : "#dd4545",
					"padding-left" : "10px",
					"vertical-align" : "middle",
					"font-size" : "9pt",
					"text-align" : "left"
				});
				container.addChild(output_7, {
					"top": "201px",
					"left": "0px",
					"width": "522px",
					"height": "23px"
				});
				var group_6 = new cpr.controls.Container();
				// Layout
				var flowLayout_1 = new cpr.controls.layouts.FlowLayout();
				flowLayout_1.scrollable = false;
				flowLayout_1.lineWrap = true;
				group_6.setLayout(flowLayout_1);
				(function(container){
					var grid_2 = new cpr.controls.Grid("communication");
					grid_2.init({
						"columns": [
							{"width": "92px"},
							{"width": "128px"},
							{"width": "89px"}
						],
						"header": {
							"rows": [
								{"height": "24px"},
								{"height": "24px"}
							],
							"cells": [
								{
									"constraint": {"rowIndex": 1, "colIndex": 0},
									"configurator": function(cell){
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "Server";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 1, "colIndex": 1},
									"configurator": function(cell){
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "Wireless LAN(Wi-Fi)";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 1, "colIndex": 2},
									"configurator": function(cell){
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "Other";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 0, "rowSpan": 1, "colSpan": 3},
									"configurator": function(cell){
										cell.text = "통신 방식";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								}
							]
						},
						"detail": {
							"rows": [{"height": "24px"}],
							"cells": [
								{
									"constraint": {"rowIndex": 0, "colIndex": 0},
									"configurator": function(cell){
										cell.control = (function(){
											var inputBox_7 = new cpr.controls.InputBox("ipb6");
											inputBox_7.bind("value").toDataMap(app.lookup("product_device"), "server");
											return inputBox_7;
										})();
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 1},
									"configurator": function(cell){
										cell.style.css({
											"vertical-align" : "middle",
											"text-align" : "center"
										});
										cell.control = (function(){
											var comboBox_1 = new cpr.controls.ComboBox("cmb1");
											comboBox_1.tabIndex = 0;
											comboBox_1.preventInput = true;
											comboBox_1.bind("value").toDataMap(app.lookup("product_device"), "wi_fi");
											(function(comboBox_1){
												comboBox_1.addItem(new cpr.controls.Item("O", "O"));
												comboBox_1.addItem(new cpr.controls.Item("X", "X"));
											})(comboBox_1);
											return comboBox_1;
										})();
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 2},
									"configurator": function(cell){
										cell.control = (function(){
											var inputBox_8 = new cpr.controls.InputBox("ipb7");
											inputBox_8.bind("value").toDataMap(app.lookup("product_device"), "other");
											return inputBox_8;
										})();
									}
								}
							]
						}
					});
					container.addChild(grid_2, {
						"autoSize": "height",
						"width": "400px",
						"height": "74px"
					});
				})(group_6);
				container.addChild(group_6, {
					"top": "265px",
					"left": "0px",
					"width": "404px",
					"height": "73px"
				});
				var group_7 = new cpr.controls.Container();
				// Layout
				var verticalLayout_1 = new cpr.controls.layouts.VerticalLayout();
				group_7.setLayout(verticalLayout_1);
				(function(container){
					var output_8 = new cpr.controls.Output();
					output_8.value = "설명";
					output_8.style.css({
						"padding-left" : "10px"
					});
					container.addChild(output_8, {
						"width": "100px",
						"height": "25px"
					});
					var textArea_1 = new cpr.controls.TextArea("txa1");
					textArea_1.bind("value").toDataMap(app.lookup("product"), "explanation");
					container.addChild(textArea_1, {
						"width": "100px",
						"height": "139px"
					});
				})(group_7);
				container.addChild(group_7, {
					"top": "348px",
					"left": "0px",
					"width": "704px",
					"height": "172px"
				});
				var group_8 = new cpr.controls.Container();
				// Layout
				var xYLayout_5 = new cpr.controls.layouts.XYLayout();
				group_8.setLayout(xYLayout_5);
				(function(container){
					var output_9 = new cpr.controls.Output();
					output_9.value = "방수/방진";
					output_9.style.css({
						"vertical-align" : "middle",
						"text-align" : "center"
					});
					container.addChild(output_9, {
						"top": "41px",
						"left": "6px",
						"width": "96px",
						"height": "25px"
					});
					var inputBox_9 = new cpr.controls.InputBox("ipb8");
					inputBox_9.bind("value").toDataMap(app.lookup("product_device"), "ip_ratings");
					container.addChild(inputBox_9, {
						"top": "41px",
						"left": "101px",
						"width": "150px",
						"height": "25px"
					});
				})(group_8);
				container.addChild(group_8, {
					"top": "265px",
					"left": "414px",
					"width": "290px",
					"height": "73px"
				});
				var output_10 = new cpr.controls.Output();
				output_10.value = "방수/방진 등급을 작성해 주세요.\r\n해당되지 않을 경우 공란으로 유지";
				output_10.style.css({
					"color" : "#dd4545",
					"padding-left" : "10px",
					"vertical-align" : "middle",
					"font-size" : "9pt",
					"text-align" : "left"
				});
				container.addChild(output_10, {
					"top": "265px",
					"left": "429px",
					"width": "275px",
					"height": "38px"
				});
				var group_9 = new cpr.controls.Container();
				// Layout
				var formLayout_3 = new cpr.controls.layouts.FormLayout();
				formLayout_3.topMargin = "0px";
				formLayout_3.rightMargin = "0px";
				formLayout_3.bottomMargin = "0px";
				formLayout_3.leftMargin = "0px";
				formLayout_3.horizontalSpacing = "0px";
				formLayout_3.verticalSpacing = "8px";
				formLayout_3.setColumns(["16fr", "34px"]);
				formLayout_3.setRows(["25px", "1fr"]);
				group_9.setLayout(formLayout_3);
				(function(container){
					var output_11 = new cpr.controls.Output();
					output_11.value = "담당 개발자";
					output_11.style.css({
						"padding-left" : "10px"
					});
					container.addChild(output_11, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var button_1 = new cpr.controls.Button();
					button_1.value = "+";
					button_1.style.css({
						"background-color" : "#eaf0ea",
						"border-bottom-color" : "#c2c2c2",
						"border-left-color" : "#c2c2c2",
						"border-top-color" : "#c2c2c2",
						"border-right-color" : "#c2c2c2",
						"background-image" : "none"
					});
					if(typeof onButtonClick2 == "function") {
						button_1.addEventListener("click", onButtonClick2);
					}
					container.addChild(button_1, {
						"colIndex": 1,
						"rowIndex": 0,
						"colSpan": 1,
						"rowSpan": 1
					});
					var grid_3 = new cpr.controls.Grid("grid_developer");
					grid_3.init({
						"dataSet": app.lookup("developerList"),
						"columns": [
							{"width": "33px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"}
						],
						"header": {
							"rows": [
								{"height": "24px"},
								{"height": "24px"}
							],
							"cells": [
								{
									"constraint": {"rowIndex": 0, "colIndex": 0, "rowSpan": 2, "colSpan": 1},
									"configurator": function(cell){
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "번호";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 1, "rowSpan": 2, "colSpan": 1},
									"configurator": function(cell){
										cell.targetColumnName = "department";
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "부서명";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 2, "rowSpan": 2, "colSpan": 1},
									"configurator": function(cell){
										cell.targetColumnName = "employees_number";
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "사원번호";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 3, "rowSpan": 2, "colSpan": 1},
									"configurator": function(cell){
										cell.targetColumnName = "employees_name";
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "성명";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 1, "colIndex": 4},
									"configurator": function(cell){
										cell.targetColumnName = "start_date";
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "시작일";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 1, "colIndex": 5},
									"configurator": function(cell){
										cell.targetColumnName = "end_date";
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "종료일";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 4, "rowSpan": 1, "colSpan": 2},
									"configurator": function(cell){
										cell.text = "담당 기간";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								}
							]
						},
						"detail": {
							"rows": [{"height": "24px"}],
							"cells": [
								{
									"constraint": {"rowIndex": 0, "colIndex": 0},
									"configurator": function(cell){
										cell.columnType = "rowindex";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 1},
									"configurator": function(cell){
										cell.columnName = "department";
										cell.control = (function(){
											var inputBox_10 = new cpr.controls.InputBox("ipb9");
											inputBox_10.bind("value").toDataColumn("department");
											return inputBox_10;
										})();
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 2},
									"configurator": function(cell){
										cell.columnName = "employees_number";
										cell.control = (function(){
											var maskEditor_1 = new cpr.controls.MaskEditor("mse1");
											maskEditor_1.mask = "000000000";
											maskEditor_1.bind("value").toDataColumn("employees_number");
											return maskEditor_1;
										})();
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 3},
									"configurator": function(cell){
										cell.columnName = "employees_name";
										cell.control = (function(){
											var inputBox_11 = new cpr.controls.InputBox("ipb10");
											inputBox_11.bind("value").toDataColumn("employees_name");
											return inputBox_11;
										})();
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 4},
									"configurator": function(cell){
										cell.columnName = "start_date";
										cell.control = (function(){
											var dateInput_1 = new cpr.controls.DateInput("dti1");
											dateInput_1.format = "YYYY-MM-DD";
											dateInput_1.bind("value").toDataColumn("start_date");
											return dateInput_1;
										})();
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 5},
									"configurator": function(cell){
										cell.columnName = "end_date";
										cell.control = (function(){
											var dateInput_2 = new cpr.controls.DateInput("dti2");
											dateInput_2.format = "YYYY-MM-DD";
											dateInput_2.bind("value").toDataColumn("end_date");
											return dateInput_2;
										})();
									}
								}
							]
						}
					});
					container.addChild(grid_3, {
						"colIndex": 0,
						"rowIndex": 1,
						"colSpan": 2,
						"rowSpan": 1
					});
				})(group_9);
				container.addChild(group_9, {
					"top": "535px",
					"left": "0px",
					"width": "704px",
					"height": "225px"
				});
				var button_2 = new cpr.controls.Button();
				button_2.value = "등록";
				button_2.style.css({
					"background-color" : "#DAF2DA",
					"border-right-style" : "none",
					"border-radius" : "10px",
					"border-left-style" : "none",
					"border-bottom-style" : "none",
					"background-image" : "none",
					"border-top-style" : "none"
				});
				if(typeof onButtonClick == "function") {
					button_2.addEventListener("click", onButtonClick);
				}
				container.addChild(button_2, {
					"top": "775px",
					"left": "624px",
					"width": "80px",
					"height": "25px"
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "0px",
				"left": "0px",
				"width": "707px",
				"height": "815px"
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
		}
	});
	app.title = "addDevice";
	cpr.core.Platform.INSTANCE.register(app);
})();
