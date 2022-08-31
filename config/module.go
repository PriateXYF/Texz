package config

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
)

//go:embed embed/modules
var modules embed.FS

// 获取modules文件夹的本地路径
func GetModlueJSPath() string {
	modulejsPath := ExecutablePathJoin("modules")
	_, err := ReadFile(modulejsPath)
	if err != nil {
		GenerateModulesDir()
		GenerateModulesFile()
	}
	return modulejsPath
}

// 生成 modules JS 文件
func GenerateModulesFile() error {
	fs.WalkDir(modules, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			module, err := modules.ReadFile(path)
			if err != nil {
				return err
			}
			modulejsPath := ExecutablePathJoin(fmt.Sprintf("modules/%s", d.Name()))
			err = OutputFile(modulejsPath, string(module))
			return err
		}
		return nil
	})
	return nil
}

// 生成默认JS模块
func GenerateModulesDir() error {
	modulesPath := ExecutablePathJoin("modules")
	err := MakeDir(modulesPath)
	if err != nil {
		return err
	}
	return nil
}

// 获取远程地址
func GetRemoteModlues() []string {
	result := []string{}
	modulesPath := GetModlueJSPath()
	fileinfo, err := ioutil.ReadDir(modulesPath)
	if err != nil {
		panic(err)
	}
	for _, fi := range fileinfo {
		if fi.IsDir() {
			continue
		}
		spFile := strings.Split(fi.Name(), ".")
		if len(spFile) == 0 || spFile[len(spFile)-1] != "js" {
			continue
		}
		result = append(result, fmt.Sprintf("http://localhost:9090/%s", fi.Name()))
	}
	return result
}
