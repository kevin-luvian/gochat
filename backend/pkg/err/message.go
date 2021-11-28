package err

var ErrMessages = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",

	ERROR_CODE_UNDEFINED: "unknown error",
}

func GetMsg(code int) string {
	if msg, ok := ErrMessages[code]; ok {
		return msg
	} else {
		return ErrMessages[ERROR_CODE_UNDEFINED]
	}
}
