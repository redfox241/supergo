package controllers

import (
	"fmt"
	"time"
	"user/user"
	"utils"
)

type userThrift struct {
}

func GetUserThrift() *userThrift {
	return &userThrift{}
}

func (this *userThrift) GetUserInfo(callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	utils.LogNotice("finish to get user info.")
	utils.LogDebug("finish to get user info.")
	for k, v := range paramMap {
		r = append(r, "key:"+k+"  value:"+v+"  ")
	}
	return
}

func (this *userThrift) Process(newUser *user.UserInfo) (err error) {
	fmt.Printf("userinfo--->userId: %d\tuserName:%s\tnickName:%s\tintro:%s\n", newUser.UserId, newUser.UserName, newUser.NickName, newUser.Intro)
	return nil
}
