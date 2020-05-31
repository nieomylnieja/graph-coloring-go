package LDFParallel

import "github.com/nieomylnieja/graph-coloring-go/graphs"

func toParallelGraph(graph *graphs.Graph) map[uint32]*Vertex {
	// due to the fact that maps are unordered we get a randomized slice
	vs := make(map[uint32]*Vertex, len(graph.M))
	chans := make(map[uint32]chan ColorEvent, len(graph.M))
	for _, v := range graph.M {
		chans[v.I] = make(chan ColorEvent, len(v.N))
	}

	for _, v := range graph.M {
		neighbours := make(map[uint32]Neighbour, len(v.N))
		for _, n := range v.N {
			neighbours[n.I] = Neighbour{
				PublicChan:   chans[n.I],
				VertexDegree: uint32(len(n.N)),
			}
		}
		vs[v.I] = &Vertex{
			Neighbours: neighbours,
			ColorsPool: make([]bool, graphs.Colors+1),
			Index:      v.I,
			Degree:     uint32(len(neighbours)),
			PublicChan: chans[v.I],
		}
		// take the zero indexed element, as It's 0 is used to signify uncolored
		vs[v.I].ColorsPool[0] = true
	}
	max := &Vertex{Degree: 0}
	for _, v := range vs {
		if v.Degree > max.Degree {
			max = v
		}
	}
	return vs
}
