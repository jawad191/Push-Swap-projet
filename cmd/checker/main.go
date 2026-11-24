package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/piles"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	nums, err := piles.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	p := &piles.Piles{A: nums}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if err := p.Exec(line); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if p.IsSorted() && len(p.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}