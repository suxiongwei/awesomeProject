package main

import (
	"fmt"
	"time"
)

func main4()  {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 100
	}()

	go func() {
		ch2 <- 200
	}()

	select {
	case data := <-ch1:
		fmt.Println("ch1读取了数据:", data)
	case data := <-ch2:
		fmt.Println("ch2读取了数据:", data)
	default:
		fmt.Println("执行了default")
	}


}
