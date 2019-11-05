package result

const (
	SUCCESS             = 0     // 成功
	ServerEc            = -500  // 服务内部错误
	ParamEc             = -1000 // 参数校验错误
	CaptchaEc           = -1001 // 验证码过期或错误
	AccountEc           = -1002 // 账号或密码错误
	DataBaseEc          = -1003 // 数据库操作错误
	ForbiddenEc         = -1004 // 权限不足
	UnauthorizedEc      = -1005 // 未登陆
	InvalidTokenEc      = -1006 // 非法或过期的token
	PasswordUnchangedEc = -1007 // 密码和之前设置一样
)

var (
	CaptchaError = &Result{
		Code: CaptchaEc,
		Msg:  "验证码错误或过期",
	}
	AcountError = &Result{
		Code: AccountEc,
		Msg:  "账号或密码错误",
	}
	DatabaseError = &Result{
		Code: DataBaseEc,
		Msg:  "数据库操作错误",
	}
	UnauthorizedError = &Result{
		Code: UnauthorizedEc,
		Msg:  "未登陆",
	}
	ForbiddenError = &Result{
		Code: ForbiddenEc,
		Msg:  "无权限访问",
	}
	InvalidTokenError = &Result{
		Code: InvalidTokenEc,
		Msg:  "非法或过期的token,重新登陆",
	}
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewError(ec int, em string) (err *Result) {
	err = &Result{
		Code: ec,
		Msg:  em,
	}
	return
}

func NewSuccess(data interface{}) (success *Result) {
	success = &Result{
		Code: SUCCESS,
		Msg:  "SUCCESS",
	}
	success.Data = data
	return
}
