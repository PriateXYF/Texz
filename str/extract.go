package str

import (
	cregex "github.com/mingrammer/commonregex"

	"regexp"
)

var Extract = Module{
	Name: "Extract",
	Note: "提取指定字符",
	Functions: []Function{
		ExtractChinese,
		ExtractLink,
	},
}

var ExtractChinese = Function{
	Name: "Extract Chinese",
	Note: "提取所有中文字符",
	Body: func(rawText string) []string {
		re := regexp.MustCompile("[\u4e00-\u9fa5]+")
		byteList := re.FindAll([]byte(rawText), -1)
		var resList []string
		for _, text := range byteList {
			resList = append(resList, string(text))
		}
		return resList
	},
}

var ExtractLink = Function{
	Name: "Extract Link",
	Note: "提取所有链接",
	Body: func(rawText string) []string {
		resList := cregex.Links(rawText)
		return resList
	},
}
