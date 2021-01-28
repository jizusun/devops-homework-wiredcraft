package main

import (
	"fmt"
	"os"
	"wiredcraft-hugo/pipeline"
)

func main() {
	args := os.Args[1:]
	err := pipeline.Execute(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// TODO: exit if an error occurred
}
