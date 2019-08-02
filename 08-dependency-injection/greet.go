package main

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
