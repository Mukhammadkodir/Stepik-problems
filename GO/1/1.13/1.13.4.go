/*
Three numbers are given - a, b, c (a <b <c) a, b, c (a <b <c) a, b, c (a <b <c) - the lengths of the sides of the triangle.
You need to check if the triangle is right-angled. If so, output "Rectangular". Otherwise display "Non-rectangular"

Sample Input:
6 8 10

Sample Output:
Rectangular
*/

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
