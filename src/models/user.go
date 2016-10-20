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
 */
func init() {

	//get conf
	appConfig := make(map[interface{}]interface{})
	appConfig = utils.GetYamlConfig(APP_CONFIG)

	db_driver := utils.GetElement("db_driver", appConfig)
	db_name := utils.GetElement("db_name", appConfig)
	user := utils.GetElement("user", appConfig)
	passwd := utils.GetElement("passwd", appConfig)
	ipaddr := utils.GetElement("ip_addr", appConfig)
	port := utils.GetElement("port", appConfig)
	db_charset := utils.GetElement("charset", appConfig)

	var err error
	strConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, passwd, ipaddr, port, db_name, db_charset)
	fmt.Println(strConnect)
	engine, err = xorm.NewEngine(db_driver, strConnect)
	engine.ShowSQL(true)

	if err != nil {
		utils.LogErr(fmt.Printf("failed to connect mysql.error:%s", err))
	}

}

/**
*获取用户信息
 */

func GetUserInfo(userId int) ([]*user.UserInfo, error) {

	userInfo := make([]*user.UserInfo, 0)
	errmu := engine.Sql("select * from user where user_id = ? ", userId).Find(&userInfo)

	if errmu != nil {
		utils.LogErr("failed to get User info, input_param:%s", errmu)
	}

	fmt.Println(userInfo)

	return userInfo, errmu
}

/**
*获取用户list,支持下拉分页
 */
func GetUserList(lastId int) ([]user.UserInfo, error) {

	userList := make([]user.UserInfo, 0)

	var strSqlCmd, strSqlCmdPrefix string
	strSqlCmdPrefix = "select * from user "

	if lastId > 0 {
		strSqlCmd = fmt.Sprintf(strSqlCmdPrefix+" where user_id > %d ", lastId)

	} else {
		strSqlCmd = strSqlCmd
	}

	utils.LogDebug(strSqlCmd)
	err := engine.Sql(strSqlCmd).Find(&userList)
	//errmn := engine.AllCols().In("id", 5, 6, 7).Find(&user)

	if err != nil {
		utils.LogErr(fmt.Printf("failed to get data info from mysql.error:%s", err))
	}
	for _, val := range userList {
		//fmt.Println(key)
		fmt.Println(val)
	}

	return userList, err
}

/**
* 新创建用户
 */
func CreateNewUser(paramMap map[string]string) (int64, int64, error) {

	user := new(User)
	user.User_id = 0
	user.User_name = paramMap["user_name"]
	user.Nick_name = paramMap["nick_name"]
	user.Intro = paramMap["intro"]

	intAffected, err := engine.Insert(user)

	return user.User_id, intAffected, err

}
