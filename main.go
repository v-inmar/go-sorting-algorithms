package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	rand.Seed(time.Now().Unix())

	sorter := Sorter{}
	sorter.Values = rand.Perm(10000) // change number to see the difference in speed between the sorting algorithm
	fmt.Printf("Length: %d Unsorted Values: %v\n\n", len(sorter.Values), sorter.Values)


	selectionSortedValues := sorter.Selection()
	fmt.Printf("Selection Sort = Elapsed Time: %s Length: %d\n", sorter.ElapsedTime, len(selectionSortedValues))
	// Uncomment below to also print sorted values
	// fmt.Printf("Sorted Values: %v\n\n", selectionSortedValues)
	
	insertionSortedValues := sorter.Insertion()
	fmt.Printf("Insertion Sort = Elapsed Time: %s Length: %d\n", sorter.ElapsedTime, len(insertionSortedValues))
	// Uncomment below to also print sorted values
	// fmt.Printf("Sorted Values: %v\n\n", insertionSortedValues)


}