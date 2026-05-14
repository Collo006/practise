package main

import "fmt"

func RepeatAlpha(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		brenda := s[i]
		if (brenda >= 'a' && brenda <= 'z') {
			position := int(brenda - 'a' +1)
			for j := 0; j < position; j++ {
				result += string(brenda)
			}
		} else if (brenda >= 'A' && brenda <= 'Z') {
           position := int(brenda - 'A' + 1)
		   for j :=0; j <position; j++ {
			result += string(brenda)
		   }
		} else {
			 result += string(brenda)
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