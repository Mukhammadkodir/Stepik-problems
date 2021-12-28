/*
Write a program that in a sequence of numbers finds the sum of two-digit numbers divisible by 8.
The program in the first line receives as input the number n - the number of numbers in the sequence,
in the second line - n numbers included in the given sequence.

Sample Input:
5
38 24 800 8 16
Sample Output:
40
*/

package main

import "fmt"

func main() {

	var n, a, c int
	fmt.Scan(&n)
	a = 0

	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		if c%8 == 0 && c < 100 && c >= 10 {
			a += c
		}
	}
	fmt.Println(a)
}
