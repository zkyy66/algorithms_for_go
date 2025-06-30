package main

import (
	"fmt"
)

//一群猴子排成一圈，按1，2，...，n依次编号。然后从第1只开始数，数到第m只,把它踢出圈，从它后面再开始数，再数到第m只，在把它踢出去...，
// 如此不停的进行下去，直 到最后只剩下一只猴子为止，那只猴子就叫做大王。要求编程模拟此过程，输入m、n,输出最后那个大王的编号
func monkeyKing(n, m int) int {
	monkeys := make([]int, n)
	for i := 0; i < n; i++ {
		monkeys[i] = i + 1
	}
	
	index := 0
	for len(monkeys) > 1 {
		index = (index + m - 1) % len(monkeys)
		monkeys = append(monkeys[:index], monkeys[index+1:]...)
	}
	return monkeys[0]
}

func main() {
	var n, m int
	fmt.Print("请输入猴子总数n 和 数到第几个淘汰m（以空格分隔）: ")
	fmt.Scan(&n, &m)
	king := monkeyKing(n, m)
	fmt.Printf("最后的大王是编号为：%d 的猴子\n", king)
}
