package status

type Code int

const (
	Success             Code = 200 //操作成功
	BadRequest          Code = 400 //请求参数不正确
	Unauthorized        Code = 401 //账号未登录
	Forbidden           Code = 403 //没有该操作权限
	NotFound            Code = 404 //请求未找到
	MethodNotAllowed    Code = 405 //请求方法不正确
	Locked              Code = 423 //请求失败，请稍后重试
	TooManyRequests     Code = 429 //请求过于频繁，请稍后重试
	InternalServerError Code = 500 //系统异常
	RepeatedRequests    Code = 900 //重复请求，请稍后重试
	DemoDeny            Code = 901 //演示模式，禁止写操作
	Unknown             Code = 999 //未知错误
)

var statusMessage = map[Code]string{
	Success:             "操作成功",
	BadRequest:          "请求参数不正确",
	Unauthorized:        "账号未登录",
	Forbidden:           "没有该操作权限",
	NotFound:            "请求未找到",
	MethodNotAllowed:    "请求方法不正确",
	Locked:              "请求失败，请稍后重试",
	TooManyRequests:     "请求过于频繁，请稍后重试",
	InternalServerError: "系统异常",
	RepeatedRequests:    "重复请求，请稍后重试",
	DemoDeny:            "演示模式，禁止写操作",
	Unknown:             "未知错误",
}

func Message(code Code) string {
	return statusMessage[code]
}

const (
	DefaultEmptyMessage  = "暂无承载数据"
	DefaultSucceeMessage = "操作成功"
	DefaultFailMessage   = "操作失败"
)
