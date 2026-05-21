package main

import (
	"fmt"
)

func HashCode(dec string) string {
	// store the results in an empty array
	result := ""
	
	for i := 0; i <  len(dec); i++ {
		brenda := dec[i]
		hash := (int(brenda) +  len(dec)) % 127 // hashcode

		//unprintable characters
		if hash < 33 {
			hash += 33
		}
		// return the result as a string
		result += string(hash)
	}
	return result
}





 func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}



/* Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

    The hash equation is computed as follows:

(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

    If the resulting character is unprintable add 33 to it.
 */
