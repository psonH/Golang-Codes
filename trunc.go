package main

import (
	"fmt"
)

func main() {
	var x float64
  fmt.Printf("Enter a floating number:")
  fmt.Scan(&x)
	fmt.Printf("\nThe truncated version is %.0f", x)
}
