/**
 * IdallocThrift TDL
 * @author redfox241
 * @time 2016.10.18
 */

namespace go  idalloc.idalloc
namespace php idalloc.idalloc

/**
 * 结构体定义
 */
struct IdallocInfo{
	1:  i64 id,
}

/*
* service定义
*/
service Idalloc {
	i64 GenId(1: map<string,string> paramMap),
}


