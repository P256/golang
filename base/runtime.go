package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
		runtime.Gosched()
	}
}

func word() {
	for i := 0; i < 10; i++ {
		fmt.Println("word")
		runtime.Gosched()
	}
}

func main() {

	go hello()

	go word()

	time.Sleep(time.Second * 1)

	fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())

}
