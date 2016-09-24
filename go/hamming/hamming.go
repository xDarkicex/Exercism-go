package hamming

import "errors"

const testVersion = 4

//Distance calculates DNA string distance
func Distance(a, b string) (dis int, err error) {
	if len(a) != len(b) {

		return -1, errors.New("Strands must be the same length")
	}
	//set difference too zero
	diff := 0
	//must build loop to loop through DNA strands
	for i := 0; i < len(a); i++ {
		//if slice position value is not the same in both slices add 1 to diff
		if a[i] != b[i] {
			diff = diff + 1
		}
	}
	//once i = len(a) return the int value of diff coounting how many differnce in each string.
	//or if none found return nil
	return diff, nil
}
