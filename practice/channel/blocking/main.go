package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 2
	fmt.Println("종료")
}
