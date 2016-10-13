package main

import (
	"batu/demo" //注意导入Thrift生成的接口包
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"time"
)

const (
	NetworkAddr = "127.0.0.1:9090" //监听地址&端口
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

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &batuThrift{}
	processor := demo.NewBatuThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
