package LDFParallel

import (
	"sync"
	"time"

	"github.com/nieomylnieja/graph-coloring-go/graphs"
)

func New() *Algorithm {
	return &Algorithm{
		N: "LDF parallel",
	}
}

type ColorEvent struct {
	Color  uint32
	Vertex uint32
}

type Algorithm struct {
	Graph        map[uint32]*Vertex
	Orchestrator *Orchestrator
	N            string
}

func (a Algorithm) Name() string {
	return a.N
}

func (a *Algorithm) Run(graph *graphs.Graph) (int, float64) {
	a.Graph = toParallelGraph(graph)

	start := time.Now()
	a.run()
	t := time.Since(start)

	elapsed := float64(t.Nanoseconds()) / 1000000
	return int(a.Orchestrator.Max), elapsed
}

func (a *Algorithm) run() {
	var stop sync.WaitGroup
	stop.Add(1)
	a.Orchestrator = NewOrchestrator(len(a.Graph), &stop)
	go a.Orchestrator.run()
	for i := range a.Graph {
		a.Graph[i].Relay = a.Orchestrator.relay()
		go a.Graph[i].color()
	}
	stop.Wait()
}
