package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for count:=0;count<10;count++{
		year := rand.Intn(100)+2001
		leap := year%400==0||(year%100!=0&&year%4==0)
		month := rand.Intn(12)+1
		//month:=2
		daypremonth :=31
		switch month {
		case 1,3,5,7,8,10,12:
			daypremonth = 31
		case 4,6,9,11:
			daypremonth=30
		default:
			if leap {
				daypremonth=29
			}else{
				daypremonth=28
			}
		}
		fmt.Println(year,month,daypremonth)
	}


}