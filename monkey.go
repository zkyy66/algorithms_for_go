package main

import (
	"fmt"
)

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
