package main

import (
	"fmt"
	"log"
	"strconv"
)

func main(){
	var start, end string
	fmt.Scan(&start)
	fmt.Scan(&end)
	s, e := strconv.Atoi(start)
	finish, e :=strconv.Atoi(end)
		if e!=nil{
			log.Fatal(e)
		}
	sum :=0
	for i:=s; i<=finish; i++{
		sum = sum + i
	}
	fmt.Println(sum)
}