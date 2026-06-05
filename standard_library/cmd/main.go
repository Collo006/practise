package main

import (
	"fmt"
	"os"
)

func main() {
	//we can use "os" to create a file
	// we first store it in a variable to access it
	file, err := os.Create("text.txt")

	if err != nil {
		fmt.Println("good job", err)
		//"Exit is used to end the current code with the status code,[0] means it is a success"
		os.Exit(1)
	}
	//closes the file rendering it unusable
	defer file.Close()

	data := "text\nto\nfile\ttabbing"

	//this line is used to write data to the created file
	file.WriteString(data)

	file_data, _ := os.ReadFile(file.Name())

	fmt.Println(string(file_data))

	//to remove a file use os.Remove("text.txt")

	//used to
	args := os.Args
	fmt.Println(args)

	hostname, _ := os.Hostname()
	fmt.Println(hostname)
}
