/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file utils db.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/
package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/go-xorm/xorm"
)

var dbconnect *Engine

const (
	APP_CONFIG = "/conf/db.conf"
)

func GetDB() (*Engine, error) {

	if dbconnect != nil {

		return dbconnect, nil
	} else {
		LogDebug("db connect init ...")
	}

	//get conf
	appConfig := make(map[interface{}]interface{})
	appConfig = GetYamlConfig(APP_CONFIG)

	db_driver := GetElement("db_driver", appConfig)
	db_name := GetElement("db_name", appConfig)
	user := GetElement("user", appConfig)
	passwd := GetElement("passwd", appConfig)
	ipaddr := GetElement("ip_addr", appConfig)
	port := GetElement("port", appConfig)
	db_charset := GetElement("charset", appConfig)

	var err error
	strConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, passwd, ipaddr, port, db_name, db_charset)
	//fmt.Println(strConnect)
	dbconnect, err = NewEngine(db_driver, strConnect)
	dbconnect.ShowSQL(true)

	if err != nil {
		LogErr(fmt.Printf("failed to connect mysql.error:%s", err))
	}
	return dbconnect, err
}
