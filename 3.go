package main

import (
	"fmt"
	"log"
	"strconv"
)

func main(){
	sumi := 0
	var start, end string
	fmt.Scan(&start)
	fmt.Scan(&end)
	s, e := strconv.Atoi(start)
	f, e := strconv.Atoi(end)
	if e!=nil{
		log.Fatal(e)
	}
	for i :=s; i<=f; i++{
		if i % 2==1{
		sumi = sumi+i
		}
	}
	fmt.Println(sumi)

}

