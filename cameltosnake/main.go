package main

import "fmt"

func CamelToSnake(s string) string {
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' ||  (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			return s
		}
		if i < len(s)-1 && s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		}
	}
	lastchar := s[len(s) - 1]
	if lastchar >= 'A' && lastchar <= 'Z' {
		return s
	}
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
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

    lowerCamelCase
    UpperCamelCase

Rules for writing in camelCase:

    The word does not end on a capitalized letter (CamelCasE).
    No two capitalized letters shall follow directly each other (CamelCAse).
    Numbers or punctuation are not allowed in the word anywhere (camelCase1).
 */

