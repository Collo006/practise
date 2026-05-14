package main

import "fmt"

func main() {
	i, j := 42, 2019
	fmt.Println(i, j)
	fmt.Println(&i, &j)
// p and q have stored the addresses of i and j
	p := &i
	q := &i
	fmt.Println(p)
	fmt.Println(q)
// *p and *q have  printed the results of i and j
	fmt.Println(*p)


	fmt.Println("%T\n",p)
// changes the value of i to 21 cause \\\8p indicates that it points to i
	*p =21
	fmt.Println(i)
	fmt.Println(*q)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}