package main

import "fmt"

func sum(numbers ...int){
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println("Total:", total)
}

func main(){
	sum(1, 2)
	sum(10,20,30)
	sum(5,5,5,5)
}