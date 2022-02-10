package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	sum := 0.0
	for sum <= 20.0 {
		switch save := rand.Intn(3) + 1; save {
		case 1:
			sum += 0.05
		case 2:
			sum += 0.10
		case 3:
			sum += 0.25
		}
	}
	fmt.Println(sum)
	fmt.Printf("%5.2f", sum)
}
