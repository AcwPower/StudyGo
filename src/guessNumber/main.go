package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num = 25
	rand.Seed(time.Now().Unix())
	var num1 = rand.Intn(100) + 1
	var count = 0
	for num1 != num {
		fmt.Println(num1)
		num1 = rand.Intn(100 + 1)
		count++
		fmt.Println(count)
	}
	fmt.Println(num1, count)
}
