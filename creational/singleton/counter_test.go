package singleton

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := GetCounter()

	fmt.Println("Initiate counter :", counter)

	counter.Add()
	fmt.Println("After counter.Add() :", counter)

	counter2 := GetCounter()
	fmt.Println("Initiate counter2 :", counter2)

	counter2.Add()
	fmt.Println("After counter2.Add() :", counter)

	fmt.Println("counter == counter2 :", counter == counter2)
}
