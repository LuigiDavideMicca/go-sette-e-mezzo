package main

import (
	"bufio"
	"os"
)

func createReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readInput() (rune, error) {
	reader := createReader()
	card, _, err := reader.ReadRune()

	return card, err
}
