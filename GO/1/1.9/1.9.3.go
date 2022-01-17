/*
A non-negative integer is given. Find and print the first digit of the number.
Input data format
The input is a natural number not exceeding 10,000.
Output data format
Print one integer - the first digit of the given number.

Sample Input:
1234
Sample Output:
one
*/

package main

import "fmt"

func main(){

	var g int
	fmt.Scanln(&g)

    for g > 0{
        if g >= 10{
            g /= 10
        }else{
            break
        }
    }
	fmt.Println(g)
}
