package response

import (
	"github.com/chuangxinyuan/fengchu/pkg/model/response/status"
	"github.com/google/uuid"
	"time"
)

// R 通用分页返回对象
type R struct {
	Code      status.Code `json:"code"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
	Timestamp string      `json:"timestamp"`
	Exception string      `json:"exception"`
}

func Result(code status.Code, success bool, data interface{}, message string, exception string) R {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return R{
		code,
		success,
		data,
		message,
		uuid.NewString(),
		currentTime,
		exception,
	}
}

func Success() R {
	return Result(status.Success, true, map[string]interface{}{}, status.Message(status.Success), "")
}

func SuccessWithMessage(message string) R {
	return Result(status.Success, true, map[string]interface{}{}, message, "")
}

func SuccessWithData(data interface{}) R {
	return Result(status.Success, true, data, status.Message(status.Success), "")
}

func SuccessWithDetailed(data interface{}, message string) R {
	return Result(status.Success, true, data, message, "")
}

//func Fail(code status.Code) {
//	Result(code, false, map[string]interface{}{}, status.Message(code), c.Errors.String())
//}

//func FailWithMessage(code status.Code, message string) {
//	Result(code, false, map[string]interface{}{}, message, c.Errors.String(), c)
//}
