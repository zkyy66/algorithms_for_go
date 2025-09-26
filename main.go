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
	// testScan()
	fmt.Println(fruits())
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

type Fruit struct {
	Name  string
	Price float64
	Qty   int64
}

func fruits() float64 {
	// items := []map[string]interface{}{
	// 	{"name": "apple", "price": 3.5, "qty": 4},
	// 	{"name": "banana", "price": 2.0, "qty": 6},
	// 	{"name": "cherry", "price": 5.0, "qty": 2},
	// }

	var fruit []*Fruit
	fruit = append(fruit, &Fruit{
		Name:  "apple",
		Price: 3.5,
		Qty:   4,
	})
	fruit = append(fruit, &Fruit{
		Name:  "banana",
		Price: 2.0,
		Qty:   6,
	})
	fruit = append(fruit, &Fruit{
		Name:  "cherry",
		Price: 5.0,
		Qty:   2,
	})
	var res float64
	for _, v := range fruit {
		res = jisuanPrice(v.Price, v.Qty)
	}
	return res
	// for _, it := range items {
	//         if it["qty"].(int) > 3 {
	//                 t += it["price"].(float64) * float64(it["qty"].(int)) * 0.9
	//         } else {
	//                 t += it["price"].(float64) * float64(it["qty"].(int))
	//         }
	// }

	// fmt.Println("Total:", t)
}
func jisuanPrice(price float64, nums int64) float64 {
	var t float64 = 0
	if nums > 3 {
		t = price * float64(nums) * 0.9
	} else {
		t = price * float64(nums)
	}
	return t
}
