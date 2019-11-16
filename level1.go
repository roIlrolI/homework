package main

import (
	"fmt"
	"time"
)
func main() {
	var a, b, c int
	b = 1
	c=0
	slice := make([]string, 10)
	datetime := "2015-01-01 00:00:00"
	timeLayout := "2006-01-02 15:04:05"
	fmt.Println("input:")
	for {


	fmt.Scanf("%d", &a)
	if a == 0 {
		println(slice[c])
		c=c+1
	} else
	{
		for {
			datetime = time.Unix(int64(a), 0).Format(timeLayout)
			str := datetime
			slice[b] = str
			b = b + 1
			println("input ok!")
			break
		}
	}
}
}
