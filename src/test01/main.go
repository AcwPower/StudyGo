package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num:=rand.Intn(3)+1
	if num==1{
		fmt.Println("fcuk 1")
	}else if num==2 {
		fmt.Println("fuck 2")
	}else {
		fmt.Println("fuck 3")
	}
	fmt.Println(num)
}
