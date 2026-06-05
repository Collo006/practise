package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNeg := false

	if n < 0 {
		isNeg = true
		n = -n
	}
	result := ""

	if n > 0 {
		result = string('0'+byte(n%10)) + result
		n = n / 10
	}

	if isNeg {
		result = "-" + result
	}
	return result
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
