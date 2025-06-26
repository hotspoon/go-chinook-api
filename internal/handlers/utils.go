package handlers

import "fmt"

func ParseInt(s string) int {
	i := 0
	fmt.Sscan(s, &i)
	return i
}
