package util

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

func JsonIter() jsoniter.API {
	return jsoniter.ConfigCompatibleWithStandardLibrary
}

func Read(file string, data interface{}) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		panic("[error] 文件，读取失败：" + err.Error())
	}
	err = JsonIter().Unmarshal(f, &data)
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

//写入文件

func WriteToJsonFile(fileName string, data interface{}) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		panic("[error] 打开文件" + fileName + "失败")
	}
	_, err = f.Write(toJson(data))
	if err != nil {
		panic("[error]写入文件" + fileName + "失败")
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
