/*
run test

  go test
  go test -bench .
  go test -bench FizzBuzz

  go test -bench . -benchtime=10s -count=5
  go test -bench . -benchmem

  go test -bench FizzBuzzLarge -benchtime=10s
  go test -bench FizzBuzz -benchmem

*/

package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}

	got := FizzBuzz(15)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("FizzBuzz(15) = %v; want %v", got, expected)
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(1000)
	}
}

func BenchmarkFizzBuzzLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(10_000)
	}
}
