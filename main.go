package main

import (
	"fmt"
	"time"
)

var (
	flagClose chan bool   = make(chan bool, count)
	results   chan uint64 = make(chan uint64, tail*count)
	endtimer  chan bool   = make(chan bool)
	count     uint64      = 4
	tail      uint64      = 3
)

func fib(x uint64) uint64 {
	if x < 2 {
		return uint64(x)
	}
	return fib(x-1) + fib(x-2)
}

func fibGor(start, stop uint64) {
	go func(x, y uint64) {
		bufX, bufY := x, y
		for {
			if bufX != bufY {
				results <- bufX
				//results <- fib(bufX)
				bufX++
			}
		}
	}(start, stop)

	for {
		flag := <-flagClose
		if flag {
			return
		}
	}
}

func main() {
	var j uint64 = 5
	var i uint64
	for i = 1; i <= count; i++ {
		go fibGor(j, j+tail-1)
		j += tail
	}

	for i = 0; i < 10; i++ {
		if uint64(len(results)) == tail*count {
			fmt.Println("Все выполнено в срок")
			return
		}
		fmt.Println(len(results))
		time.Sleep(200 * time.Millisecond)
	}

	for i = 1; i <= count; i++ {
		flagClose <- true
	}
	fmt.Println("Срок вышел!")
}
