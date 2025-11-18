package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			runtime.Gosched()
		}
	}()
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("old goroutine: i = %d\n", i)
			if i == 5 {
				break
			}
		}
	}()
	i := 0
	for {
		runtime.Gosched()
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		if i == 3 {
			break
		}
	}
}
