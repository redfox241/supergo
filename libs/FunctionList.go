package libs

import (
	"batu/demo"
	"fmt"
	"time"
)

type batuThrift struct {
}

func (this *batuThrift) GetUserInfo(callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	for k, v := range paramMap {
		r = append(r, "key:"+k+"  value:"+v+"  ")
	}
	return
}

func (this *batuThrift) CallBack(callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	r = append(r, "key:"+paramMap["a"]+"    value:"+paramMap["b"])
	return
}

func (this *batuThrift) Put(s *demo.Article) (err error) {
	fmt.Println("Article--->id: %d\tTitle:%s\tContent:%t\tAuthor:%d\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}

func (this *batuThrift) Process(s *demo.Article) (err error) {
	fmt.Println("Article--->id: %d\tTitle:%s\tContent:%t\tAuthor:%d\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}
