package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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


func main() {
	//  Defer the print statement until the end of the func
	defer fmt.Println("The func is now complete...")
	//  Extract int from sys args and run getMultiples
	num := os.Args[1]
	log.Println("Finding all multiples of 3 and 5 under "+num)
	parsedNum, err := strconv.Atoi(num)
	if err != nil {
		panic(err)}
	sum := getMultiples(parsedNum)
	println(sum)
}