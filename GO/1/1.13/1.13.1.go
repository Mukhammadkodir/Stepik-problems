package main
import "fmt"

func main() {
    var N, count int
	fmt.Scan(&N)
	var number int
	

    for 0 < N{
		number = 0
        if N < 10{
            count += N
			break
        }
		number = N % 10
		N /= 10
		count += number
		
	}
    fmt.Println(count)


}
