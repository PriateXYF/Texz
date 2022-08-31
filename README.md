<div align="center">
  <p>
    <img width="284" src="./build/appicon.png">
  </p>
  <h1><strong>Texz</strong></h1>
  <p>一个跨平台的轻量级文本处理小工具。</p>
  <p>支持自定义 JS 处理模块，基于 Go 语言开发。</p>
</div>



<div align="center">

![license](https://img.shields.io/github/license/priatexyf/texz?style=for-the-badge)
![version](https://img.shields.io/github/package-json/v/priatexyf/texz?filename=frontend%2Fpackage.json&style=for-the-badge)
![last](https://img.shields.io/github/last-commit/priatexyf/texz?style=for-the-badge)
![downloads](https://img.shields.io/github/downloads/priatexyf/texz/total?style=for-the-badge)
</div>

Texz 目前仍处于**开发**中，目前仅支持 Mac 下使用，请耐心等待。

## 开发

* Development

```bash
wails dev
```

* Build

```bash
wails build
```

## 快捷键

* `CMD + 1` : 切换到上一版本。
* `CMD + 2` : 切换到下一版本。
* `CMD + Enter/F` : 查询处理方法。
* `TAB` : 在查询处理方法界面切换所选模块。
* `CMD + R` : 重新载入应用程序。
* `CMD + D` : 清空当前文本框内容。
* `ALT + C` : 复制当前文本框内容。
* `CMD + S` : 将当前文本框内容保存到文件。 `TODO`
* `CMD + Z` : 还原文本框内容。
* `CMD + W` : 最小化程序。
* `CMD + Q` : 退出程序。


## 模块

### Decode (解码指定字符串)

* Base64 Decode : 对字符串进行 Base64 解码。
* Unicode Decode : 对 Unicode 字符串进行解码。
* URL Decode : 对字符串进行URL解码。

### Encode (编码指定字符串)

* Base64 Encode : 对字符串进行 Base64 解码。
* MD5 Crypto : 对字符串进行 MD5 加密。
* Unicode Encode : 对字符串进行 Unicode 编码。
* URL Encode : 对字符串进行 URL 编码。

### Extract (提取指定字符串)

* Extract Chinese : 提取所有中文字符。
* Extract Link : 提取所有链接。

### Remove (清除指定字符串)

* Remove Chinese : 移除所有中文字符。
* Remove Space : 移除所有空格。 `Custom`

### Format (格式化指定字符串)

* Format JSON : 格式化 JSON 字符串。 `Custom`
* Format GO : 格式化 GO 字符串。 `TODO`

### Convert (转换字符串格式)

* Markdown to HTML : Markdown 转 HTML。 `TODO`
* Cookie to JSON : 将 Cookie 字符串转为 JSON 格式。 `TODO`
* JSON to Cookie : 将 JSON 格式 Cookie 转为字符串。 `TODO`

## 自定义模块

### 编写模块

Texz 为用户提供了自定义模块处理的功能，可以通过 JavaScript 实现一个自己的模块。简单步骤如下 :

1. 点击右上角最后一个按钮，打开配置文件夹。
2. 进入 `modules` 文件夹。
3. 创建一个 `.js` 文件，编写你的模块。

编写模块非常简单，只需要实现一个文本处理方法即可，以下是两个模块示例。

* `RemoveSpace.js` : 移除文本中的所有空格。

```js
// RemoveSpace.js
const RemoveSpace = {
	name : "Remove Space",
	note : "移除空格",
	func : function(str){
		return str.replaceAll(' ', '')
	}
}

CostomModules.functions.push(RemoveSpace)
```

* `FormatJSON.js` : 格式化 JSON 字符串。

```js
// FormatJSON.js
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
```

### 调试模块

调试模块同样非常简单，在每次修改模块文件后，只需按下 `CMD + R` 或点击右上角倒数第二个按钮重新载入应用，新的模块文件即可生效。

### 默认模块

默认模块即为应用自带模块，具体内容可参考 [模块](#模块) 章节，默认模块分为**应用默认模块**与**自定义默认模块**，二者区别如下 :

1. 应用默认模块基于 Go 语言实现，**无法修改和删除**。
2. 自定义默认模块基于 JS 实现，**可以修改和删除**，与用户自定义模块的唯一区别是，当配置文件夹下的 `modules` 文件夹被删除时，重新打开应用会**自动创建**全部的自定义默认模块。

## 演示

![](https://i.imgur.com/Iefagkt.png)
![](https://i.imgur.com/Im4yjEy.png)

