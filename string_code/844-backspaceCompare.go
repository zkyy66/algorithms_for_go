/**
 * @Date 2022/6/28
 * @Name 844-backspaceCompare
 * @VariableName
**/
//https://leetcode.cn/problems/backspace-string-compare/solution/bi-jiao-han-tui-ge-de-zi-fu-chuan-by-leetcode-solu/
//https://leetcode.cn/problems/backspace-string-compare/solution/shuang-zhi-zhen-bi-jiao-han-tui-ge-de-zi-8fn8/
package string_code

//栈
func backspaceCompare(s, t string) bool {
	return build(s) == build(t)
}
func build(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}

//双指针
//思路：
//将介绍一种常量级空间复杂度的解法：双指针，并且比官方解思路更简单清晰！
//
//由于 # 号只会消除左边的一个字符，所以对右边的字符无影响，所以我们选择从后往前遍历 SS，TT 字符串。
//
//思路解析：
//
//准备两个指针 ii, jj 分别指向 SS，TT 的末位字符，再准备两个变量 skipSskipS，skipTskipT 来分别存放 SS，TT 字符串中的 # 数量。
//从后往前遍历 SS，所遇情况有三，如下所示：
//2.1 若当前字符是 #，则 skipSskipS 自增 11；
//2.2 若当前字符不是 #，且 skipSskipS 不为 00，则 skipSskipS 自减 11；
//2.3 若当前字符不是 #，且 skipSskipS 为 00，则代表当前字符不会被消除，我们可以用来和 TT 中的当前字符作比较。
//若对比过程出现 SS, TT 当前字符不匹配，则遍历结束，返回 falsefalse，若 SS，TT 都遍历结束，且都能一一匹配，则返回 truetrue。
//
//作者：demigodliu
//链接：https://leetcode.cn/problems/backspace-string-compare/solution/shuang-zhi-zhen-bi-jiao-han-tui-ge-de-zi-8fn8/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func BackspaceCompare(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
