/**
 * UserThrift TDL
 * @author redfox241
 * @time 2016.10.18
 */

namespace go  demo.demo
namespace php demo.demo


/**
 * 结构体定义
 */
struct DemoInfo{
	1:  i64 demo_id,
    2:  string demo_name,
}

/*
* service定义
*/
service Demo {
    list<DemoInfo> GetDemoInfo(1:map<string, string> paramMap),
	i64 ProcessDemo(1: map<string,string> paramMap),
	i64 DoDemo(1: map<string,string> paramMap),
}
