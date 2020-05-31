package graphs

import (
	"fmt"
)

const (
	Colors    uint32 = 256
	Uncolored uint32 = 0
	Colored   uint32 = 1
	Taken     uint32 = 1
	Available uint32 = 0
)

type Vertex struct {
	N Neighbours
	C uint32
	I uint32
}

type Edge [2]uint32

type Neighbours map[uint32]*Vertex

type Graph struct {
	M    map[uint32]*Vertex
	E    []Edge
	I    []uint32
	Name string
}

func NewGraph(name string) *Graph {
	return &Graph{
		M:    make(map[uint32]*Vertex),
		Name: name,
	}
}

func (g *Graph) Add(v1, v2 uint64) {
	u32V1, u32V2 := uint32(v1), uint32(v2)
	g.E = append(g.E, [2]uint32{u32V1, u32V2})
	if _, ok := g.M[u32V1]; !ok {
		g.M[u32V1] = &Vertex{
			N: make(Neighbours),
			C: Uncolored,
		}
	}
	if _, ok := g.M[u32V2]; !ok {
		g.M[u32V2] = &Vertex{
			N: make(Neighbours),
			C: Uncolored,
		}
	}
	g.M[u32V1].N[u32V2] = g.M[u32V2]
	g.M[u32V2].N[u32V1] = g.M[u32V1]
}

func (g *Graph) ReIndexVertices() {
	i := uint32(0)
	for _, v := range g.M {
		v.C = Uncolored
		v.I = i
		i++
	}
}

func (g Graph) String() string {
	maxDegree := 0
	for _, v := range g.M {
		if len(v.N) > maxDegree {
			maxDegree = len(v.N)
		}
	}
	return fmt.Sprintf("Graph: %s [%d | %d] {MAX degree: %d}", g.Name, len(g.M), len(g.E), maxDegree)
}
