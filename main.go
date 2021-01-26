package main

import (
	"os"
	"wiredcraft-hugo/pipeline"
)

func main() {
	envName := os.Args[1]
	err := pipeline.Execute(envName)
	if err != nil {
		panic(err)
	}
	// TODO: exit if an error occurred
}
