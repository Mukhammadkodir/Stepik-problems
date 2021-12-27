/*
It is required to write a program that reads two natural numbers A and
B from the keyboard (each no more than 100, A <B).
Print the sum of all numbers from A to B inclusive.

Sample Input:
15
Sample Output:
15
*/

package main

import "fmt"

func main() {

	var a, b, count int
	fmt.Scanln(&a, &b)

	for ; a <= b; a++ {
		count += a
	}
	fmt.Println(count)
}
