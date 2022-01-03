/*
You are given a sequence of integers. Write a program that counts the number of positive numbers among the elements of a sequence.

Input data
First, the number NNN is given - the number of elements in the sequence (1≤N≤1001 \ leq N \ leq1001≤N≤100). Then NNN numbers are written separated by a space - the elements of the sequence. The sequence consists of integers.

Output
You must print a single number - the number of positive elements in the sequence.

Sample Input:
5
1 2 3 -1 -4

Sample Output:
3
*/

package main
import "fmt"

func main() {
    var N, count int
	fmt.Scan(&N)


	for i := 0; i < N; i++ {
		var number int
		fmt.Scan(&number)
        if number >= 0{
            count++
        }
	}
    fmt.Println(count)
   
}