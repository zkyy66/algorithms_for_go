/**
 * @Date 2022/7/2
 * @Name 76-windowSubString
 * @VariableName
**/
//https://leetcode.cn/problems/minimum-window-substring/solution/golangshuang-zhi-zhen-dan-ha-xi-biao-by-8cn0r/
package array_code

import "math"

func minWindows(s, t string) string {
	var result string
	//cnt := match.MaxInt32
	cnt := math.MaxInt32
	hashMap := make(map[byte]int)
	l, r := 0, 0
	for i := 0; i < len(t); i++ {
		hashMap[t[i]]++
	}
	for r < len(s) {
		hashMap[s[r]]--
		for check(hashMap) {
			if r-l+1 < cnt {
				cnt = r - l + 1
				result = s[l : r+1]
			}
			hashMap[s[l]]++
			l++
		}
		r++
	}
	return result
}
func check(hashMap map[byte]int) bool {
	for _, v := range hashMap {
		if v > 0 {
			return false
		}
	}
	return true
}
