package main

import "time"

type Sorter struct {
	ElapsedTime time.Duration
	Values []int
}

func (s *Sorter) Selection()[]int{
	// Make a copy of the Values slice
	// so we can preserved it for other sorting algorithm
	t := make([]int, len(s.Values))
	copy(t, s.Values)
	start := time.Now()
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
	s.ElapsedTime = time.Since(start)
	return t
}

func (s *Sorter)Insertion()[]int{
	// Make a copy of the Values slice
	// so we can preserved it for other sorting algorithm
	t := make([]int, len(s.Values))
	copy(t, s.Values)

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
	s.ElapsedTime = time.Since(start)
	return t
}