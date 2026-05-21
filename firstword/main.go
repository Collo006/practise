package main

import "fmt"

func firstWord(s string) string {
// first check the length of the string
if len(s) == 0 {
	return ""
}
// we then initialize "i"
i := len(s)-1; 
for i > 0 && s[i] != ' ' {
	i--
}
return s[i:]
}

func main(){
	fmt.Println(firstWord("Hello World"))
	fmt.Println(firstWord("come home"))

}

/*Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

    A word is a sequence of characters delimited by spaces or by the start/end of the argument.
*/