package libs

import (
	"fmt"
	"time"
)

//type batuThrift struct {
//}

func (this *batuThrift) GetUserInfo(callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	for k, v := range paramMap {
		r = append(r, "key:"+k+"  value:"+v+"  ")
	}
	return
}
