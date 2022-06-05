package str

import (
	"regexp"
)

var Remove = Module{
	Name: "Remove",
	Note: "清除指定字符",
	Functions: []Function{
		RemoveChinese,
	},
}

var RemoveChinese = Function{
	Name: "Remove Chinese",
	Note: "移除所有中文字符",
	Body: func(rawText string) string {
		re := regexp.MustCompile("[\u4e00-\u9fa5]")
		resText := string(re.ReplaceAll([]byte(rawText), []byte("")))
		return resText
	},
}
