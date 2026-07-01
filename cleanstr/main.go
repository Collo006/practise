package main

import (
	"os"
	"fmt"
)

//we use "os" package to access the command-line arguments

func main() {
	//check for the command-line  arguments
	if len(os.Args) != 2 {
		fmt.Println('\n')
		return
	}
    //we store the first argument in str
	str := os.Args[1]
	//this is used to keep track of the first character
	i :=0
	//this used to check for the first word
	firstWord := true

	//keep moving throug the loop until we reach the end
	for i < len(str) {
		//skip spaces and tabs
		for i < len(str) && (str[i]== ' ' || str[i] == '\t'){
			i++
		}

		//End if nothing left
		if i >= len(str){
			break
		}
		//print one spacebefore every word except the first 
		if !firstWord {
			fmt.Print(' ')
		}
		firstWord = false

		//print current word
		for i < len(str) && (str[i] !=' ' && str[i] != '\t') {
			fmt.Print(rune(str[i]))
			i++
		}
	}
	fmt.Println('\n')
}