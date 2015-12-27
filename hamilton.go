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

func IsEdgeInCycle(e graph.Edge, cycle []int) bool {
	for i := range cycle {
		if i == 0 {
			continue
		}
		if (e == graph.Edge{cycle[i], cycle[i-1]}) {
			return true
		}
		if (e == graph.Edge{cycle[i-1], cycle[i]}) {
			return true
		}
	}
	return false
}

// Load graph, and setups another graphs(morphing and encryption)
func (alice *Alice) LoadGraph(filename string) {
	alice.rsa.Init()

	// Loaded graph G
	alice.G = graph.ReadGraph(filename)

	// Isomorphing graph H
	rand := crypto.Random(2, 1423)
	graph.Copy(&alice.G, &alice.H) // from g to h
	for i := range alice.H.Edges {
		alice.H.Edges[i].A += rand
		alice.H.Edges[i].B += rand
	}
	for k := range alice.H.Cycle {
		alice.H.Cycle[k] += rand
	}

	// Enncrypt graph
	graph.Copy(&alice.H, &alice.F)
	for i := range alice.F.Edges {
		alice.F.Edges[i].A = crypto.Pow(alice.F.Edges[i].A, alice.rsa.D, alice.rsa.N)
		alice.F.Edges[i].B = crypto.Pow(alice.F.Edges[i].B, alice.rsa.D, alice.rsa.N)
	}
	fmt.Println("Loaded graph:", alice.G)
	fmt.Println("Mutate graph:", alice.H)
	fmt.Println("Encrypted graph:", alice.F)

}

// Answer for question 1
func (alice *Alice) GetCycle() graph.Graph {
	var M graph.Graph
	graph.Copy(&alice.H, &M)
	fmt.Println("Making cycle:")
	for k := range M.Edges {
		if graph.IsEdgeInCycle(M.Edges[k], M.Cycle) {
			continue
		} else {
			M.Edges[k].A = crypto.Pow(M.Edges[k].A, alice.rsa.D, alice.rsa.N)
			M.Edges[k].B = crypto.Pow(M.Edges[k].B, alice.rsa.D, alice.rsa.N)
		}
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
