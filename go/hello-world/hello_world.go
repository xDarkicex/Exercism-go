package hello

import "unicode/utf8"

const testVersion = 2

//HelloWorld Prints hello world
func HelloWorld(name string) string {

	if utf8.RuneCountInString(name) > 1 {
		return "Hello, " + name + "!"
	}
	return "Hello, World!"
}
