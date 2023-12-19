package main

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/cpu"
)

func main() {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	total := float64(before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(before.Idle)/total*100)
}
