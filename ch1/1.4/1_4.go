// Exercise 1.4
// Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineAppearances := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "", counts, lineAppearances)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, lineAppearances)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s%s\n", n, line, lineAppearances[line])
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, lineAppearances map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		// Probably definitely not very efficient
		if filename != "" {
			s := []string{lineAppearances[input.Text()], filename}
			lineAppearances[input.Text()] = strings.Join(s, "\t")
		}

	}
}
