package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case ch <- 1:
		fmt.Println("case")
	}
	fmt.Println("종료")
}
