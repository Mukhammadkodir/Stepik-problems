/*
It's time for tasks where you can apply the knowledge gained in practice.

Mandatory execution conditions: data from the standard input is read by the readTask () function,
which returns 3 values ​​of the empty interface type. This function uses the encoding / json, fmt, and os packages - do not remove them from the import. You will most likely need the "fmt" package, but you can use any other package to write to stdout without removing fmt.

So, you get 3 values ​​of type empty interface: if everything is successful, 
then the first 2 values ​​will be float64. The third value, ideally, will be a string, 
which can have the following values: "+", "-", "*", "/" (a certain mathematical operation). 
But such ideal cases will not always be, you can get values ​​of other types, 
and the string in the third value may not belong to one of the specified mathematical operations.

The result of executing the program should be as follows:

    in a normal situation, you should print the result of performing a mathematical operation to 
	standard output with an accuracy of 4 digits after the decimal point (fmt.Printf (%. 4f));
    if the first or second value is not of type float64, you should print an error message like: 
	value = received_value: value_type (for example: value = true: bool)
    if the third value is of the wrong type or a sign is passed that is not related to the above mathematical operations, 
	the error message should be of the form: unknown operation

It is guaranteed that there can be only one error in the arguments, 
so if you saw that it contains an error when checking the first value, 
print an error message and exit the program, checking the rest of the arguments no longer matters, 
and the checking system will accept 2 error messages as a violation of the condition for completing the assignment.
*/
package main

import (
	"encoding/json" // the package is used to validate the response, do not delete it
	"fmt"           // the package is used to validate the response, do not delete it
	"os"            // the package is used to validate the response, do not delete it
)


 // raw data is obtained using this function
 // all received values are of the empty interface type



func readTask()(value1, value2, operation interface{}) {
	return 5.0, 5.6, "/"
}

func printError(value interface{}) {
	fmt.Printf("value=%v: %T\n", value, value)
}

func main() {
	value1, value2, operation := readTask()
	v1, ok := value1.(float64)
	if !ok {
		printError(value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		printError(value2)
		return
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			result := v1 + v2
			fmt.Printf("%.4f\n", result)
		case "-":
			result := v1 - v2
			fmt.Printf("%.4f\n", result)
		case "*":
			result := v1 * v2
			fmt.Printf("%.4f\n", result)
		case "/":
			result := v1 / v2
			fmt.Printf("%.4f\n", result)
		default:
			fmt.Println("unknown operation")
			return
		}
	} else {
		fmt.Println("unknown operation")
		return
	}
}

