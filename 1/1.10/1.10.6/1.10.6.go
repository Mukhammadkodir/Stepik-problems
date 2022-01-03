/*
Write a program that reads integers from the console, one number per line.

For each entered number, check:
     if the number is less than 10, then we skip this number;
     if the number is more than 100, then stop reading the numbers;
     otherwise, print that number back to the console on a separate line.

Sample Input:
thirty
eleven
7
101

Sample Output:
thirty
eleven
*/

package main

import "fmt"

func main() {

	var x int

	for {
		fmt.Scan(&x)
		if x < 10 {
			continue
		} else if x > 100 {
			break
		} else {
			fmt.Println(x)
		}
	}
	
}
