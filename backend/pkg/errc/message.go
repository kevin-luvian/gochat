package errc

var ErrMessages = map[int]string{
	Success: "ok",
	Error:   "error",

	ErrorCodeUndefined: "unknown error",
}

func GetMsg(code int) string {
	if msg, ok := ErrMessages[code]; ok {
		return msg
	} else {
		return ErrMessages[ErrorCodeUndefined]
	}
}