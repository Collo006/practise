package main

import "fmt"

func CamelToSnake(s string) string {
	//first start with checking the length of the string
	if len(s) == 0 {
		return ""
	}
	//run a for loop to check if the string contains two consecutive capital letters or contains any lowercase letters
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' ||  (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			return s
		}
		// start checking from the last character as you come forward
		if i < len(s)-1 && s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		}
	}
	//checking if the last character is uppercase
	lastchar := s[len(s) - 1]
	if lastchar >= 'A' && lastchar <= 'Z' {
		return s
	}
	//function to change the camel case to snake case
	result := "" // create an empty varibale to store the result
	for i := 0; i < len(s); i++ { // then loop through the string
		if s[i] >= 'A' && s[i] <= 'Z' { // check if the string contains capital letters
			if i != 0 { // if there is no capital letters at index 0 then put an underscore '_'
				result += "_" // return the result '_'
			}
			result += string(s[i] + 32) //if it is capital convert to lowercase
		} else {
			result += string(s[i]) // else return the string as it is
		}
	}
	return result
}

func main() {
	fmt.Println(CamelToSnake("camelCase"))
	fmt.Println(CamelToSnake("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnake("HelloWorld"))
	fmt.Println(CamelToSnake("helloWorld"))||
	fmt.Println(CamelToSnake("camelToSnakeCase"))
	fmt.Println(CamelToSnake("hey2"))
}
/* Write a function that converts a string from camelCase to snake_case.

    If the string is empty, return an empty string.
    If the string is not camelCase, return the string unchanged.
    If the string is camelCase, return the snake_case version of the string.

For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:
acquired
    lowerCamelCase
    UpperCamelCase

Rules for writing in camelCase:

    The word does not end on a capitalized letter (CamelCasE).
    No two capitalized letters shall follow directly each other (CamelCAse).
    Numbers or punctuation are not allowed in the word anywhere (camelCase1).
 */

