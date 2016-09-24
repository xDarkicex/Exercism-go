package hamming

import "errors"

const testVersion = 4

//Distance calculates DNA string distance
func Distance(a, b string) (dis int, err error) {
	if len(a) != len(b) {

		return -1, errors.New("Strands must be the same length")
	}
	diff := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff = diff + 1
		}
	}
	return diff, nil
}
