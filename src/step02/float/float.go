package main

import (
	"fmt"
)
//声明零值的过程中，必须指定变量的类型，
func main() {
//	var a float64
//	fmt.Println(a)
//	price:=0.0
//	fmt.Println(price)
pig:=0.0
for i:=0;i<11;i++{
	pig+=0.1
}
	fmt.Printf("%.1f",pig)
//	fmt.Printf("%v\n",third)
//	fmt.Printf("%f\n",third)
}