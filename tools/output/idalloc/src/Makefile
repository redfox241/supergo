build:
	rm -f ./bin/go_client ./bin/go_server
	thrift -r --gen go ./data/thrift/user.thrift
	thrift -r --gen php ./data/thrift/user.thrift
	cp -r  ./gen-go/user  /Users/xiaojing/projectcode/src/
	go build  -o ./bin/go_client  ./client/go_client.go
	go build  -o ./bin/go_server  ./service/go_server.go
	cp -r  ./gen-php/*  ./clientphp/Protocol/
	rm -rf gen-php gen-go

