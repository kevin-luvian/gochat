package errc

var ErrMessages = map[int]string{
	Success: "ok",
	Error:   "error",

	InvalidParams:      "request parameter is not valid",
	ErrorCodeUndefined: "unknown error",
}

func GetMsg(code int) string {
	if msg, ok := ErrMessages[code]; ok {
		return msg
	} else {
		return ErrMessages[ErrorCodeUndefined]
	}
}
