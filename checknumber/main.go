package main

import "fmt"

func CheckNumber(arg string) bool {

	if arg >= "0" && arg <= "9" {
		return true
	} else {
		return false
	}
	
}

func main() {
	fmt.Println(CheckNumber("9"))
		fmt.Println(CheckNumber("a"))
}

 