package service

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
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
	//err = json.Unmarshal(f, &data)
	//if err != nil {
	//	panic("[error] 项目配置文件，无法解析成预期的格式")
	//}

	var json2 = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json2.Unmarshal(f, &data)
	if err != nil {
		return
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
