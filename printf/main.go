package main

import "fmt"

func PrintIfNot(str string) string {
	if str < "3" || str == ""{
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

// Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.

//    If it's an empty string return G followed by a newline \n.

// Write a function that takes a string and returns:

//    "G\n" if the string's length is less than 3 (including empty string).

//    "Invalid Input\n" otherwise.
