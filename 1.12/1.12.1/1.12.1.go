/*
Inside the main function (you don't need to declare the function), 
you need to write a program:

In the first step, 10 positive integers are fed to standard input, 
which must be written in the order they were entered into an array of 10 elements. 
The type of numbers included in the array must correspond to the smallest possible unsigned integer. 
The name of the array that you must create yourself workArray (the condition is required). 
The fmt package has already been imported to read from standard input.

At the second stage, 3 more pairs of numbers are fed to the standard input - the indices of the elements of this array,
which need to be swapped (if such a pair of numbers is 3 and 7, then in the array the element with index 3 must be swapped with the element whose index is 7).
The elements of the resulting array should be output separated by a space on the standard output. 
Next, the used types will be checked automatically, the result of which will be added to your answer.

Using an array is a must!

Sample Input:
99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8
Sample Output:
99 137 151 187 117 71 20 187 93 67 type ok
*/
package main

import "fmt"

func main(){

var workArray [10]uint8

	for i := 0; i < 10; i++ {
		var a uint8
		fmt.Scan(&a)
		workArray[i] = a
	}
	for j := 0; j < 3; j++ {
		var i1, i2 uint8
		fmt.Scan(&i1, &i2)
		workArray[i1], workArray[i2] = workArray[i2], workArray[i1]
	}
	for _, value := range workArray {
		fmt.Print(value, " ")
	}
}