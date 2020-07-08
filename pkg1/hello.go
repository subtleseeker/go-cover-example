package pkg1

import "fmt"

func Hello() string {
	x := 1
	fmt.Println("hello")
	if x == 0 {
		fmt.Println(x)
	}
	return "hello"
}
