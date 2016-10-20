package controllers

import (
	"models"
	"strconv"
	"time"
	"user/user"
	"utils"
)

type userThrift struct {
}

func GetUserThrift() *userThrift {
	return &userThrift{}
}

func (this *userThrift) GetUserInfo(callTime int64, name string, paramMap map[string]string) ([]*user.UserInfo, error) {

	var err error

	intUserId, _ := strconv.Atoi(paramMap["user_id"])
	newUser := make([]*user.UserInfo, 0)
	newUser, err = models.GetUserInfo(intUserId)

	utils.LogErr("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	utils.LogNotice("finish to get user info.")

	return newUser, err
}

func (this *userThrift) CreateNewUser(newUser *user.UserInfo) (int64, error) {

	userinfo := make(map[string]string)
	userinfo["user_id"] = "0"
	userinfo["user_name"] = newUser.UserName
	userinfo["nick_name"] = newUser.NickName
	userinfo["intro"] = newUser.Intro

	newUserId, intAffects, err := models.CreateNewUser(userinfo)
	if intAffects > 0 {
		return newUserId, err
	} else {
		return newUserId, err
	}

}
