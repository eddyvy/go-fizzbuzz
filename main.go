package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputNumber := TakeInputNumber()
	FizzBuzzIt(inputNumber)
}

// Prints numbers 1 by 1 until the number passed (lastNum)
// Multiples of 3 and 5 are changes with "fizz", "buzz" or "fizzbuzz"
func FizzBuzzIt(lastNum uint64) {
	for i := uint64(1); i <= lastNum; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

// Asks for the user input
func TakeInputNumber() uint64 {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce the number for FizzBuzz: ")
	scanner.Scan()

	if scanner.Err() != nil {
		log.Fatal("An error occurred -> ", scanner.Err())
	}

	inputNumber, err := strconv.ParseUint(scanner.Text(), 10, 0)

	if err != nil {
		log.Fatal("An error occurred -> ", err, "\nPlease write an integer number!")
	}

	return inputNumber
}
