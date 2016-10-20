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
	1: i64 user_id,
    2: string user_name,
    3: string nick_name,
	4: string intro,
}

/*
* service定义
*/
service User {
    list<UserInfo> GetUserInfo(1:i64 callTime, 2:string name, 3:map<string, string> paramMap),
	i64 CreateNewUser(1: UserInfo newUser),
}
