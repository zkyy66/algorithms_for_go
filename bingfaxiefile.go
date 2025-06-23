package main

import (
	"fmt"
	"os"
	"sync"
)

func fileDemo() {
	f, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	ch := make(chan string)
	done := make(chan struct{})
	var wg sync.WaitGroup
	go writer(f, ch, done)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	wg.Wait()
	close(ch)
	<-done
}
func worker(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fmt.Sprintf("worker %d starting\n", id)
}
func writer(file *os.File, ch <-chan string, done chan<- struct{}) {
	for line := range ch {
		fmt.Println(line)
	}
	done <- struct{}{}
}
