package main

import (
	"fmt"
)

func main(){
	toSort := []int{345,2,756,723,34,12,45,67,678,567,234,894,904,324,4536,456}
	fmt.Printf("Length: %d Unsorted Values: %v\n", len(toSort), toSort)

	toSort = selection(toSort)
	fmt.Printf("Length: %d Unsorted Values: %v\n", len(toSort), toSort)
}