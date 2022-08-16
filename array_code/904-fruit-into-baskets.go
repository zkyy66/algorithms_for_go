/**
 * @Date 2022/8/16
 * @Name 904-fruit-into-baskets
 * @VariableName
**/
////根据题意写判断函数
//func check(n Type) {/*statement*/}
////形参列表要根据题意变换
//func slidingWindow(nums []int) {
//    n := len(nums)
//    //使用i指针遍历整个数组
//    for i, j := 0, 0; i < n; i++ {
//        //调整j指针使[i, j]符合题意
//        for j <= i && check() {
//            /*statement*/
//            j++
//        }
//    }
//}
//
//作者：ballache
//链接：https://leetcode.cn/problems/fruit-into-baskets/solution/goyu-yan-ti-jie-fu-dai-hua-dong-chuang-k-l264/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
package array_code

func totalFruit(fruits []int) int {
	hashMap := map[int]int{}
	i, j, res, n := 0, 0, 0, len(fruits)
	for ; i < n; i++ {
		hashMap[fruits[i]]++
		for j < i && len(hashMap) == 3 {
			hashMap[fruits[j]]--
			if hashMap[fruits[j]] == 0 {
				delete(hashMap, fruits[j])
			}
			j++
		}
		if res < i-j+1 {
			res = i - j + 1
		}
	}
	return res
}

//php :https://leetcode.cn/problems/fruit-into-baskets/solution/hua-dong-chuang-kou-by-wuzhh-2kur/
