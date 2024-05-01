package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

func synchronousPipeline(input *csv.Reader) {
	fmt.Println("--Synchronous pipeline----")
	// Ignore the first row
	_, err := input.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		rec, err := input.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		out := encode(convert(parse(rec)))
		fmt.Println(string(out))
	}
}
