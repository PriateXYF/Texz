package str

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

var Decode = Module{
	Name: "Decode",
	Note: "解码指定字符",
	Functions: []Function{
		DecodeURL,
		DecodeBase64,
		DecodeUnicode,
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
			return "解码失败!请检查输入格式是否正确!"
		}
		return string(resText)
	},
}

var DecodeUnicode = Function{
	Name: "Unicode Decode",
	Note: "对Unicode字符进行解码",
	Body: func(rawText string) string {
		sUnicodev := strings.Split(rawText, "\\u")
		var resText string
		for _, v := range sUnicodev {
			if len(v) < 1 {
				continue
			}
			temp, err := strconv.ParseInt(v, 16, 32)
			if err != nil {
				return "解码失败!请检查输入格式是否正确!"
			}
			resText += fmt.Sprintf("%c", temp)
		}
		return resText
	},
}
