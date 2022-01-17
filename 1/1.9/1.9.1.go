/*
An integer is fed into the input.
If the number is positive - display the message "Number is positive",
if the number is negative - "Number is negative". If zero is supplied,
display the message "Zero". Display message without quotes.

Sample Input:
5

Sample Output:
Number positive
*/
package main

import "fmt"

func main() {

    var a int
    fmt.Scan(&a)

   if a > 0 {
        fmt.Println("Число положительное")
    }else if a < 0{
        fmt.Println("Число отрицательное")
    }else if a == 0 {
        fmt.Println("Ноль")
    }
}
