package str

import (
	"encoding/base64"
	"net/url"
	"regexp"
)

var Encode = Module{
	Name: "Encode",
	Note: "编码指定字符",
	Functions: []Function{
		EncodeURL,
		EncodeMD5,
		EncodeBase64,
	},
}
var EncodeURL = Function{
	Name: "URL Encode",
	Note: "对字符进行URL编码",
	Body: func(rawText string) string {
		resText := url.QueryEscape(rawText)
		return resText
	},
}
var EncodeMD5 = Function{
	Name: "MD5 Crypto",
	Note: "对字符进行MD5加密",
	Body: func(rawText string) string {
		re := regexp.MustCompile("[\u4e00-\u9fa5]")
		resText := string(re.ReplaceAll([]byte(rawText), []byte("")))
		return resText
	},
}
var EncodeBase64 = Function{
	Name: "Base64 Encode",
	Note: "对字符Base64编码",
	Body: func(rawText string) string {
		resText := base64.StdEncoding.EncodeToString([]byte(rawText))
		return resText
	},
}
