package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	fmt.Println("Graph v0.1")

	graphFilename := "input_graph"
	graphFile, _ := os.Open(graphFilename)
	defer graphFile.Close()

	Ints, _ := ReadInts(graphFile)
	fmt.Println(Ints)

	return
}
