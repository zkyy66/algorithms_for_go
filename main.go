package main

import (
	"fmt"
)

func testScan() {
	a := 0
	b := 0
	n, _ := fmt.Scan(&a, &b)
	for {
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
func main() {
	testScan()
	//fmt.Println("二分查找\n")
	// binary_list := []int{1}
	// binary_target := 1
	// //res := array_code.BinarySearch(binary_list, binary_target)
	// //fmt.Printf("结果：%d\n", res)
	// //fmt.Println("二分查找-左闭右开\n")
	// //resTwo := array_code.BinarySearchVersionTwo(binary_list, binary_target)
	// fmt.Println("二分查找-Demo\n")
	// resTwo := array_code.BinarySearchVersionTwo704(binary_list, binary_target)
	// fmt.Printf("结果：%d\n", resTwo)
	// //fmt.Println("35:搜索插入位置")
	// //fmt.Printf("结果：%d\n", array_code.SearchInsertion(binary_list, binary_target))

	// fmt.Println("26去重")
	// //nums := []int{1, 2, 2, 3, 3, 4, 4} //1,2,3,4
	// ////fmt.Println("One:", array_code.RemoveDuplicates(nums))
	// //fmt.Println("Two:", array_code.RemoveDuTwo(nums))
	// ////num, arrRes := array_code.RemoveDuThree(nums)
	// ////fmt.Println("Three:", num, arrRes)
	// //i, arra := array_code.Kuaimanzhizhen(nums)
	// //fmt.Println("Kuaimanzhizhen:", i, arra)

	// //a := []int{0, 1, 2, 0, 4, 5, 6}
	// //fmt.Println(array_code.MoveZerosTwo(a))
	// //fmt.Println("844##########")
	// //s := "ab#c"
	// //t := "ab#c"
	// //fmt.Println(string_code.BackspaceCompare(s, t))
	// //fmt.Println("977")
	// //
	// //fmt.Println(array_code.SortedSquares([]int{-4, -1, 0, 3, 10}))

	// fmt.Println("好未来算法题：反转单词")
	// howStr := "how are you ?"
	// res := string_code.ReverseStringWord(howStr)
	// fmt.Println(res)
}
