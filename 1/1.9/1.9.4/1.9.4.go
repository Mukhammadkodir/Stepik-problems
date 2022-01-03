/*
Determine if the ticket is lucky.
A ticket is considered lucky if in its six-digit number the sum of
the first three digits coincides with the sum of the last three.

Input data format
The ticket number is entered at the entrance - one six-digit number.
Output data format
Print "YES" if the ticket is lucky, otherwise - "NO".

Sample Input:
613244
Sample Output:
YES
*/

package main

import "fmt"

func main(){

	var g int
	fmt.Scanln(&g)
    var count, count1 int

    for i := 0; i <= 6; i++{
        if g >= 10{
            if i >= 3{
                count1 += g % 10
                g /= 10
            }else{
                count += g % 10
                g /= 10
            }
        }else{
            count1 += g
            break
        }
    }

    if count == count1{
	    fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}

