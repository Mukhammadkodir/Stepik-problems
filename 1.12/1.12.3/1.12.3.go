/*
Five integers are fed into the input, which are written to the array.
However, this part of the program has already been written.
You need to write a piece of code with which you can find and display the maximum number in this array.

Sample Input:
2
3
56
45
21

Sample Output:
56
*/

package main

import "fmt"

func main() {
	array := [5]int{}
	var a int

	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	max := array[0]
	for j,_ := range array {
		if array[j] > max {
			max = array[j]
		}
	}
	fmt.Println(max)
}
