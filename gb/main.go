package main

import (
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type MainFile struct {
	Data map[string]Object
}

type Object struct {
	name string
	value int
}

var data MainFile

const (
	CURLY_OPEN = iota
	CURLY_CLOSE
)

func next_token(data byte[]) int {
}

func main() {
	if len(os.Args) == 1 {
		fmt.Print("format: ./gb jsonfile");
		return;
	}

	filename := os.Args[1];

	json, err := os.ReadFile(filename);

	check(err);

	for _, c := range json {
		fmt.Printf("%c\n", c);
	}
}
