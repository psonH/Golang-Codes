package main

import (
	"fmt"
  "strconv"
)

func main() {
  var appleNum int
	fmt.Printf("Enter the number of apples:")
	fmt.Scan(&appleNum)
	fmt.Println("\nThere are " + strconv.Itoa(appleNum) + " apples.")
}
