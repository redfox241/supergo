/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file user.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package controllers

import (
	"models"
	"strconv"
	"user/user"
	"utils"
)

type userThrift struct {
}

func GetUserThrift() *userThrift {
	return &userThrift{}
}

func (this *userThrift) GetUserInfoByUserId(paramMap map[string]string) ([]*user.UserInfo, error) {

	var err error

	intUserId, _ := strconv.Atoi(paramMap["user_id"])
	newUser := make([]*user.UserInfo, 0)
	newUser, err = models.GetUserInfo(intUserId)

	utils.LogErr("-->from client Call:", paramMap)
	utils.LogNotice("finish to get user info.")

	return newUser, err
}

func (this *userThrift) GetUserInfo(paramMap map[string]string) ([]*user.UserInfo, error) {

	var err error
	
	intUserId, _ := strconv.Atoi(paramMap["user_id"])
	newUser := make([]*user.UserInfo, 0)
	newUser, err = models.GetUserInfo(intUserId)

	utils.LogErr("-->from client Call:", paramMap)
	utils.LogNotice("finish to get user info.")

	return newUser, err
}

func (this *userThrift) GetUserList(paramMap map[string]string) ([]*user.UserInfo, error) {

	var err error

	intLastId, _ := strconv.Atoi(paramMap["last_id"])
	newUser := make([]*user.UserInfo, 0)
	newUser, err = models.GetUserList(intLastId)

	utils.LogErr("-->from client Call:", paramMap)
	utils.LogNotice("finish to get user info.")

	return newUser, err
}

func (this *userThrift) CreateNewUser(paramMap map[string]string) (int64, error) {

	userinfo := make(map[string]string)
	userinfo["user_id"] = "0"
	userinfo["user_name"] = paramMap["user_name"]
	userinfo["nick_name"] = paramMap["nick_name"]
	userinfo["intro"] = paramMap["intro"]
	
	newUserId, intAffects, err := models.CreateNewUser(userinfo)
	if intAffects > 0 {
		return newUserId, err
	} else {
		return newUserId, err
	}

}
