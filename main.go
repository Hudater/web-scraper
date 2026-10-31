package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) - 1 < 1 {
		fmt.Println("no website provided\n")
		os.Exit(1)
	}
	
	if len(os.Args) - 1 > 1 {
		fmt.Println("too many arguments provided\n")
		os.Exit(1)
	}
	
	if len(os.Args) - 1 == 1 {
		fmt.Printf("starting crawl of: %v\n", os.Args[1:])
	}
}
