/*
 * App URI: output/outputContentView
 * Source Location: output/outputContentView.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("output/outputContentView", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * outputContentView.js
			 * Created at 2022. 6. 6. 오후 1:45:09.
			 *
			 * @author PSE
			 ************************************************/
			
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				
				
				app.lookup("output_id").setValue("output_id", app.getHost().initValue.output_id); 
				console.log("output_id : " + app.lookup("output_id").getValue("output_id"));
				
				app.lookup("getOutputContent").send();
				console.log("getOutputContent 서브미션 실행");
				
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onGetOutputContentSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var getOutputContent = e.control;
				
				app.lookup("productName").redraw();
				app.lookup("outputType").redraw();
				app.lookup("writeDate").redraw();
				app.lookup("outputTitle").redraw();
				app.lookup("outputContent").redraw();
				
				var attachmentList = app.lookup("attachmentList");
				var i
				for(i = 0; i < attachmentList.getRowCount(); i++) {
					var fileName = attachmentList.getValue(i, "real_file_name");
					var fileSize = attachmentList.getValue(i, "file_size");
					var save_path = attachmentList.getValue(i, "save_path")
					app.lookup("file_upload").addUploadedFile(
						{
							name : fileName, 
							size : fileSize, 
							properties : {rowIndex : i}
						}
					);
				}
				
				
			}
			
			
			/*
			 * "수정" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				
				var embeddedApp = app.getHost();
				
				cpr.core.App.load("output/modifyOutput", function(loadedApp){
					if(loadedApp){
						embeddedApp.initValue = {
							output_id : app.lookup("output_id").getValue("output_id")
						}
			    		embeddedApp.app = loadedApp;	    		
			  		}
				});
				
			}
			
			
			/*
			 * (다운로드)파일 업로드에서 sendbutton-click 이벤트 발생 시 호출.
			 * 파일을 전송하는 button을 클릭 시 발생하는 이벤트. 서브미션을 통해 전송 버튼에 대한 구현이 필요합니다.
			 */
			function onFile_uploadSendbuttonClick(/* cpr.events.CEvent */ e){
				/** 
				 * @type cpr.controls.FileUpload
				 */
				var file_upload = e.control;
				console.log("download");
			
				var checkedFiles = file_upload.getSelection();
				var downloadFileList = app.lookup("downloadFileList");
				var attachmentList = app.lookup("attachmentList");
				var file = app.lookup("file");
				
				var i;
				var fileIndex;
				var fileName;
				var savePath;
				
				for(i = 0; i < checkedFiles.length; i++){
					fileIndex = file_upload.getIndex(checkedFiles[i]);
					fileName = attachmentList.getValue(fileIndex, "real_file_name");
					savePath = attachmentList.getValue(fileIndex, "save_path");
					
					downloadFileList.addRowData(
						{
							"file_name" : fileName,
							"save_path" : savePath
						}
					);
					
				} // end for
				
				 app.lookup("downloadAttachmentList").send();
				 
				 /*
				for(i = 0; i < checkedFiles.length; i++){
					fileIndex = file_upload.getIndex(checkedFiles[i]);
					fileName = attachmentList.getValue(fileIndex, "real_file_name");
					savePath = attachmentList.getValue(fileIndex, "save_path");
					
					file.setValue("file_name", fileName);
					file.setValue("save_path", savePath);
					
					app.lookup("downloadAttachment").send();
				}
				*/
				
			}
			
			/*
			 * 파일 업로드에서 download-click 이벤트 발생 시 호출.
			 * 파일을 다운받는 button을 클릭 시 발생하는 이벤트. 서브미션을 통해 다운로드 버튼에 대한 구현이 필요합니다.
			 */
			function onFile_uploadDownloadClick(/* cpr.events.CUploadedFileEvent */ e){
				/** 
				 * @type cpr.controls.FileUpload
				 */
				var file_upload = e.control;
				
				var clickFile = e.uploadedFile;
				var clickFileName = clickFile.name;
				console.log("clickFileName : " + clickFileName);
				app.lookup("file").setValue("file_name", clickFileName);
				
				var savePath = app.lookup("attachmentList").getValue(clickFile.getProperty("rowIndex"), "save_path");
				
				app.lookup("file").setValue("save_path", savePath);
				/*
				var savePath = clickFile.getProperty("savePath");	
				app.lookup("file").setValue("save_path", savePath);
				*/
				
				//app.lookup("downloadAttachment").action = "/productMangement/downloadAttachment?" + clickFileName;
				app.lookup("downloadAttachment").send();
				
			}
			
			
			
			/*
			 * "삭제" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick2(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				var output_id = app.lookup("output_id").getValue("output_id");
				
				app.lookup("deleteOutput").action = "/productMangement/deleteOutput?" + output_id;
			
				app.lookup("deleteOutput").send();
				console.log("deleteOutput 서브미션 실행");
				
			}
			
			/*
			 * deleteOutput 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onDeleteOutputSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var deleteOutput = e.control;
				
				app.getRootAppInstance().dialogManager.getDialogByName("outputContentView").close(1);
				
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("attachmentList");
			dataSet_1.parseData({
				"columns" : [
					{"name": "real_file_name"},
					{"name": "save_file_name"},
					{"name": "save_path"},
					{
						"name": "file_size",
						"dataType": "number"
					}
				]
			});
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("downloadFileList");
			dataSet_2.parseData({
				"columns" : [
					{"name": "file_name"},
					{"name": "save_path"}
				]
			});
			app.register(dataSet_2);
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
			
			var dataMap_3 = new cpr.data.DataMap("product_id");
			dataMap_3.parseData({
				"columns" : [{
					"name": "product_id",
					"dataType": "number"
				}]
			});
			app.register(dataMap_3);
			
			var dataMap_4 = new cpr.data.DataMap("output_id");
			dataMap_4.parseData({
				"columns" : [{
					"name": "output_id",
					"dataType": "number"
				}]
			});
			app.register(dataMap_4);
			
			var dataMap_5 = new cpr.data.DataMap("result");
			dataMap_5.parseData({
				"columns" : [{
					"name": "resultCode",
					"dataType": "number"
				}]
			});
			app.register(dataMap_5);
			
			var dataMap_6 = new cpr.data.DataMap("file");
			dataMap_6.parseData({
				"columns" : [
					{"name": "file_name"},
					{"name": "save_path"}
				]
			});
			app.register(dataMap_6);
			var submission_1 = new cpr.protocols.Submission("getOutputContent");
			submission_1.action = "/productMangement/getOutputContent";
			submission_1.addRequestData(dataMap_4);
			submission_1.addResponseData(dataMap_1, false);
			submission_1.addResponseData(dataMap_2, false);
			submission_1.addResponseData(dataSet_1, false);
			if(typeof onGetOutputContentSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onGetOutputContentSubmitDone);
			}
			app.register(submission_1);
			
			var submission_2 = new cpr.protocols.Submission("deleteOutput");
			submission_2.method = "delete";
			submission_2.action = "/productMangement/deleteOutput";
			submission_2.addResponseData(dataMap_5, false);
			if(typeof onDeleteOutputSubmitDone == "function") {
				submission_2.addEventListener("submit-done", onDeleteOutputSubmitDone);
			}
			app.register(submission_2);
			
			var submission_3 = new cpr.protocols.Submission("downloadAttachment");
			submission_3.action = "/productMangement/downloadAttachment";
			submission_3.responseType = "filedownload";
			submission_3.addRequestData(dataMap_6);
			if(typeof onDownloadAttachmentSubmitDone == "function") {
				submission_3.addEventListener("submit-done", onDownloadAttachmentSubmitDone);
			}
			app.register(submission_3);
			
			var submission_4 = new cpr.protocols.Submission("downloadAttachmentList");
			submission_4.action = "/productMangement/downloadAttachmentList";
			submission_4.responseType = "filedownload";
			submission_4.addRequestData(dataSet_2);
			app.register(submission_4);
			
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
				formLayout_1.setColumns(["100px", "190px"]);
				formLayout_1.setRows(["30px"]);
				group_2.setLayout(formLayout_1);
				(function(container){
					var output_1 = new cpr.controls.Output();
					output_1.value = "제품명";
					output_1.style.css({
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
					container.addChild(output_1, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var output_2 = new cpr.controls.Output("productName");
					output_2.style.css({
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
					output_2.bind("value").toDataMap(app.lookup("product"), "product_name");
					container.addChild(output_2, {
						"colIndex": 1,
						"rowIndex": 0
					});
				})(group_2);
				container.addChild(group_2, {
					"top": "30px",
					"left": "20px",
					"width": "297px",
					"height": "38px"
				});
				var group_3 = new cpr.controls.Container();
				// Layout
				var formLayout_2 = new cpr.controls.layouts.FormLayout();
				formLayout_2.topMargin = "0px";
				formLayout_2.rightMargin = "0px";
				formLayout_2.bottomMargin = "0px";
				formLayout_2.leftMargin = "0px";
				formLayout_2.horizontalSpacing = "0px";
				formLayout_2.verticalSpacing = "0px";
				formLayout_2.setColumns(["120px", "245px"]);
				formLayout_2.setRows(["30px"]);
				group_3.setLayout(formLayout_2);
				(function(container){
					var output_3 = new cpr.controls.Output();
					output_3.value = "산출물 종류";
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
					var output_4 = new cpr.controls.Output("outputType");
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
					output_4.bind("value").toDataMap(app.lookup("product_output"), "output_type");
					container.addChild(output_4, {
						"colIndex": 1,
						"rowIndex": 0,
						"colSpan": 1,
						"rowSpan": 1
					});
				})(group_3);
				container.addChild(group_3, {
					"top": "78px",
					"left": "20px",
					"width": "367px",
					"height": "38px"
				});
				var group_4 = new cpr.controls.Container();
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
				group_4.setLayout(formLayout_3);
				(function(container){
					var output_5 = new cpr.controls.Output();
					output_5.value = "제목";
					output_5.style.css({
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
					container.addChild(output_5, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var output_6 = new cpr.controls.Output("outputTitle");
					output_6.style.css({
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
					output_6.bind("value").toDataMap(app.lookup("product_output"), "output_title");
					container.addChild(output_6, {
						"colIndex": 1,
						"rowIndex": 0,
						"colSpan": 1,
						"rowSpan": 1
					});
				})(group_4);
				container.addChild(group_4, {
					"top": "126px",
					"left": "20px",
					"width": "693px",
					"height": "38px"
				});
				var group_5 = new cpr.controls.Container();
				// Layout
				var formLayout_4 = new cpr.controls.layouts.FormLayout();
				formLayout_4.topMargin = "0px";
				formLayout_4.rightMargin = "0px";
				formLayout_4.bottomMargin = "0px";
				formLayout_4.leftMargin = "0px";
				formLayout_4.horizontalSpacing = "0px";
				formLayout_4.verticalSpacing = "0px";
				formLayout_4.setColumns(["100px", "195px"]);
				formLayout_4.setRows(["30px"]);
				group_5.setLayout(formLayout_4);
				(function(container){
					var output_7 = new cpr.controls.Output();
					output_7.value = "작성일";
					output_7.style.css({
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
					container.addChild(output_7, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var maskEditor_1 = new cpr.controls.MaskEditor("writeDate");
					maskEditor_1.readOnly = true;
					maskEditor_1.mask = "0000-00-00";
					maskEditor_1.style.css({
						"text-align" : "center"
					});
					maskEditor_1.bind("value").toDataMap(app.lookup("product_output"), "write_date");
					container.addChild(maskEditor_1, {
						"colIndex": 1,
						"rowIndex": 0
					});
				})(group_5);
				container.addChild(group_5, {
					"top": "78px",
					"left": "416px",
					"width": "297px",
					"height": "38px"
				});
				var fileUpload_1 = new cpr.controls.FileUpload("file_upload");
				fileUpload_1.buttons = ["send"];
				fileUpload_1.readOnly = true;
				fileUpload_1.sendLabel = "다운로드";
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
				if(typeof onFile_uploadSendbuttonClick == "function") {
					fileUpload_1.addEventListener("sendbutton-click", onFile_uploadSendbuttonClick);
				}
				if(typeof onFile_uploadDownloadClick == "function") {
					fileUpload_1.addEventListener("download-click", onFile_uploadDownloadClick);
				}
				container.addChild(fileUpload_1, {
					"top": "174px",
					"left": "19px",
					"width": "695px",
					"height": "202px"
				});
				var group_6 = new cpr.controls.Container();
				// Layout
				var verticalLayout_1 = new cpr.controls.layouts.VerticalLayout();
				group_6.setLayout(verticalLayout_1);
				(function(container){
					var output_8 = new cpr.controls.Output();
					output_8.value = "내용";
					output_8.style.css({
						"padding-left" : "10px"
					});
					container.addChild(output_8, {
						"width": "100px",
						"height": "27px"
					});
					var textArea_1 = new cpr.controls.TextArea("outputContent");
					textArea_1.readOnly = true;
					textArea_1.style.css({
						"padding-top" : "5px",
						"padding-left" : "10px",
						"padding-bottom" : "5px",
						"padding-right" : "10px"
					});
					textArea_1.bind("value").toDataMap(app.lookup("product_output"), "output_content");
					container.addChild(textArea_1, {
						"width": "100px",
						"height": "225px"
					});
				})(group_6);
				container.addChild(group_6, {
					"top": "386px",
					"left": "19px",
					"width": "695px",
					"height": "264px"
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
				if(typeof onButtonClick == "function") {
					button_1.addEventListener("click", onButtonClick);
				}
				container.addChild(button_1, {
					"top": "660px",
					"left": "540px",
					"width": "80px",
					"height": "25px"
				});
				var button_2 = new cpr.controls.Button();
				button_2.value = "삭제";
				button_2.style.css({
					"border-right-style" : "none",
					"background-color" : "#DAF2DA",
					"border-radius" : "10px",
					"border-left-style" : "none",
					"border-bottom-style" : "none",
					"background-image" : "none",
					"border-top-style" : "none"
				});
				if(typeof onButtonClick2 == "function") {
					button_2.addEventListener("click", onButtonClick2);
				}
				container.addChild(button_2, {
					"top": "660px",
					"left": "633px",
					"width": "80px",
					"height": "25px"
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "0px",
				"left": "0px",
				"width": "725px",
				"height": "700px"
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
		}
	});
	app.title = "outputContentView";
	cpr.core.Platform.INSTANCE.register(app);
})();
