package graph

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type Graph struct {
	N, M  int
	Cycle []int
	Edges []Edge
}

type Edge struct {
	A, B int
}

func Copy(from, to *Graph) {
	to.M = from.M
	to.N = from.N
	to.Cycle = make([]int, 0)
	for _, v := range from.Cycle {
		to.Cycle = append(to.Cycle, v)
	}
	to.Edges = make([]Edge, 0)
	for _, v := range from.Edges {
		to.Edges = append(to.Edges, v)
	}
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	var result []int
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func ReadGraph(filename string) Graph {
	var G Graph
	graphFile, _ := os.Open(filename)
	defer graphFile.Close()
	Ints, _ := ReadInts(graphFile)
	G.N = Ints[0]
	G.M = Ints[1]
	G.Cycle = Ints[2*G.N+6:]
	G.Edges = make([]Edge, 0)
	for i := 2; i < G.N*2+6; i += 2 {
		G.Edges = append(G.Edges, Edge{Ints[i], Ints[i+1]})
	}
	return G
}
