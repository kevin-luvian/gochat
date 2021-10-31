package util

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

func AnyToStringQuery(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return "'" + v + "'"
	case bool:
		return strconv.FormatBool(v)
		// case datetime.DateTime:
		// 	return v.String()
	default:
		logrus.Panic("Type of ", v, " is not supported!!")
		return ""
	}
}
