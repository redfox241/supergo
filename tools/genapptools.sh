appname=$1

output="./output"
datapath="../data"
gopath="/Users/xiaojing/projectcode/src/"

echo "生成 "$appname" 模块基本开发框架"

##判断是否存在output,不存在，创建output
if [ ! -d $output ]; then
    mkdir -p $output
fi

if [ ! -d "$output/$appname" ]; then
    mkdir -p $output/$appname
fi

## copy source code,为支持多次调整thrift数据修改的，源码支持增量调整
if [ ! -d "$output/$appname/src" ];then
    cp -r  ../src  $output/$appname/
fi

## 生成thrift文件
cd $output

thrift -r --gen go $datapath/$appname/$appname.thrift
thrift -r --gen php $datapath/$appname/$appname.thrift


cp -r  ./gen-go/$appname  $gopath
cp -r  ./gen-php/*  ./$appname/src/clientphp/Protocol/


## 创建bin pkg文件夹

if [ ! -d "./$appname/bin" ]; then
    mkdir -p ./$appname/bin
fi

if [ ! -d "./$appname/pkg" ]; then
    mkdir -p ./$appname/pkg
fi


if [ -d "./gen-php" ]; then
    rm -rf gen-php
fi

if [ -d "./gen-go" ]; then
    rm -rf gen-go
fi

echo "模块 "$appname" 开发模块生成完毕"
