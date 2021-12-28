/*
The input is a number of type float64. You need to output the converted number according to the rule: the number is squared,
then the fractional part is cut off so that 4 decimal places remain. But if the number is equal to or less than zero - output:
"number R does not match", where R is the original number with 2 digits after the decimal point and 2 in width.
And if the number is more than 10,000 - display the original number in exponential notation (lowercase exponent sign).

Sample input:
-000012.2123

Sample output:
the number -12.21 is not suitable

Sample input:
1000001

Sample output:
1.000001e + 06

Sample Input:
12.12345678

Sample Output:
146.9782
*/

package main

import "fmt"

func main() {

	var a float64
	fmt.Scanln(&a)

	if a <= 0 {
		fmt.Printf("число %.2f не подходит\n", a)
	} else if a > 10000 {
		fmt.Printf("%e\n", a)
	} else if a > 0 {
		fmt.Printf("%.4f\n", a*a)
	}

}
