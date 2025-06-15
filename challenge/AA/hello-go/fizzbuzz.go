package main

import (
	"strconv"
)

func FizzBuzz(n int) []string {
	result := make([]string, n)

	var number = 1
	for i := 0; i < n; i++ {
		if number%3 == 0 && number%5 == 0 {
			result[i] = "FizzBuzz"
		} else if number%3 == 0 {
			result[i] = "Fizz"
		} else if number%5 == 0 {
			result[i] = "Buzz"
		} else {
			result[i] = strconv.Itoa(number)
		}
		number++
	}

	return result
}
