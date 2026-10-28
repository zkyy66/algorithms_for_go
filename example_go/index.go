package example_go

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Exindex() {
	//taoyi
	x := fooTao()
	fmt.Println(*x)
	makeSlice()
	//taoyi end

	//defer
	// fmt.Println("defer value:", f())
	// a, b := 1, 2
	// f(a, b)
	// deferDemo()
	// defer end
	// x, y := func() (int, int) {
	// 	fmt.Println("func")
	// 	return 3, 4
	// }()
	// func(a, b int) {
	// 	fmt.Println("a*a+b*b=", a*a, b*b)
	// }(x, y)

	// func(x int) {
	// 	fmt.Println("x*x+y*y=", x*x+y*y)
	// }(y)

	// func() {
	// 	fmt.Println("x*x+y*y=", x*x+y*y)
	// }()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// // wg.Add(2)
	// // go exampleWaitGroup("hi", 10)
	// // go exampleWaitGroup("hello", 10)
	// // wg.Wait()
	// fmt.Println("defer fun\n")
	// defer fmt.Println("9")
	// fmt.Println("0")
	// defer fmt.Println("8")
	// fmt.Println("1")
	// if false {
	// 	defer fmt.Println("not reachable")
	// }
	// defer func() {
	// 	defer fmt.Println("7")
	// 	fmt.Println("3")
	// 	defer func() {
	// 		fmt.Println("5")
	// 		fmt.Println("6")
	// 	}()
	// 	fmt.Println("4")
	// }()
	// fmt.Println("2")
	// return
	// defer fmt.Println("not reachable")
}

func exampleWaitGroup(greet string, times int) {

	for i := 0; i < times; i++ {
		d := time.Second * 2
		time.Sleep(d)
	}
	wg.Done()
}

//	func mergeFile() error {
//		f, _ := os.Open("xxx.txt")
//		if f != nil {
//			defer func(f io.Closer) {
//				if err := f.Close(); err != nil {
//					fmt.Printf("%v\n", err)
//				}
//			}(f)
//		}
//		f, _ = os.Open("xxx.txt2")
//		if f != nil {
//			defer func(f io.Closer) {
//				if err := f.Close(); err != nil {
//					fmt.Printf("%v\n", err)
//				}
//			}(f)
//		}
//		return nil
//	}
// func f() (r int) {
// 	t := 5
// 	defer func() {
// 		t = t + 5
// 	}()
// 	return t
// }

func deferDemo() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")
	fmt.Println("env:", user)
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover suceess err:", err)

			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()
			if user == "" {
				panic("should set user env")
			}
			fmt.Println("after panic")
		}()
	}()
	time.Sleep(10)
	fmt.Println("end of main function")
}
func deferSum(a, b int) {
	c := a + b
	fmt.Println("sum:", c)
}

func f(a, b int) {
	defer deferSum(a, b)
	fmt.Println(a, b)
}
func exSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]
	s2 = append(s2, 100)
	s2 = append(s2, 200)
	s1[2] = 300
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
func makeSlice() {
	s := make([]int, 0)
	fmt.Println(s)
	s = append(s, 1)
	fmt.Println(s)
	s = append(s, []int{2, 3, 4}...)
	fmt.Println(s)
	s = append(s, 5, 6, 7)
	fmt.Println(s)
}
