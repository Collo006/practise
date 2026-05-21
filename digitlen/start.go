package main

import "fmt"

func DigitLen(n, base int) int {
	//check if the base number is less than or greater than 36
	if base < 2 || base > 36 {
		return -1
	}
	// convert a negative number to a positive number
	if n < 0 {
		n = -n
	}
	//start the count at index 0
	count := 0
	for n > 0 {
		n = n/base
		count ++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
/* Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero.

    The second int must be between 2 and 36. If not, the function returns -1.
    If the first int is negative, reverse the sign and count the digits.

*/