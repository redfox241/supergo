/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file utils Conf.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//解析文件，取出所有参数
func GetYamlConfig(filePath string) map[interface{}]interface{} {
	var strAppPath, strFileName string

	strAppPath = getParentDirectory(getCurrentDirectory())
	strFileName = strAppPath + "/" + filePath

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

//解析文件，取出所有参数
func GetAppConf(filePath string) map[interface{}]interface{} {
	var strAppPath, strFileName string

	//将解析出的参数转为map的形式
	mapAppConf := make(map[interface{}]interface{})

	strAppPath = getParentDirectory(getCurrentDirectory())
	strFileName = strAppPath + "/" + filePath

	data, err := ioutil.ReadFile(strFileName)

	if err != nil {
		LogErr("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &mapAppConf)
	return mapAppConf
}

func GetValuesByKey(key string, themap map[interface{}]interface{}) string {

	if value, ok := themap[key]; ok {

		fmt.Println(value)
		//return value
	}

	LogErr("Can't find the *.yaml")
	return ""
}
