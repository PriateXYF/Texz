const FormatJSON = {
	name : "Format JSON",
	note : "格式化JSON字符串",
	func : function(str){
		var data
		try{
			data = JSON.parse(str)
		}catch(err){
			return "解析JSON失败!"
		}
		return JSON.stringify(data, null, "\t")
	}
}

CostomModules.functions.push(FormatJSON)