package LDFParallel

import (
	"testing"

	"github.com/nieomylnieja/graph-coloring-go/graphs"
)

var a *Algorithm
var g *graphs.Graph

func BenchmarkAlgorithm_LDFParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a.Graph = toParallelGraph(g)
		a.run()
	}
}

// TODO change path to support reading the instances from other packages than main
func init() {
	r := graphs.DimacsReader{}
	g = r.Read("gc500")
	g.ReIndexVertices()
	a = &Algorithm{}
}
