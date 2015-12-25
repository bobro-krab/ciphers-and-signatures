package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"zi/crypto"
	"zi/shifr"
)

type Alice struct {
	rsa shifr.RSA
	G   Graph // Initial graph
	H   Graph // Isomorphic graph
	F   Graph // Encripted isomoprhic graph
}

type Graph struct {
	N, M  int
	Cycle []int
	Edges []Edge
}

type Edge struct {
	a, b int
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

// Load graph, and setups another graphs(morphing and encryption)
func (alice Alice) LoadGraph(filename string) {
	alice.rsa.Init()

	// Loaded graph G
	alice.G = ReadGraph(filename)
	fmt.Println("Loaded graph:", alice.G)

	// Isomorphing graph H
	rand := crypto.Random(2, 1423)
	alice.H = alice.G
	for _, v := range alice.H.Edges {
		v.a += rand
		v.b += rand
	}
	fmt.Println("Mutate graph:", alice.H)

	// Enncrypt graph
	alice.F = alice.G
	for _, v := range alice.F.Edges {
		v.a = crypto.Pow(v.a, alice.rsa.D, alice.rsa.N)
		v.b = crypto.Pow(v.b, alice.rsa.D, alice.rsa.N)
	}
	fmt.Println("Encrypted graph:", alice.F)
}

func main() {
	fmt.Println("Graph v0.1")

	var alice Alice
	alice.LoadGraph("input_graph")

	return
}
