package str

import "strings"

// 基础函数(将一个字符串转为另一个字符串)
type Function struct {
	// 函数名称
	Name string `json:"name"`
	// 函数备注
	Note string `json:"note"`
	// 函数内容
	Body interface{} `json:"-"`
}

// 模块
type Module struct {
	Name      string     `json:"name"`
	Note      string     `json:"note"`
	Functions []Function `json:"functions"`
}

func (f Function) Call(rawText string) string {
	function := f.Body
	switch fn := function.(type) {
	case func(string) string:
		return fn(rawText)
	case func(string) []string:
		return strings.Join(fn(rawText), "\n")
	default:
		return "处理发生异常!函数 " + f.Name + " 返回不正确!"
	}
}

var modules = []Module{
	Decode,
	Encode,
	Extract,
	Remove,
}

// 获取所有模块
func GetModules() []Module {
	return modules
}

// 通过函数名称获取函数
func GetFunction(funcName string) *Function {
	for _, module := range modules {
		for _, function := range module.Functions {
			if function.Name == funcName {
				return &function
			}
		}
	}
	return nil
}
