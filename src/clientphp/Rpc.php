<?php
/**
 * Thrift RPC - PHPClient
 * @author redfox241
 * @time 2016.10.13
 */


define("ROOT_DIR" , realpath(dirname(__FILE__).'/'));
define("GEN_DIR" , realpath(dirname(__FILE__).'/').'/Protocol');
require_once ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;
use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\TSocketPool;
use Thrift\Transport\TFramedTransport;
use Thrift\Transport\TBufferedTransport;


Class Rpc {

    static  private  function _init(){

    }

   static public function call($servicename,$method,$paramslist){
        $startTime = self::_getMillisecond();//记录开始时间

        $loader = new ThriftClassLoader();
        $loader->registerNamespace('Thrift',ROOT_DIR);
        $loader->registerDefinition($servicename.'\\'.$servicename, GEN_DIR);
        $loader->register();


        #require_once ( __DIR__ . "/config.php");

        $thriftHost = '127.0.0.1'; //UserServer接口服务器IP
        $thriftPort = 9090;            //UserServer端口

        $socket = new TSocket($thriftHost,$thriftPort);
        $socket->setSendTimeout(10000);#Sets the send timeout.
        $socket->setRecvTimeout(20000);#Sets the receive timeout.
        //$transport = new TBufferedTransport($socket); #传输方式：这个要和服务器使用的一致 [go提供后端服务,迭代10000次2.6 ~ 3s完成]
        $transport = new TFramedTransport($socket); #传输方式：这个要和服务器使用的一致[go提供后端服务,迭代10000次1.9 ~ 2.1s完成，比TBuffer快了点]
        $protocol = new TBinaryProtocol($transport);  #传输格式：二进制格式
       $ClassClient =  "\\".$servicename."\\".$servicename."\\".ucwords($servicename)."Client";

       //$client = new \user\user\UserClient($protocol);# 构造客户端
       $client = new $ClassClient($protocol);

        $transport->open();
        $socket->setDebug(true);

//       $s = new \user\user\UserInfo();

       $arrResult = $client->$method(time(),"php client",$paramslist);
       return $arrResult;

        $endTime = self::_getMillisecond();

        echo "本次调用用时: :".$endTime."-".$startTime."=".($endTime-$startTime)."毫秒<br>\n";

        $transport->close();
    }

    static private function _getMillisecond() {
        list($t1, $t2) = explode(' ', microtime());
        return (float)sprintf('%.0f', (floatval($t1) + floatval($t2)) * 1000);
    }

}


$arrUserInfo["user_id"] = 8;

$arrRes = Rpc::call("user","GetUserInfo",$arrUserInfo);

var_dump( $arrRes );
