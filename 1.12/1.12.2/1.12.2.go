/*
Write a program that takes as input the number N (N≥4) N (N \ geq 4) N (N≥4), 
and then NNN integers-elements of the slice. The 4th (index 3) element of this slice must be sent to the output.

Sample Input:
5
41 -231 24 49 6

Sample Output:
49
*/

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)

	for i := 0; i < N; i++ {
		var number int
		fmt.Scan(&number)
		slice[i] = number
	}
	fmt.Println(slice[3])
}