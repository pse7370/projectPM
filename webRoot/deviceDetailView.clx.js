/*
 * App URI: deviceDetailView
 * Source Location: deviceDetailView.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("deviceDetailView", {
		onPrepare: function(loader){
			loader.addCSS("theme/css/main.css");
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * deviceDetailView.js
			 * Created at 2022. 5. 29. 오전 3:27:55.
			 *
			 * @author PSE
			 ************************************************/;
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
				"rows": []
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
					{"name": "save_path"},
					{"name": "explanation"}
				]
			});
			app.register(dataMap_2);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 748px) and (max-width: 1023px)", "deatilView");
			app.supportMedia("all and (min-width: 748px) and (max-width: 747px)", "layout");
			app.supportMedia("all and (min-width: 500px) and (max-width: 747px)", "tablet");
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
			var image_1 = new cpr.controls.Image();
			(function(image_1){
			})(image_1);
			container.addChild(image_1, {
				"top": "20px",
				"left": "20px",
				"width": "180px",
				"height": "180px"
			});
			
			var group_1 = new cpr.controls.Container();
			// Layout
			var formLayout_1 = new cpr.controls.layouts.FormLayout();
			formLayout_1.topMargin = "0px";
			formLayout_1.rightMargin = "0px";
			formLayout_1.bottomMargin = "0px";
			formLayout_1.leftMargin = "0px";
			formLayout_1.horizontalSpacing = "0px";
			formLayout_1.verticalSpacing = "0px";
			formLayout_1.setColumns(["150px", "1fr"]);
			formLayout_1.setRows(["35px", "35px", "35px", "35px"]);
			group_1.setLayout(formLayout_1);
			(function(container){
				var output_1 = new cpr.controls.Output();
				output_1.value = "제품명";
				output_1.style.css({
					"border-right-style" : "solid",
					"background-color" : "#eaf0ea",
					"border-top-width" : "1px",
					"border-right-width" : "1px",
					"border-left-style" : "solid",
					"border-left-color" : "#b4b4b4",
					"border-top-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-left-width" : "1px",
					"border-top-style" : "solid",
					"text-align" : "center"
				});
				container.addChild(output_1, {
					"colIndex": 0,
					"rowIndex": 0
				});
				var output_2 = new cpr.controls.Output();
				output_2.value = "제품 사이즈";
				output_2.style.css({
					"border-right-style" : "solid",
					"background-color" : "#eaf0ea",
					"border-right-width" : "1px",
					"border-left-style" : "solid",
					"border-left-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-left-width" : "1px",
					"text-align" : "center"
				});
				container.addChild(output_2, {
					"colIndex": 0,
					"rowIndex": 2
				});
				var output_3 = new cpr.controls.Output();
				output_3.value = "버전";
				output_3.style.css({
					"border-right-style" : "solid",
					"border-top-width" : "1px",
					"border-bottom-color" : "#b4b4b4",
					"border-right-width" : "1px",
					"border-left-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-left-width" : "1px",
					"border-top-style" : "solid",
					"background-color" : "#eaf0ea",
					"border-left-style" : "solid",
					"border-bottom-width" : "1px",
					"border-top-color" : "#b4b4b4",
					"border-bottom-style" : "solid",
					"text-align" : "center"
				});
				container.addChild(output_3, {
					"colIndex": 0,
					"rowIndex": 1
				});
				var output_4 = new cpr.controls.Output();
				output_4.value = "방수/방진";
				output_4.style.css({
					"border-right-style" : "solid",
					"border-top-width" : "1px",
					"border-bottom-color" : "#b4b4b4",
					"border-right-width" : "1px",
					"border-left-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-left-width" : "1px",
					"border-top-style" : "solid",
					"background-color" : "#eaf0ea",
					"border-left-style" : "solid",
					"border-bottom-width" : "1px",
					"border-top-color" : "#b4b4b4",
					"border-bottom-style" : "solid",
					"text-align" : "center"
				});
				container.addChild(output_4, {
					"colIndex": 0,
					"rowIndex": 3
				});
				var output_5 = new cpr.controls.Output();
				output_5.style.css({
					"border-right-style" : "solid",
					"border-top-width" : "1px",
					"border-right-width" : "1px",
					"padding-left" : "10px",
					"border-top-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-top-style" : "solid"
				});
				output_5.bind("value").toDataMap(app.lookup("product"), "product_name");
				container.addChild(output_5, {
					"colIndex": 1,
					"rowIndex": 0
				});
				var output_6 = new cpr.controls.Output();
				output_6.style.css({
					"border-right-style" : "solid",
					"border-top-width" : "1px",
					"border-bottom-color" : "#b4b4b4",
					"border-right-width" : "1px",
					"padding-left" : "10px",
					"border-bottom-width" : "1px",
					"border-top-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-bottom-style" : "solid",
					"border-top-style" : "solid"
				});
				output_6.bind("value").toDataMap(app.lookup("product"), "product_version");
				container.addChild(output_6, {
					"colIndex": 1,
					"rowIndex": 1
				});
				var output_7 = new cpr.controls.Output();
				output_7.value = "Output";
				output_7.style.css({
					"border-right-style" : "solid",
					"border-right-width" : "1px",
					"padding-left" : "10px",
					"border-right-color" : "#b4b4b4"
				});
				container.addChild(output_7, {
					"colIndex": 1,
					"rowIndex": 2
				});
				var output_8 = new cpr.controls.Output();
				output_8.style.css({
					"border-right-style" : "solid",
					"border-top-width" : "1px",
					"border-bottom-color" : "#b4b4b4",
					"border-right-width" : "1px",
					"padding-left" : "10px",
					"border-bottom-width" : "1px",
					"border-top-color" : "#b4b4b4",
					"border-right-color" : "#b4b4b4",
					"border-bottom-style" : "solid",
					"border-top-style" : "solid"
				});
				output_8.bind("value").toDataMap(app.lookup("product_device"), "ip_ratings");
				container.addChild(output_8, {
					"colIndex": 1,
					"rowIndex": 3
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "35px",
				"left": "222px",
				"width": "494px",
				"height": "146px"
			});
			
			var group_2 = new cpr.controls.Container();
			// Layout
			var verticalLayout_1 = new cpr.controls.layouts.VerticalLayout();
			group_2.setLayout(verticalLayout_1);
			(function(container){
				var grid_1 = new cpr.controls.Grid("authentication");
				grid_1.readOnly = true;
				grid_1.init({
					"dataSet": app.lookup("authenticationList"),
					"autoRowHeight": "all",
					"columns": [
						{"width": "83px"},
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
									cell.targetColumnName = "auth_type";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "인증 방식";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 1},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_one_max_user";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : 1";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 2},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_many_max_user";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : N";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 3},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_one_max_template";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : 1";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 4},
								"configurator": function(cell){
									cell.targetColumnName = "one_to_many_max_template";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "1 : N";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1, "rowSpan": 1, "colSpan": 2},
								"configurator": function(cell){
									cell.text = "최대 등록 가능 사용자 수";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3, "rowSpan": 1, "colSpan": 2},
								"configurator": function(cell){
									cell.text = "최대 등록 가능 템플릿 수";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
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
									cell.columnName = "auth_type";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.columnName = "one_to_one_max_user";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.columnName = "one_to_many_max_user";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.columnName = "one_to_one_max_template";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.columnName = "one_to_many_max_template";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							}
						]
					}
				});
				container.addChild(grid_1, {
					"autoSize": "height",
					"width": "400px",
					"height": "177px"
				});
			})(group_2);
			container.addChild(group_2, {
				"top": "218px",
				"left": "20px",
				"width": "696px",
				"height": "187px"
			});
			
			var group_3 = new cpr.controls.Container();
			// Layout
			var verticalLayout_2 = new cpr.controls.layouts.VerticalLayout();
			group_3.setLayout(verticalLayout_2);
			(function(container){
				var grid_2 = new cpr.controls.Grid("communication");
				grid_2.init({
					"columns": [
						{"width": "100px"},
						{"width": "129px"},
						{"width": "109px"}
					],
					"header": {
						"rows": [
							{"height": "27px"},
							{"height": "27px"}
						],
						"cells": [
							{
								"constraint": {"rowIndex": 1, "colIndex": 0},
								"configurator": function(cell){
									cell.text = "Server";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 1},
								"configurator": function(cell){
									cell.text = "Wireless LAN(Wi-Fi)";
									cell.style.css({
										"background-color" : "#eaf0ea"
									});
								}
							},
							{
								"constraint": {"rowIndex": 1, "colIndex": 2},
								"configurator": function(cell){
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
						"rows": [{"height": "27px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.columnName = "server";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.columnName = "wi_fi";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.columnName = "other";
								}
							}
						]
					}
				});
				container.addChild(grid_2, {
					"autoSize": "height",
					"width": "591px",
					"height": "94px"
				});
			})(group_3);
			container.addChild(group_3, {
				"top": "418px",
				"left": "20px",
				"width": "586px",
				"height": "102px"
			});
			
			var group_4 = new cpr.controls.Container();
			// Layout
			var verticalLayout_3 = new cpr.controls.layouts.VerticalLayout();
			group_4.setLayout(verticalLayout_3);
			(function(container){
				var output_9 = new cpr.controls.Output();
				output_9.value = "설명";
				output_9.style.css({
					"padding-left" : "10px"
				});
				container.addChild(output_9, {
					"width": "100px",
					"height": "25px"
				});
				var textArea_1 = new cpr.controls.TextArea("txa1");
				textArea_1.readOnly = true;
				textArea_1.bind("value").toDataMap(app.lookup("product"), "explanation");
				container.addChild(textArea_1, {
					"autoSize": "none",
					"width": "100px",
					"height": "143px"
				});
			})(group_4);
			container.addChild(group_4, {
				"top": "530px",
				"left": "20px",
				"width": "696px",
				"height": "181px"
			});
			
			var group_5 = new cpr.controls.Container();
			// Layout
			var verticalLayout_4 = new cpr.controls.layouts.VerticalLayout();
			group_5.setLayout(verticalLayout_4);
			(function(container){
				var output_10 = new cpr.controls.Output();
				output_10.value = "담당 개발자";
				output_10.style.css({
					"padding-left" : "10px"
				});
				container.addChild(output_10, {
					"width": "100px",
					"height": "25px"
				});
				var grid_3 = new cpr.controls.Grid("grid_developer");
				grid_3.readOnly = true;
				grid_3.init({
					"dataSet": app.lookup("developerList"),
					"columns": [
						{"width": "36px"},
						{"width": "109px"},
						{"width": "98px"},
						{"width": "100px"},
						{"width": "100px"},
						{"width": "100px"}
					],
					"header": {
						"rows": [{"height": "25px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "번호";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.targetColumnName = "department";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "부서명";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.targetColumnName = "employees_number";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "사원 번호";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.targetColumnName = "employees_name";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "성명";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.targetColumnName = "start_date";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "시작일";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 5},
								"configurator": function(cell){
									cell.targetColumnName = "end_date";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "종료일";
									cell.style.css({
										"background-color" : "#eaf0ea",
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							}
						]
					},
					"detail": {
						"rows": [{"height": "25px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.columnType = "rowindex";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.columnName = "department";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.columnName = "employees_number";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.columnName = "employees_name";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.columnName = "start_date";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 5},
								"configurator": function(cell){
									cell.columnName = "end_date";
									cell.style.css({
										"vertical-align" : "middle",
										"text-align" : "center"
									});
								}
							}
						]
					}
				});
				container.addChild(grid_3, {
					"width": "400px",
					"height": "164px"
				});
			})(group_5);
			container.addChild(group_5, {
				"top": "721px",
				"left": "20px",
				"width": "696px",
				"height": "200px"
			});
			
			var group_6 = new cpr.controls.Container();
			// Layout
			var xYLayout_2 = new cpr.controls.layouts.XYLayout();
			group_6.setLayout(xYLayout_2);
			(function(container){
			})(group_6);
			container.addChild(group_6, {
				"top": "0px",
				"left": "5px",
				"width": "730px",
				"height": "970px"
			});
			
			var button_1 = new cpr.controls.Button();
			button_1.value = "수정";
			button_1.style.css({
				"background-color" : "#DAF2DA",
				"border-right-style" : "none",
				"border-radius" : "10px",
				"border-left-style" : "none",
				"border-bottom-style" : "none",
				"background-image" : "none",
				"border-top-style" : "none"
			});
			container.addChild(button_1, {
				"top": "931px",
				"left": "539px",
				"width": "80px",
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
				"top": "931px",
				"left": "636px",
				"width": "80px",
				"height": "25px"
			});
		}
	});
	app.title = "deviceDetailView";
	cpr.core.Platform.INSTANCE.register(app);
})();
