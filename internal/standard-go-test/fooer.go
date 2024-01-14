package main

import (
	"strconv"
)

// If the number is divisible by 3, write "Foo" otherwise, the number

func Fooer(input int) string {
	isFoo := (input % 3) == 0
	if isFoo {
		return "Foo"
	}
	return strconv.Itoa(input)
}
