package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 1:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}
	fmt.Println("종료")
}
