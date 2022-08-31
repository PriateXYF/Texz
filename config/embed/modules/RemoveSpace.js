const RemoveSpace = {
	name : "Remove Space",
	note : "移除空格",
	func : function(str){
		return str.replaceAll(' ', '')
	}
}

CostomModules.functions.push(RemoveSpace)