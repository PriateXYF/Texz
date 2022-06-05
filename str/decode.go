package str

import (
	"encoding/base64"
	"net/url"
)

var Decode = Module{
	Name: "Decode",
	Note: "解码指定字符",
	Functions: []Function{
		DecodeURL,
		DecodeBase64,
	},
}

var DecodeURL = Function{
	Name: "URL Decode",
	Note: "对字符进行URL解码",
	Body: func(rawText string) string {
		enEscapeUrl, _ := url.QueryUnescape(rawText)
		return enEscapeUrl
	},
}

var DecodeBase64 = Function{
	Name: "Base64 Decode",
	Note: "对字符进行Base64解码",
	Body: func(rawText string) string {
		resText, err := base64.StdEncoding.DecodeString(rawText)
		if err != nil {
			return ""
		}
		return string(resText)
	},
}
