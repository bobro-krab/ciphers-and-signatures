package main

import (
	"fmt"
	"zi/crypto"
	"zi/graph"
	"zi/shifr"
)

type Alice struct {
	rsa shifr.RSA
	G   graph.Graph // Initial graph
	H   graph.Graph // Isomorphic graph
	F   graph.Graph // Encripted isomoprhic graph
}

// Load graph, and setups another graphs(morphing and encryption)
func (alice Alice) LoadGraph(filename string) {
	alice.rsa.Init()

	// Loaded graph G
	alice.G = graph.ReadGraph(filename)
	fmt.Println("Loaded graph:", alice.G)

	// Isomorphing graph H
	rand := crypto.Random(2, 1423)
	alice.H = alice.G
	for i := range alice.H.Edges {
		alice.H.Edges[i].A += rand
		alice.H.Edges[i].B += rand
	}
	fmt.Println("Mutate graph:", alice.H)

	// Enncrypt graph
	alice.F = alice.G
	for i := range alice.F.Edges {
		alice.F.Edges[i].A = crypto.Pow(alice.F.Edges[i].A, alice.rsa.D, alice.rsa.N)
		alice.F.Edges[i].B = crypto.Pow(alice.F.Edges[i].B, alice.rsa.D, alice.rsa.N)
	}
	fmt.Println("Encrypted graph:", alice.F)

}

// Answer for question 1
func (alice Alice) GetCycle() graph.Graph {
	dict := map[graph.Edge]graph.Edge{}
	var M graph.Graph
	M = alice.F
	for k := range M.Edges {
		dict[M.Edges[k]] = alice.H.Edges[k]
	}
	return M
}

func main() {
	fmt.Println("Graph v0.1")

	var alice Alice
	alice.LoadGraph("input_graph")
	fmt.Println(alice.GetCycle())

	return
}
