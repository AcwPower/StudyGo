package main

import (
	"fmt"
	"math/rand"
	"time"
)
const distance=62100000
var hourperday=24
var seconds=3600

func calculate(){
	company := make([]string, 0)
	company = append(company, "Space Adventures",
		"SpaceX",
		"Virgin Galacic")
	companyname:=company[rand.Intn(len(company))]
	returncode:=make([]string,0)
	returncode=append(returncode, "往返",
	  "单程")											
	isreturncode:=returncode[rand.Intn(len(returncode))]
	speed :=rand.Intn(15)+16
	oncecost :=speed+20
	days:=distance/(hourperday*seconds*speed)
	switch isreturncode {
	case "往返":
		fmt.Printf("%-19v %-10v %-7v %+4v\n",companyname,days,isreturncode,oncecost*2)
	default:
		fmt.Printf("%-19v %-10v %-7v %+4v\n",companyname,days,isreturncode,oncecost)
	}

}
func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("太空航行公司","\t\t","飞行天数","\t","飞行类型","\t","价格(百万美元)")
	for count:=0;count<10;count++{
		calculate()
	}
}