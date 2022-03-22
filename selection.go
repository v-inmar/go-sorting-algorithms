package main

func selection(s []int) []int{
	// create a new slice same length as argument s
	t := make([]int, len(s))

	// copy the s slice into t slice
	// :this is to make sure that GC cleans up s slice
	copy(t, s)

	// Loop through the slice
	for i := 0; i < len(t); i++{
		// get the current index (you can assume that this index holds the lowest value)
		minIndex := i
		
		// loop through the same slice again, but this time the initial index is the outside loop
		// current index plus 1
		for j := i+1; j < len(t); j++{
			// check that the value of the index minIndex is greater that the value of the j index
			if t[minIndex] > t[j]{
				// make j index the new minIndex
				minIndex = j
			}
		}

		// swap
		t[i], t[minIndex] = t[minIndex], t[i]
	}

	return t
}