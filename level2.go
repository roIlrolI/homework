package main

import (
	"fmt"
	"time"
)
var ch = make(chan map[int]int,400)
var myres = make(map[int] int, 20)
var myres1 = make(map[int] int, 20)
func factorial(n int){

	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
		myres[n] = res
		ch <- myres
	}
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	myres1=<- ch
	for i, v := range myres1 {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
	time.Sleep(time.Second)
}