package main

import (
	"encoding/json"
	"errors"
	"os"
)

// 目录是否存在
func IsExistDir(dirPath string) (bool, error) {
	f, err := os.Stat(dirPath)
	if err == nil {
		if f.IsDir() {
			return true, nil
		}
		return false, errors.New(dirPath + " 不是目录")
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 创建文件夹
func MkDir(dirPath string) error {
	exist, err := IsExistDir(dirPath)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return errors.New("目录创建失败：" + err.Error())
	}

	return nil
}

//写入文件

func WriteToJsonFile(dir string, name string, data interface{}) {
	file, err := os.OpenFile(dir+"/"+name, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		panic("[error] 打开文件" + dir + "/" + name + "失败")
	}
	_, err = file.Write(toJson(data))
	if err != nil {
		panic("[error]写入文件" + dir + "/" + name + "失败")
	}
}

//解析成json
func toJson(data interface{}) []byte {
	marshal, err := json.Marshal(data)
	if err != nil {
		panic("[error] 解析失败")
	}
	return marshal
}
