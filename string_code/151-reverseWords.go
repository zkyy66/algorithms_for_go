package string_code

import "strings"

// versionOne
func reverseWordsOne(s string) string {
	t := strings.Fields(s)
	for i := 0; i < len(t)/2; i++ {
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	newStr := strings.Join(t, " ")
	return newStr
}
func reverseWordsTwo(s string) string {
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	reverse151(&b, 0, len(b)-1)
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		
		}
		reverse151(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}
func reverse151(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
func main() {
	reverseWordsOne("yaoyuan")
	reverseWordsTwo("yaoyuan")
}
