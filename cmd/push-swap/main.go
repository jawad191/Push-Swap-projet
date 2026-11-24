package main

import (
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
	p.Sort()

	for _, op := range p.Operations {
		fmt.Println(op)
	}
}