package Daily13

func FindMajorElem(elems []int) int {
	// return first element if length of slice is only 1
	if len(elems) == 1 {
		return elems[0]
	}

	// create map to store elements and their count
	mapElems := make(map[int]int)

	for _, elem := range elems {
		// check if element exists in map
		if _, ok := mapElems[elem]; ok {
			// increment count
			mapElems[elem]++
			// check if element is already major element
			if mapElems[elem] > len(elems)/2 {
				return elem
			}
		} else {
			// add element to map
			mapElems[elem] = 1
		}
	}

	// return zero in case of empty slice or errors
	return 0
}
