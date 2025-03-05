package string_code

import "strings"

func stringPath(s string) string {
	if len(s) == 0 {
		return ""
	}
	strLen := []byte(s)
	res := make([]byte, 0)
	for i := 0; i <= len(strLen); i++ {
		if strLen[i] == '.' {
			res = append(res, []byte(" ")...)
		} else {
			res = append(res, strLen[i])
		}
	}
	return string(res)
}

func stringReplaceFun(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.Replace(s, ".", " ", -1)
	return s
}
