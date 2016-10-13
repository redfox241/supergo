<?php

$startTime = getMillisecond();//记录开始时间

$ROOT_DIR = realpath(dirname(__FILE__).'/');
$GEN_DIR = realpath(dirname(__FILE__).'/').'/gen-php';
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;
use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\TSocketPool;
use Thrift\Transport\TFramedTransport;
use Thrift\Transport\TBufferedTransport;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift',$ROOT_DIR);
$loader->registerDefinition('batu\demo', $GEN_DIR);
$loader->register();



?>
