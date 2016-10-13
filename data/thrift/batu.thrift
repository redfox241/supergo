/**
 * BatuThrift TDL
 * @author liuxinming
 * @time 2015.5.13
 */

namespace go batu.demo
namespace php batu.demo

/**
 * 结构体定义
 */
struct Article{
    1: i32 id, 
    2: string title,
    3: string content,
    4: string author,
}

const map<string,string> MAPCONSTANT = {'hello':'world', 'goodnight':'moon'}

service batuThrift {        
    list<string> CallBack(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
    list<string> GetUserInfo(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
    void Put(1: Article newArticle),
    void Process(1: Article newArticle),
}
