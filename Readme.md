
supergo 一个基于thrift的golang研发框架，支持go+go，php+go应用

目前主要是支持service端研发，不支持数据的api接口

#####一、使用生成工具生成demo

    1、书写thrift文件（按照thrift语法）
		
		namespace go  demo.demo
		namespace php demo.demo
		
		
		/**
		  结构体定义
		 */
		struct DemoInfo{
			1:  i64 demo_id,
		    2:  string demo_name,
		}
		
		/*
		  service定义
		*/
		service Demo {
		    list<DemoInfo> GetDemoInfo(1:map<string, string> paramMap),
			i64 ProcessDemo(1: map<string,string> paramMap),
		}

	2、使用demo生成工具生成demo
		
		cd tools
		sh genapptools.sh demo
			
			output
			|	demo
			|	|	bin
			|	|	pkg
			|__	|__	src
					|	Makefile    
					|	bin         
					|	client      
					|	clientphp   
					|	conf        
					|	controllers 
					|	libs        
					|	load        
					|	models      
					|	routers     
					|	service     
					|	utils
	

		
	  3、编译
		make 
		启动server ./bin/go_server
		使用goclient测试 ./bin/go_client
		使用php测试  php ./clientphp/demo.php
		  
		  