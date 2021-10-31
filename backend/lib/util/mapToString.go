package util

import "strings"

func MapToString(maps *map[string]string, using string, join string) string {
	mstring := make([]string, 0, len(*maps))
	for k, v := range *maps {
		mstring = append(mstring, k+using+v)
	}
	return strings.Join(mstring, join)
}
