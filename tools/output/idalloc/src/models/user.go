/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file models user.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"user/user"
	"utils"
)

var engine *xorm.Engine

const (
	APP_CONFIG = "/conf/db.conf"
)

type UserOld struct {
	User_id       int64
	Role_id       int8
	From_type     string
	User_name     string
	Nick_name     string
	Sex           int8
	Mobile_prefix int
	Mobile_no     string
	Openid        string
	Portrait      string
	Email         string
	Intro         string
	Lng           float64
	Lat           float64
	Cur_city_id   int64
	City_id       int32
	Follow_num    int64
	Followed_num  int64
	Feed_num      int64
	Login_time    int64
	Create_time   int64
	Ext           string
	Member_time   int64
	Is_official   int8
	Birthday      int64
	Update_time   int64
	Uuid          string
}

type User struct {
	User_id   int64  `xorm:"user_id pk autoincr"`
	User_name string `xorm:"user_name"`
	Nick_name string `xorm:"nick_name"`
	Intro     string `xorm:"intro"`
}

/**
 * 初始化方法
 **/
func init() {

	//初始化数据库连接
	engine, _ = utils.GetDB()
}

/**
*获取用户信息
 */

func GetUserInfo(userId int) ([]*user.UserInfo, error) {

	newUser := make([]User, 0)
	errmu := engine.Sql("select * from user where user_id = ? ", userId).Find(&newUser)
	//errmu := engine.Table("user").In("user_id", 200, 201, 202).Find(&newUser)

	if errmu != nil {
		utils.LogErr("failed to get User info, input_param:%s", errmu)
	}

	userList := make([]*user.UserInfo, len(newUser))

	for k, v := range newUser {
		userInfo := new(user.UserInfo)
		userInfo.UserID = v.User_id
		userInfo.UserName = v.User_name
		userInfo.NickName = v.Nick_name
		userInfo.Intro = v.Intro
		userList[k] = userInfo
		userInfo = nil
	}

	return userList, errmu
}

/**
*获取用户list,支持下拉分页
 */
func GetUserList(lastId int) ([]*user.UserInfo, error) {

	newUser := make([]User, 0)

	var strSqlCmd string

	if lastId > 0 {
		strSqlCmd = fmt.Sprintf(" select * from user  where user_id < %d order by user_id desc limit 20", lastId)

	} else {
		strSqlCmd = fmt.Sprintf(" select * from user  order by user_id desc limit 20")
	}

	utils.LogDebug(strSqlCmd)
	err := engine.Sql(strSqlCmd).Find(&newUser)
	//errmn := engine.AllCols().In("id", 5, 6, 7).Find(&user)

	if err != nil {
		utils.LogErr(fmt.Printf("failed to get data info from mysql.error:%s", err))
	}

	userList := make([]*user.UserInfo, len(newUser))
	for key, val := range newUser {

		userInfo := new(user.UserInfo)
		userInfo.UserID = val.User_id
		userInfo.UserName = val.User_name
		userInfo.NickName = val.Nick_name
		userInfo.Intro = val.Intro
		userList[key] = userInfo

		userInfo = nil
	}

	return userList, err
}

/**
* 新创建用户
 */
func CreateNewUser(paramMap map[string]string) (int64, int64, error) {

	user := new(User)
	//user.User_id = 0
	user.User_name = paramMap["user_name"]
	user.Nick_name = paramMap["nick_name"]
	user.Intro = paramMap["intro"]

	intAffected, err := engine.Insert(user)

	return user.User_id, intAffected, err

}
