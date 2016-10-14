// Conf.go
package utils

import (
	"github.com/redfox241/supergo/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//解析文件，取出所有参数
func GetYamlConfig() map[interface{}]interface{} {

	data, err := ioutil.ReadFile("../conf/app.conf")

	//将解析出的参数转为map的形式
	m := make(map[interface{}]interface{})
	if err != nil {
		LogErr("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &m)

	return m
}

//根据需求取出对应值
func GetElement(key string, themap map[interface{}]interface{}) string {

	if value, ok := themap[key]; ok {
		//return value.(string)
		return value.(string)
	}

	LogErr("Can't find the *.yaml")
	return ""
}
