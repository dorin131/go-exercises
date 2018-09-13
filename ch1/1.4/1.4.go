// print the lines which repeat either in files or standard input
// e.g. go run 1.4.go ./file1.txt ./file2.txt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, locations)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(f, counts, locations)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(locations[line], ","))
		}
	}
}

func countLines(f *os.File, counts map[string]int, locations map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !contains(locations[input.Text()], f.Name()) {
			locations[input.Text()] = append(locations[input.Text()], f.Name())
		}
	}
}

func contains(slice []string, name string) bool {
	for _, val := range slice {
		if val == name {
			return true
		}
	}
	return false
}
