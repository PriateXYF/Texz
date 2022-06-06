package str

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
)

var Encode = Module{
	Name: "Encode",
	Note: "编码指定字符",
	Functions: []Function{
		EncodeURL,
		EncodeMD5,
		EncodeBase64,
		EncodeUnicode,
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
		data := []byte(rawText)
		has := md5.Sum(data)
		resText := fmt.Sprintf("%x", has)
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
var EncodeUnicode = Function{
	Name: "Unicode Encode",
	Note: "对字符进行Unicode编码",
	Body: func(rawText string) string {
		textQuoted := strconv.QuoteToASCII(rawText)
		textUnquoted := textQuoted[1 : len(textQuoted)-1]
		return textUnquoted
	},
}
