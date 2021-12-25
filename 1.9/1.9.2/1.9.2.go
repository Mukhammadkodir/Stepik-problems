/*
For this three-digit number, determine if all of its numbers are different.
Input data format
The input is one natural three-digit number.
Output data format
Print "YES" if all digits of the number are different, otherwise - "NO".

Sample Input 1:
237
Sample Output 1:
YES

Sample Input 2:
117
Sample Output 2:
NO
*/
package main

import "fmt"

func main() {
var j int

    fmt.Scan(&j)
    one := j % 10
    j /= 10
    two := j % 10
    j /= 10


    if one != two && one != j {

        if two != j {
            fmt.Print("YES")
        }else{
            fmt.Print("NO")
        }
        
    }else{
       fmt.Print("NO")
        }     
}





