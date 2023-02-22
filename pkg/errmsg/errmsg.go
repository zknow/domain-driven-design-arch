package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 自定義 ErrorCode..
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "Fail",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
