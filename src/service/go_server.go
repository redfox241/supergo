package main

import (
	"controllers"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"user/user" //注意导入Thrift生成的接口包
	"utils"
)

const (
	APP_CONFIG = "/conf/app.conf"
)

func main() {

	//get conf
	appConfig := make(map[interface{}]interface{})
	appConfig = utils.GetYamlConfig(APP_CONFIG)

	ip_addr := utils.GetElement("ip_addr", appConfig)
	port := utils.GetElement("port", appConfig)

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(ip_addr + ":" + port)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := controllers.GetUserThrift()
	processor := user.NewUserProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server on :", ip_addr+":"+port)
	server.Serve()

}
