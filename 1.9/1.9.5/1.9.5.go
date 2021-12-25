/*
It is required to determine whether a given year is a leap year, recall:
A year is a leap year if it meets at least one of the following conditions:
- multiple of 400;
- a multiple of 4, but not a multiple of 100.

Input data
A single number is entered - the number of the year (integer, positive, does not exceed 10000).

Output
You want to output the word YES if the year is a leap year and NO otherwise.

Sample Input:
2000
Sample Output:
YES
*/

package main

import "fmt"

func main(){
	var g int
	fmt.Scanln(&g)
    
    if g % 400 == 0 || g % 4 == 0 && g % 100 != 0{
    	    fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}
