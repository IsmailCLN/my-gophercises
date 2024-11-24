package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit("Failed to open the CSV file: %s", *csvFileName)
	}
	_ = file
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
