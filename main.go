package main

import (
	"algorithms_for_go/array_code"
	"algorithms_for_go/string_code"
	"fmt"
)

func main() {
	//fmt.Println("二分查找\n")
	//binary_list := []int{1, 2, 3, 4, 6, 9, 10}
	//binary_target := 11
	//res := array_code.BinarySearch(binary_list, binary_target)
	//fmt.Printf("结果：%d\n", res)
	//fmt.Println("二分查找-左闭右开\n")
	//resTwo := array_code.BinarySearchVersionTwo(binary_list, binary_target)
	//fmt.Printf("结果：%d\n", resTwo)
	//fmt.Println("35:搜索插入位置")
	//fmt.Printf("结果：%d\n", array_code.SearchInsertion(binary_list, binary_target))

	fmt.Println("26\n")
	nums := []int{1, 2, 2}
	fmt.Println(array_code.RemoveDuplicates(nums))
	fmt.Println(array_code.RemoveDuTwo(nums))
	a := []int{0, 1, 2, 0, 4, 5, 6}
	fmt.Println(array_code.MoveZerosTwo(a))
	fmt.Println("844##########")
	s := "ab#c"
	t := "ab#c"
	fmt.Println(string_code.BackspaceCompare(s, t))
}
