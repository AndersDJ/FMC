package errCode

//ErrCode  错误结构
type ErrCode struct {
	Code int
	Msg  string
}

//定义错误code
var (
	Success = &ErrCode{Code: 0, Msg: "成功"}

	AuthErr            = &ErrCode{Code: 1, Msg: "认证错误"}
	GetCookieTokenErr  = &ErrCode{Code: 1, Msg: "获取CookieToken错误"}
	GetEsbTokenErr     = &ErrCode{Code: 1, Msg: "获取EsbToken错误"}
	AuthCookieTokenErr = &ErrCode{Code: 1, Msg: "AuthCookieToken错误"}
	GetUserInfoErr     = &ErrCode{Code: 1, Msg: "获取用户信息失败"}

	AuthNotFound     = &ErrCode{Code: 1, Msg: "未登录"}
	AuthTimeout      = &ErrCode{Code: 2, Msg: "登录超时"}
	ProtoFormatErr   = &ErrCode{Code: 3, Msg: "协议格式错误"}
	ProtoValidataErr = &ErrCode{Code: 4, Msg: "协议格式验证错误"}
	ParamErr         = &ErrCode{Code: 9, Msg: "参数错误"}
	LoginErr         = &ErrCode{Code: 5, Msg: "用户名或密码错误"}

	DBAddTaskErr = &ErrCode{Code: 5, Msg: "往数据库中添加任务失败"}
	DBErr        = &ErrCode{Code: 6, Msg: "查询数据似乎出现问题了"}
	DBScanErr    = &ErrCode{Code: 7, Msg: "解析查询数据似乎出现问题了"}
	DBDelErr     = &ErrCode{Code: 8, Msg: "删除数据错误"}
	TableNameErr = &ErrCode{Code: 10, Msg: "表不存在或没有授权"}

	CMDErr       = &ErrCode{Code: 11, Msg: "指令没授权"}
	SaltTokenErr = &ErrCode{Code: 11, Msg: "salt获取token失败"}
	CallSaltErr  = &ErrCode{Code: 11, Msg: "调用salt失败"}
	RspFormatErr = &ErrCode{Code: 11, Msg: "返回数据格式错误"}

	FlushGameListLockErr     = &ErrCode{Code: 11, Msg: "有人正在同步, 请稍后再试"}
	FlushGameListErr         = &ErrCode{Code: 11, Msg: "从CMDB获取游戏列表失败"}
	FlushGameListValidateErr = &ErrCode{Code: 11, Msg: "从CMDB获取游戏列表数据格式错误"}
	UpdateGameListErr        = &ErrCode{Code: 11, Msg: "更新DB数据失败"}
)
