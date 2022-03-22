package main

import (
	"fmt"
	"time"
)


func insertion(s []int) []int{
	// create a new slice same length as argument s
	// copy the s slice into t slice
	// :this is to make sure that GC cleans up s slice
	t := make([]int, len(s))
	copy(t, s)

	// start timer
	start := time.Now()

	// interate over the slice using key/value range for loop
	for k, v := range t{

		// No need to use the very first entry in the loop
		// So we only start from index 1 and above
		if k > 0{

			// get the index of the previous entry
			// this will be use for shifting the values
			i := k-1

			// Go's "while loop"
			// Goes inside the loop if current i is bigger or equals to 0
			// AND
			// value of t[i] is bigger than current value (v)
			for i >= 0 && t[i] > v{

				// move the t[i] value one index over
				t[i+1] = t[i]

				// change i index to the next lower index
				i = i-1
			}

			// Insert the current value (v) infront of index i, so i+1 index
			// since t[i] is small than value v
			t[i+1] = v
		}
	}

	// get the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Sort time: %s\n", elapsed)
	
	return t
}