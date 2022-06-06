/*
 * App URI: output/modifyOutput
 * Source Location: output/modifyOutput.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("output/modifyOutput", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * modifyOutput.js
			 * Created at 2022. 6. 6. 오후 2:48:59.
			 *
			 * @author PSE
			 ************************************************/;
			// End - User Script
			
			// Header
			var dataMap_1 = new cpr.data.DataMap("product");
			dataMap_1.parseData({
				"columns" : [{"name": "product_name"}]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("product_output");
			dataMap_2.parseData({
				"columns" : [
					{
						"name": "product_id",
						"dataType": "number"
					},
					{
						"name": "output_type",
						"dataType": "string"
					},
					{"name": "output_title"},
					{"name": "output_content"},
					{
						"name": "write_date",
						"dataType": "string"
					}
				]
			});
			app.register(dataMap_2);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 745px) and (max-width: 1023px)", "new-screen");
			app.supportMedia("all and (min-width: 500px) and (max-width: 744px)", "tablet");
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
				var group_2 = new cpr.controls.Container();
				// Layout
				var formLayout_1 = new cpr.controls.layouts.FormLayout();
				formLayout_1.topMargin = "0px";
				formLayout_1.rightMargin = "0px";
				formLayout_1.bottomMargin = "0px";
				formLayout_1.leftMargin = "0px";
				formLayout_1.horizontalSpacing = "0px";
				formLayout_1.verticalSpacing = "0px";
				formLayout_1.setColumns(["110px", "240px"]);
				formLayout_1.setRows(["30px"]);
				group_2.setLayout(formLayout_1);
				(function(container){
					var output_1 = new cpr.controls.Output();
					output_1.value = "산출물 종류";
					output_1.style.css({
						"text-align" : "center"
					});
					container.addChild(output_1, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var inputBox_1 = new cpr.controls.InputBox("ipb1");
					inputBox_1.bind("value").toDataMap(app.lookup("product_output"), "output_type");
					container.addChild(inputBox_1, {
						"colIndex": 1,
						"rowIndex": 0,
						"colSpan": 1,
						"rowSpan": 1
					});
				})(group_2);
				container.addChild(group_2, {
					"top": "30px",
					"left": "358px",
					"width": "357px",
					"height": "38px"
				});
				var fileUpload_1 = new cpr.controls.FileUpload("file_upload");
				fileUpload_1.buttons = ["add", "remove"];
				fileUpload_1.maxFileCount = 5;
				fileUpload_1.style.button.css({
					"background-color" : "#daf2da",
					"border-right-style" : "none",
					"border-left-style" : "none",
					"font-size" : "10pt",
					"border-bottom-style" : "none",
					"background-image" : "none",
					"border-top-style" : "none"
				});
				fileUpload_1.style.header.css({
					"background-color" : "#eaf0ea",
					"background-image" : "none"
				});
				container.addChild(fileUpload_1, {
					"top": "140px",
					"left": "20px",
					"width": "695px",
					"height": "202px"
				});
				var group_3 = new cpr.controls.Container();
				// Layout
				var verticalLayout_1 = new cpr.controls.layouts.VerticalLayout();
				group_3.setLayout(verticalLayout_1);
				(function(container){
					var output_2 = new cpr.controls.Output();
					output_2.value = "내용";
					output_2.style.css({
						"padding-left" : "10px"
					});
					container.addChild(output_2, {
						"width": "100px",
						"height": "27px"
					});
					var textArea_1 = new cpr.controls.TextArea("txa1");
					textArea_1.bind("value").toDataMap(app.lookup("product_output"), "output_content");
					container.addChild(textArea_1, {
						"width": "100px",
						"height": "252px"
					});
				})(group_3);
				container.addChild(group_3, {
					"top": "355px",
					"left": "20px",
					"width": "695px",
					"height": "290px"
				});
				var button_1 = new cpr.controls.Button();
				button_1.value = "수정";
				button_1.style.css({
					"border-right-style" : "none",
					"background-color" : "#DAF2DA",
					"border-radius" : "10px",
					"border-left-style" : "none",
					"border-bottom-style" : "none",
					"background-image" : "none",
					"border-top-style" : "none"
				});
				container.addChild(button_1, {
					"top": "655px",
					"left": "635px",
					"width": "80px",
					"height": "25px"
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "0px",
				"left": "0px",
				"width": "745px",
				"height": "700px"
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
			formLayout_2.setColumns(["100px", "190px"]);
			formLayout_2.setRows(["30px"]);
			group_4.setLayout(formLayout_2);
			(function(container){
				var output_3 = new cpr.controls.Output();
				output_3.value = "제품명";
				output_3.style.css({
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
				container.addChild(output_3, {
					"colIndex": 0,
					"rowIndex": 0
				});
				var output_4 = new cpr.controls.Output("productName");
				output_4.style.css({
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
				output_4.bind("value").toDataMap(app.lookup("product"), "product_name");
				container.addChild(output_4, {
					"colIndex": 1,
					"rowIndex": 0
				});
			})(group_4);
			container.addChild(group_4, {
				"top": "30px",
				"left": "20px",
				"width": "297px",
				"height": "38px"
			});
			
			var group_5 = new cpr.controls.Container();
			// Layout
			var formLayout_3 = new cpr.controls.layouts.FormLayout();
			formLayout_3.topMargin = "0px";
			formLayout_3.rightMargin = "0px";
			formLayout_3.bottomMargin = "0px";
			formLayout_3.leftMargin = "0px";
			formLayout_3.horizontalSpacing = "0px";
			formLayout_3.verticalSpacing = "0px";
			formLayout_3.setColumns(["60px", "630px"]);
			formLayout_3.setRows(["30px"]);
			group_5.setLayout(formLayout_3);
			(function(container){
				var output_5 = new cpr.controls.Output();
				output_5.value = "제목";
				output_5.style.css({
					"padding-left" : "13px",
					"text-align" : "left"
				});
				container.addChild(output_5, {
					"colIndex": 0,
					"rowIndex": 0
				});
				var inputBox_2 = new cpr.controls.InputBox("ipb2");
				inputBox_2.bind("value").toDataMap(app.lookup("product_output"), "output_title");
				container.addChild(inputBox_2, {
					"colIndex": 1,
					"rowIndex": 0,
					"colSpan": 1,
					"rowSpan": 1
				});
			})(group_5);
			container.addChild(group_5, {
				"top": "84px",
				"left": "20px",
				"width": "695px",
				"height": "38px"
			});
		}
	});
	app.title = "modifyOutput";
	cpr.core.Platform.INSTANCE.register(app);
})();