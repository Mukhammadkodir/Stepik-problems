/*
The sequence consists of natural numbers and ends with the number 0.
Determine the number of elements of this sequence that are equal to its largest element.

Input data format
A non-empty sequence of natural numbers is introduced, ending with the number 0
(the number 0 itself is not included in the sequence, but serves as a sign of its end).

Output data format
Print the answer to the problem.
*/

package main

import "fmt"

func main() {
	var x, count, max int
	max = 0

	for fmt.Scan(&x); x != 0; fmt.Scan(&x) {
		if x > max {
			max = x
			count = 1
		} else if x == max {
			count++
		}
	}
	fmt.Println(count)
}
