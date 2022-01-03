/*
Input data
Three natural numbers a, b, c are given. Determine if there is a triangle with such sides.

Output
If the triangle exists, print the line "Exists", otherwise print the line "Doesn't exist".
Print the string without quotes.

Sample Input:
4 5 6

Sample Output:
Exists
*/

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
