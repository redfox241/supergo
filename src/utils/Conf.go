// Conf.go
package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//解析文件，取出所有参数
func GetYamlConfig(filePath string) map[interface{}]interface{} {
	var strAppPath, strFileName string

	strAppPath, _ = os.Getwd()
	strFileName = strAppPath + "/../" + filePath

	data, err := ioutil.ReadFile(strFileName)

	//将解析出的参数转为map的形式
	mapAppConf := make(map[interface{}]interface{})
	if err != nil {
		LogErr("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &mapAppConf)
	return mapAppConf
}

//根据需求取出对应值
func GetElement(key string, themap map[interface{}]interface{}) string {

	if value, ok := themap[key]; ok {
		return value.(string)
	}

	LogErr("Can't find the *.yaml")
	return ""
}

func GetValuesByKey(key string, themap map[interface{}]interface{}) string {

	if value, ok := themap[key]; ok {

		fmt.Println(value)
		//return value
	}

	LogErr("Can't find the *.yaml")
	return ""
}

func getCurrentDirectory() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
