package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	textPtr := flag.String("t", "", "Text to parse. (Required)")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Hello world: %s \n", *textPtr)
}
