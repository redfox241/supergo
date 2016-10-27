appname=$1
echo $appname

##判断是否存在output,不存在，创建output
if [ ! -d "./output" ]; then
    mkdir -p output
fi

if [ ! -d "./output/$appname" ]; then
    mkdir -p ./output/$appname
fi

cp -r ../src  ./output/$appname


