package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Read(file string, data interface{}) {
	// 1.1 读取配置文件
	f, err := ioutil.ReadFile(file)
	if err != nil {
		panic("[error] 文件，读取失败：" + err.Error())
	}

	// 1.2 json解析公共库配置
	err = json.Unmarshal(f, &data)
	if err != nil {
		panic("[error] 项目配置文件，无法解析成预期的格式")
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
