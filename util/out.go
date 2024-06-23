package util

import "fmt"

func Sout(arr ...any) {
	fmt.Print("[")
	for i, v := range arr {
		fmt.Print(v)
		if i != len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
	fmt.Println()
}
