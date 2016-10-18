/**
 * UserThrift TDL
 * @author redfox241
 * @time 2016.10.18
 */

namespace go user.user
namespace php user.user


/**
 * 结构体定义
 */
struct UserInfo{
    1: i32 userId,
    2: string userName,
    3: string nickName,
    4: string portrait,
	5: string intro,
	6: string mobileNo,
	7: string email
}

/*
* service定义
*/
service User {
    list<string> GetUserInfo(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
    void Process(1: UserInfo newUser),
}
