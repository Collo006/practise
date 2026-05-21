package main

import "fmt"

func RepeatAlpha(s string) string {
	//create an empty variable
	result := ""
	//loop through the whole string
	for i := 0; i < len(s); i++ {
		char := s[i]
		if (char >= 'a' && char <= 'z') { // check if the character is a letter
			position := int(char - 'a' +1) // find the position of the character
			for j := 0; j < position; j++ {
				result += string(char)
			}
		} else if (char >= 'A' && char <= 'Z') {
           position := int(char - 'A' + 1)
		   for j :=0; j <position; j++ {
			result += string(char)
		   }
		} else { 
			 result += string(char) // if the character is not an alphabet return the character
		}
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}


/*Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc... */