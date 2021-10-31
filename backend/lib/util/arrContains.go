package util

func ArrStringContains(arr []string, o string) bool {
	for _, a := range arr {
		if a == o {
			return true
		}
	}
	return false
}
