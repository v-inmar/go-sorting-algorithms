package main

import (
	"fmt"
)

func main(){

	sorter := Sorter{}
	sorter.Values = []int{345,2,756,723,34,12,45,67,678,567,234,894,904,324,4536,456}
	fmt.Printf("Length: %d Unsorted Values: %v\n", len(sorter.Values), sorter.Values)


	selectionSortedValues := sorter.Selection()
	fmt.Printf("Selection Sort = Elapsed Time: %s Length: %d Sorted Values: %v\n", sorter.ElapsedTime, len(selectionSortedValues), selectionSortedValues)
	insertionSortedValues := sorter.Insertion()
	fmt.Printf("Insertion Sort = Elapsed Time: %s Length: %d Sorted Values: %v\n", sorter.ElapsedTime, len(insertionSortedValues), insertionSortedValues)


}