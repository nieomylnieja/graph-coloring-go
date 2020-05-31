package main

import (
	"fmt"

	"github.com/nieomylnieja/graph-coloring-go/LDFParallel"
	"github.com/nieomylnieja/graph-coloring-go/graphs"
)

type Algorithm interface {
	Run(graph *graphs.Graph) (int, float64)
}

func main() {
	r := graphs.DimacsReader{}
	g := r.Read(true)
	fmt.Println(g)
	samples := 10
	algorithms := []Algorithm{&LDFParallel.Algorithm{}}
	for _, a := range algorithms {
		runAlgorithm(g, a, samples)
	}
}

func runAlgorithm(graph *graphs.Graph, algorithm Algorithm, samples int) {
	totalT, totalC := 0., 0
	for i := 0; i < samples; i++ {
		graph.ReIndexVertices()
		c, t := algorithm.Run(graph)
		totalC += c
		totalT += t
	}
	avgC := totalC / samples
	avgT := totalT / float64(samples)
	fmt.Printf("Samples: %d | Average colors used: %d | Average time: %f [ms]\n",
		samples, avgC, avgT)
}
