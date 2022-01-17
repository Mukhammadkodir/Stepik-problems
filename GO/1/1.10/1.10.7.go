/*
The deposit in the bank is x rubles. It increases annually by p percent,
after which the fractional part of kopecks is discarded.
Each year, the amount of the deposit becomes larger.
Determine how many years the contribution will be at least y rubles.

Input data
The program receives three natural numbers as input: x, p, y.

Output
The program should print one integer.

Sample Input:
one hundred
10
200

Sample Output:
eight
*/

package main

import "fmt"

func main() {
	
	var x, p, y, count int
	count = 0
	fmt.Scan(&x, &p, &y)

	for {
		percent := x * p / 100
		x += percent
		count++
		if x > y {
			break
		}
	}
	fmt.Println(count)
}
