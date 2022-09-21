package main

import (
	"fmt"
	"time"
)

func main() {
	demo4()
}

func demo1() {
	fmt.Println("main start")
	go func() {
		fmt.Println("goroutine")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func demo2() {
	fmt.Println("main start")
	ch := make(chan string, 10)
	ch <- "a"
	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func demo3() {
	fmt.Println("main start")
	ch := make(chan string, 10)
	go func() {
		fmt.Println("producer start")
		ch <- "a"
		ch <- "b"
		ch <- "c"
		ch <- "d"
		fmt.Println("producer end")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func demo4() {
	fmt.Println("main start")
	ch := make(chan string, 10)
	go func() {
		fmt.Println("producer start")
		ch <- "a"
		ch <- "b"
		ch <- "c"
		ch <- "d"
		ch <- "e"
		ch <- "f"
		fmt.Println("producer end")
	}()
	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
