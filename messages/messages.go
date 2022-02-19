package messages

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("hello, %v\n", name)
}

//an unpublished function
func depart(name string) string {
	return fmt.Sprintf("bye now, %v", name)
}
