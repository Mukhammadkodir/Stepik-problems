/*
An array of integers is given. The numbering of elements starts from 0. 
Write a program that displays the elements of an array whose indices are even (0, 2, 4 ...).

Input data
First, the number NNN is given - the number of elements in the array (1≤N≤1001 \ leq N \ leq 1001≤N≤100). 
Then NNN numbers are written separated by a space - the elements of the array. The array consists of integers.

Output
It is necessary to display all the elements of the array with even indices.

Sample Input:
6
4 5 3 4 2 3

Sample Output:
4 3 2
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
    for i := 0; i < N; i++{
        if i == 0 || i % 2 == 0{   
    	    fmt.Print(slice[i]," ")
        }
    }
}