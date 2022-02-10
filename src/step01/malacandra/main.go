package main

import "fmt"

func main() {
	const (
		road       = 56000000
		hourpreday = 24
	)
	fmt.Println(road / hourpreday / 28)

}
