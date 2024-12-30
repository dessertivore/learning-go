package main

import (
	"fmt"
	"log"
)


func getMultiples(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}


func Problem1(num int) {
	//  Defer the print statement until the end of the func
	defer fmt.Println("The first Project Euler func is now complete...")
	log.Println("Finding all multiples of 3 and 5 under ",num)
	sum := getMultiples(num)
	println(sum)
}