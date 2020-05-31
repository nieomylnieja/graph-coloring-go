package main

import (
	"fmt"

	"github.com/nieomylnieja/graph-coloring-go/LDFParallel"
	"github.com/nieomylnieja/graph-coloring-go/LDFSequential"
	"github.com/nieomylnieja/graph-coloring-go/graphs"
	"github.com/nieomylnieja/graph-coloring-go/greedy"
)

type Algorithm interface {
	Run(graph *graphs.Graph) (int, float64)
	Name() string
}

var mainInstances = []string{"queen6_6", "miles250", "gc1000_300013", "le450_5a", "myciel4", "fpsol2.i.1", "inithx.i.1", "mulsol.i.1"}

func main() {
	for _, instance := range mainInstances {
		r := graphs.DimacsReader{}
		g := r.Read(instance)
		fmt.Println(g)
		samples := 100
		algorithms := []Algorithm{greedy.New(), LDFSequential.New(), LDFParallel.New()}
		for _, a := range algorithms {
			runAlgorithm(g, a, samples)
		}
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
	fmt.Printf("Algorithm: %s | Samples: %d | Average colors used: %d | Average time: %f [ms]\n",
		algorithm.Name(), samples, avgC, avgT)
}
