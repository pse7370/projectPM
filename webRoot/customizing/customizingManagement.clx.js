/*
 * App URI: customizing/customizingManagement
 * Source Location: customizing/customizingManagement.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("customizing/customizingManagement", {
		onPrepare: function(loader){
			loader.addCSS("theme/css/addProduct_style.css");
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * customizingMangement.js
			 * Created at 2022. 6. 3. 오전 12:03:57.
			 *
			 * @author PSE
			 ************************************************/
			
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				var product_id = app.getRootAppInstance().getAppProperty("product_id"); // 부모화면 데이터 셋
				console.log("product_id : " + product_id);
				
				var dataProduct_id = app.lookup("product_id");
				dataProduct_id.setValue("product_id", Number(product_id));
				console.log(dataProduct_id.getValue("product_id"));
				
				app.lookup("getCustomizingList").send();
				console.log("getCustomizingList 서브미션 실행");
				
			}
			
			/*
			 * getCustomizingList 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onGetCustomizingListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var getCustomizingList = e.control;
				
				app.lookup("productName").redraw();
				app.lookup("grid_customizing").redraw();
				
			}
			
			
			
			/*
			 * "추가/수정" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				
				var embeddedApp = app.getHost();
				
				cpr.core.App.load("customizing/modifyCustomizing", function(loadedApp){
					if(loadedApp){
						embeddedApp.initValue = {
							
						}
			    		embeddedApp.app = loadedApp;	    		
			  		}
				});
				
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("product_customizing");
			dataSet_1.parseData({
				"columns": [
					{"name": "customizing_id"},
					{"name": "customizing_version"},
					{"name": "customized_function"},
					{"name": "department"},
					{
						"name": "employees_number",
						"dataType": "number"
					},
					{"name": "employees_name"},
					{
						"name": "start_dates",
						"dataType": "string"
					},
					{"name": "end_date"}
				],
				"rows": []
			});
			app.register(dataSet_1);
			var dataMap_1 = new cpr.data.DataMap("product");
			dataMap_1.parseData({
				"columns" : [
					{
						"name": "product_id",
						"dataType": "number"
					},
					{"name": "product_name"}
				]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("product_id");
			dataMap_2.parseData({
				"columns" : [{
					"name": "product_id",
					"dataType": "number"
				}]
			});
			app.register(dataMap_2);
			var submission_1 = new cpr.protocols.Submission("getCustomizingList");
			submission_1.action = "/productMangement/getCustomizingList";
			submission_1.addRequestData(dataMap_2);
			submission_1.addResponseData(dataMap_1, false);
			submission_1.addResponseData(dataSet_1, false);
			if(typeof onGetCustomizingListSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onGetCustomizingListSubmitDone);
			}
			app.register(submission_1);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 770px) and (max-width: 1023px)", "new-screen");
			app.supportMedia("all and (min-width: 740px) and (max-width: 769px)", "dialog");
			app.supportMedia("all and (min-width: 500px) and (max-width: 739px)", "tablet");
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
			var group_1 = new cpr.controls.Container();
			// Layout
			var xYLayout_2 = new cpr.controls.layouts.XYLayout();
			group_1.setLayout(xYLayout_2);
			(function(container){
				var output_1 = new cpr.controls.Output();
				output_1.value = "커스터마이징 이력";
				output_1.style.css({
					"font-weight" : "bold",
					"vertical-align" : "middle",
					"font-size" : "14pt",
					"text-align" : "left"
				});
				container.addChild(output_1, {
					"top": "30px",
					"left": "20px",
					"width": "201px",
					"height": "41px"
				});
				var group_2 = new cpr.controls.Container();
				// Layout
				var formLayout_1 = new cpr.controls.layouts.FormLayout();
				formLayout_1.topMargin = "0px";
				formLayout_1.rightMargin = "0px";
				formLayout_1.bottomMargin = "0px";
				formLayout_1.leftMargin = "0px";
				formLayout_1.horizontalSpacing = "0px";
				formLayout_1.verticalSpacing = "0px";
				formLayout_1.setColumns(["100px", "180px"]);
				formLayout_1.setRows(["30px"]);
				group_2.setLayout(formLayout_1);
				(function(container){
					var output_2 = new cpr.controls.Output();
					output_2.value = "제품명";
					output_2.style.css({
						"border-right-style" : "solid",
						"border-top-width" : "1px",
						"border-bottom-color" : "#b4b4b4",
						"border-right-width" : "1px",
						"border-left-color" : "#b4b4b4",
						"border-right-color" : "#b4b4b4",
						"border-left-width" : "1px",
						"border-top-style" : "solid",
						"border-left-style" : "solid",
						"border-bottom-width" : "1px",
						"border-top-color" : "#b4b4b4",
						"border-bottom-style" : "solid",
						"text-align" : "center"
					});
					container.addChild(output_2, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var output_3 = new cpr.controls.Output("productName");
					output_3.style.css({
						"border-right-style" : "solid",
						"border-top-width" : "1px",
						"border-bottom-color" : "#b4b4b4",
						"border-right-width" : "1px",
						"border-bottom-width" : "1px",
						"border-top-color" : "#b4b4b4",
						"border-bottom-style" : "solid",
						"border-right-color" : "#b4b4b4",
						"border-top-style" : "solid",
						"text-align" : "center"
					});
					output_3.bind("value").toDataMap(app.lookup("product"), "product_name");
					container.addChild(output_3, {
						"colIndex": 1,
						"rowIndex": 0
					});
				})(group_2);
				container.addChild(group_2, {
					"top": "93px",
					"left": "20px",
					"width": "284px",
					"height": "38px"
				});
				var group_3 = new cpr.controls.Container();
				// Layout
				var verticalLayout_1 = new cpr.controls.layouts.VerticalLayout();
				group_3.setLayout(verticalLayout_1);
				(function(container){
					var grid_1 = new cpr.controls.Grid("grid_customizing");
					grid_1.readOnly = true;
					grid_1.init({
						"dataSet": app.lookup("product_customizing"),
						"autoRowHeight": "1, 2, 3, 4, 5, 6",
						"collapsible": true,
						"columns": [
							{"width": "25px"},
							{"width": "31px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"},
							{"width": "100px"}
						],
						"header": {
							"rows": [
								{"height": "27px"},
								{"height": "27px"}
							],
							"cells": [
								{
									"constraint": {"rowIndex": 0, "colIndex": 0, "rowSpan": 2, "colSpan": 1},
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
									"constraint": {"rowIndex": 0, "colIndex": 1, "rowSpan": 2, "colSpan": 1},
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
									"constraint": {"rowIndex": 0, "colIndex": 2, "rowSpan": 2, "colSpan": 1},
									"configurator": function(cell){
										cell.targetColumnName = "customized_function";
										cell.filterable = false;
										cell.sortable = false;
										cell.text = "기능";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 1, "colIndex": 3},
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
									"constraint": {"rowIndex": 1, "colIndex": 4},
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
									"constraint": {"rowIndex": 1, "colIndex": 5},
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
									"constraint": {"rowIndex": 1, "colIndex": 6},
									"configurator": function(cell){
										cell.targetColumnName = "start_dates";
										cell.filterable = false;
										cell.sortable = true;
										cell.text = "시작일";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 1, "colIndex": 7},
									"configurator": function(cell){
										cell.targetColumnName = "end_date";
										cell.filterable = false;
										cell.sortable = true;
										cell.text = "종료일";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 3, "rowSpan": 1, "colSpan": 5},
									"configurator": function(cell){
										cell.text = "담당 개발자";
										cell.style.css({
											"background-color" : "#eaf0ea"
										});
									}
								}
							]
						},
						"detail": {
							"rows": [{"height": "27px"}],
							"cells": [
								{
									"constraint": {"rowIndex": 0, "colIndex": 0},
									"configurator": function(cell){
										cell.columnType = "checkbox";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 1},
									"configurator": function(cell){
										cell.columnType = "rowindex";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 2},
									"configurator": function(cell){
										cell.columnName = "customized_function";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 3},
									"configurator": function(cell){
										cell.columnName = "department";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 4},
									"configurator": function(cell){
										cell.columnName = "employees_number";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 5},
									"configurator": function(cell){
										cell.columnName = "employees_name";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 6},
									"configurator": function(cell){
										cell.columnName = "start_dates";
									}
								},
								{
									"constraint": {"rowIndex": 0, "colIndex": 7},
									"configurator": function(cell){
										cell.columnName = "end_date";
									}
								}
							]
						},
						"rowGroup": [{
							"groupCondition": "customizing_version",
							"gheader": {
								"rows": [{"height": "30px"}],
								"cells": [{
									"constraint": {"rowIndex": 0, "colIndex": 0, "rowSpan": 1, "colSpan": 8},
									"configurator": function(cell){
										cell.expr = "\"[버전]     \" + customizing_version";
										cell.style.css({
											"font-weight" : "bold",
											"padding-left" : "15px",
											"text-align" : "left"
										});
									}
								}]
							}
						}]
					});
					container.addChild(grid_1, {
						"autoSize": "both",
						"width": "739px",
						"height": "475px"
					});
				})(group_3);
				container.addChild(group_3, {
					"top": "141px",
					"left": "20px",
					"width": "735px",
					"height": "481px"
				});
				var button_1 = new cpr.controls.Button();
				button_1.value = "추가/수정";
				button_1.style.css({
					"background-color" : "#DAF2DA",
					"border-right-style" : "none",
					"border-radius" : "10px",
					"border-left-style" : "none",
					"border-bottom-style" : "none",
					"background-image" : "none",
					"border-top-style" : "none"
				});
				if(typeof onButtonClick == "function") {
					button_1.addEventListener("click", onButtonClick);
				}
				container.addChild(button_1, {
					"top": "640px",
					"left": "526px",
					"width": "111px",
					"height": "25px"
				});
				var button_2 = new cpr.controls.Button();
				button_2.value = "삭제";
				button_2.style.css({
					"background-color" : "#DAF2DA",
					"border-right-style" : "none",
					"border-radius" : "10px",
					"border-left-style" : "none",
					"border-bottom-style" : "none",
					"background-image" : "none",
					"border-top-style" : "none"
				});
				container.addChild(button_2, {
					"top": "640px",
					"left": "651px",
					"width": "80px",
					"height": "25px"
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "0px",
				"left": "0px",
				"width": "760px",
				"height": "680px"
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
			if(typeof onBodyInit == "function"){
				app.addEventListener("init", onBodyInit);
			}
		}
	});
	app.title = "customizingManagement";
	cpr.core.Platform.INSTANCE.register(app);
})();
