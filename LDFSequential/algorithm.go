package LDFSequential

import (
	"log"
	"time"

	"github.com/nieomylnieja/graph-coloring-go/graphs"
)

func New() *Algorithm {
	return &Algorithm{
		N: "LDF sequential",
	}
}

type Algorithm struct {
	*graphs.Graph
	Max uint32
	N   string
}

func (a Algorithm) Name() string {
	return a.N
}

func (a *Algorithm) Run(graph *graphs.Graph) (int, float64) {
	a.Graph = graph
	start := time.Now()
	for {
		finished := true
	Outer:
		for _, e := range a.M {
			if e.C == graphs.Uncolored {
				avb := make([]uint32, graphs.Colors)
				avb[0] = graphs.Taken
				finished = false

				for _, n := range e.N {
					if n.C == graphs.Uncolored && (len(e.N) < len(n.N) || (len(e.N) == len(n.N) && e.I < n.I)) {
						continue Outer
					}
					if n.C != graphs.Uncolored {
						avb[n.C] = graphs.Taken
					}
				}

				var i uint32
				for ; i < uint32(len(avb)); i++ {
					if avb[i] == graphs.Available {
						e.C = i
						break
					}
				}
			}
		}
		if finished {
			break
		}
	}
	t := time.Since(start)
	elapsed := float64(t.Nanoseconds()) / 1000000
	for _, e := range a.M {
		if e.C > a.Max {
			a.Max = e.C
		}
		for _, n := range e.N {
			if e.C == n.C {
				log.Fatal("Bad coloring!")
			}
		}
	}
	return int(a.Max), elapsed
}
