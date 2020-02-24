package utils

import "errors"

// 服务器和客户端的交流错误码
const (
	ErrCodeSuccess            = iota
	ErrCodeParamter           = 1001
	ErrCodeUserExit           = 1002
	ErrCodeServerBusy         = 1003
	ErrCodeUserNotExist       = 1004
	ErrCodeUserNameOrPwdWrong = 1005
	ErrCodeCaptionSensitive   = 1006
	ErrCodeContentSensitive   = 1007
	ErrCodeNotLogin			  = 1008
)

// 程序内部的错误交流模块
var (
	ErrUserAlreadyExisted = errors.New("user not existed")
	ErrUserNotExisted     = errors.New("user not existed")
	ErrUserPwdWrong       = errors.New("user passwd wrong")
)

// 根据错误码获取对应的消息
func GetMessage(code int) (message string) {
	switch code {
	case ErrCodeParamter:
		message = "参数错误"
	case ErrCodeSuccess:
		message = " Success "
	case ErrCodeUserExit:
		message = "用户名已存在"
	case ErrCodeServerBusy:
		message = "服务器繁忙"
	case ErrCodeUserNotExist:
		message = "不户名不存在"
	case ErrCodeUserNameOrPwdWrong:
		message = "用户名或密码错误"
	case ErrCodeCaptionSensitive:
		message = "标题不能含有敏感词"
	case ErrCodeContentSensitive:
		message = "内容中含有敏感词!请修改后再提交"
	case ErrCodeNotLogin:
		message = "未登陆，请重新登录"
	default:
		message = "未知错误"
		break
	}
	return
}
