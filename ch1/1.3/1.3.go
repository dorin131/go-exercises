package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// echo using for loop
	start1 := time.Now()
	for _, arg := range os.Args[1:] {
		fmt.Print(arg + " ")
	}
	fmt.Printf("\n%.6f seconds\n", time.Since(start1).Seconds())

	// echo using strings.Join
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.6f seconds\n", time.Since(start2).Seconds())
}
