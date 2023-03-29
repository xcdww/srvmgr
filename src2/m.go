package main

import (
	"fmt"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	a, b := rCommand("asd")
	fmt.Println(a, b)
}

func rCommand(command string) (string, string) {
	co := strings.Split(command, "_")
	return co[0], co[1]
}
