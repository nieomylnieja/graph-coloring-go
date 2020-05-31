package greedy

import (
	"fmt"
	"log"
	"time"

	"github.com/nieomylnieja/graph-coloring-go/graphs"
)

func New() *Algorithm {
	return &Algorithm{
		N: "Greedy",
	}
}

type Algorithm struct {
	*graphs.Graph
	N   string
	Max int
}

func (a Algorithm) Name() string {
	return a.N
}

func (a *Algorithm) Run(graph *graphs.Graph) (int, float64) {
	a.Graph = graph
	start := time.Now()
	for _, e := range a.M {
		a.lowestColor(e)
	}
	t := time.Since(start)
	elapsed := float64(t.Nanoseconds()) / 1000000
	uncoloredCtr := 0
	for i, e := range a.M {
		if e.C == graphs.Uncolored {
			uncoloredCtr++
			fmt.Printf("Not colored: %d\n", i)
			continue
		}
		for _, n := range e.N {
			if e.C == n.C {
				log.Fatal("Bad coloring!")
			}
		}
	}
	if uncoloredCtr > 0 {
		fmt.Printf("Uncolored Num: %d\n", uncoloredCtr)
	}
	return a.Max, elapsed
}

func (a *Algorithm) lowestColor(e *graphs.Vertex) {
	avb := make([]uint32, graphs.Colors)
	avb[0] = graphs.Taken
	for _, n := range e.N {
		if n.C != graphs.Uncolored {
			avb[n.C] = graphs.Taken
		}
	}
	for i, c := range avb {
		if c == graphs.Available {
			e.C = uint32(i)
			if i > a.Max {
				a.Max = i
			}
			return
		}
	}
}
