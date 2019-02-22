package main

import (
	"fmt"
	"time"
)

func ch1(c1 chan int) {
	defer close(c1)
	for i := 1; i <= 10; i++ {
		c1 <- i
	}
}

func ch2(c1, c2 chan int) {
	for {
		if v, ok := <-c1; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	c2 <- 1
}

func main() {
	c1 := make(chan int, 2)
	c2 := make(chan int, 2)

	go ch1(c1)

	go ch2(c1, c2)

	//<-c2

	select {
	case <-c2:
		fmt.Println("已收到数据")
	case <-time.After(5 * time.Second):

	}

}
