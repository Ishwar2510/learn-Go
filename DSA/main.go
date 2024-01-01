// effective go
// statically typed language
package main

import (
	"errors"
	"fmt"
)

func main() {
	// forLoop();
	ifElseWithError(10)
}

// if statement

func ifElseWithError(num int) {
	if num, err := test(num); err != nil {
		fmt.Println("errr", err)
	} else {
		fmt.Println("ans is s", num)
	}

}
func ifexam(num float64) {
	if num > 10 {
		fmt.Println("gtr than 10")
		return
	}
	fmt.Println("less than 10")
}

func forLoop() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func test(num int) (int, error) {
	if num >= 10 {
		return 10, nil
	}
	return -1, errors.New("number greter than 10")
}
