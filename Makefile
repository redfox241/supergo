build:
	rm -f ./bin/go_client ./bin/go_server
	thrift -r --gen go ./data/thrift/batu.thrift
	thrift -r --gen php ./data/thrift/batu.thrift 
	go build  -o ./bin/go_client  ./client/go_client.go
	go build  -o ./bin/go_server  ./service/go_server.go
	cp -r  ./gen-php  ./clientphp/
	cp -r  ./gen-go/batu  /Users/xiaojing/projectcode/src/
	rm -rf gen-php gen-go

