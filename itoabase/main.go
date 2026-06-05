package main

import "fmt"

func Itoa(n int) string {
	//if the number is zero, we directly return "0"
	if n == 0 {
		return "0"
	}
	//then convert the -ve number to positive
	isNeg := false
	//if the number is negative, we set it to true and convert it back to a positive number
	if n < 0 {
		isNeg = true
		n = -n
	}
	//store the result in an empty string
	result := ""
	//if n is s number now convert it to a string by taking the last digit and converting it to a character, then adding it to the result string. We repeat this process until n becomes zero.
	for n > 0 {
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

//converts a number from an integer to a string. It handles negative numbers and zero as well. The function works by repeatedly taking the last digit of the number and converting it to a character, building the resulting string in reverse order.
