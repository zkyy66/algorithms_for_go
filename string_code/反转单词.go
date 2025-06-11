package string_code

import (
	"strings"
)

func ReverseStringWord(s string) string {
	if len(s) == 0 {
		return ""
	}
	// 使用strings.Fields分割字符串为单词数组
	words := strings.Fields(s)
	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// 使用strings.Join将单词数组重新组合成字符串
	return strings.Join(words, " ")
}

//此方法也行
func Reverse1(s string) string {
	if len(s) == 0 {
		return ""
	}
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
