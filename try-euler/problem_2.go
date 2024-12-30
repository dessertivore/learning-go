package main

import (
	"fmt"
)

var cache = map[int]int{}

func fib(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	if val, ok := cache[num]; ok {
		return val
	}
	cache[num] = fib(num-1) + fib(num-2)
	return cache[num]
}

func Problem2() {
	defer fmt.Println("The second Project Euler func is now complete...")
	var sum int
	var x int
	
	for fib(x) < 4000000 {
		if fib(x)%2 == 0 {
			{
				sum += fib(x)
			}
		
		}
		x++
	}
	fmt.Println(sum)}
