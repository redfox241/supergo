package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"strconv"
	"time"
	"user/user"
)

const (
	HOST = "127.0.0.1"
	PORT = "9090"
)

func main() {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := user.NewUserClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	model := user.UserInfo{
		100,
		"诸葛亮",
		"孔明",
		"鞠躬尽瘁，死而后已",
	}
	intNewUserId, _ := client.CreateNewUser(&model)

	fmt.Println("new_user_id : ", intNewUserId)

	for i := 0; i < 1; i++ {
		paramMap := make(map[string]string)
		paramMap["user_id"] = strconv.FormatInt(intNewUserId, 10)

		userInfo, _ := client.GetUserInfo(time.Now().Unix(), "go client", paramMap)
		fmt.Println("GOClient Call->", userInfo)
	}

	endTime := currentTimeMillis()
	fmt.Printf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
