/*
Let's put your knowledge of structures, methods and interfaces into practice and implement an object that satisfies the fmt.
Stringer interface. Let's call it "Battery".

First, you must declare a new type that conforms to the fmt.Stringer interface.

Your type should assume that it will look like this on print: [XXXX]: 
where spaces are "empty" battery capacity and X is "charged".

Secondly, on standard input you get a string consisting of exactly 10 digits: 
0 or 1 (the order of 0/1 is random). Your task is to read this string in any possible way and create an 
object of the type you declared at the first stage based on this string: 
I hope you understand that the string symbolizes the capacity of the battery: 0 is the "empty" part, and 1 is the "charged" part.

Third, the object you create must be named batteryForTest (use of this name is required).

You have virtually the entire file at your disposal, BUT the closing curly brace of main () is not visible to you, 
but it is present. There is a function (which you can't see either) in front of this bracket that takes an object of
 type fmt.Stringer - batteryForTest as an argument and directs it to standard output, so you don't need to print anything yourself.
*/


package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	charge string
}

func (b *Battery) String() string {
	zeros := strings.Count(b.charge, "0")
	ones := strings.Count(b.charge, "1")
	result := "[" + strings.Repeat(" ", zeros) + strings.Repeat("X", ones) + "]"
	return result
}

func main() {
	var s string
	fmt.Scan(&s)
	batteryForTest := &Battery{charge: s}
	fmt.Println(batteryForTest)
}
