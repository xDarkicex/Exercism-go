package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

//IsLeapYear is a function
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	return year%4 == 0
}
