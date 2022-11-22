package server

const (
	ResponseCodeUnknownError = -10000 - iota
	ResponseCodeParamParseError
	ResponseCodeParamNotEnough
	ResponseCodePasswordStrengthInvalid
	ResponseCodeDuplicated
	ResponseCodeDatabase
)

var (
	ResponseMsgUnknownError            = "系统错误"
	ResponseMsgParamParseError         = "参数解析失败"
	ResponseMsgParamNotEnough          = "参数不足"
	ResponseMsgPasswordStrengthInvalid = "密码强度不够"
	ResponseMsgDuplicated              = "数据重复"
	ResponseMsgDatabase                = "执行数据库语句失败"
)
