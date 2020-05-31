package LDFParallel

import (
	"github.com/nieomylnieja/graph-coloring-go/graphs"
)

const (
	taken     = true
	available = false
)

type Vertex struct {
	Index      uint32
	Degree     uint32
	PublicChan chan ColorEvent
	Neighbours map[uint32]Neighbour
	ColorsPool []bool
	Relay      OrchestratorRelay
}

type Neighbour struct {
	VertexDegree uint32
	PublicChan   chan ColorEvent
}

func (v *Vertex) color() {
	color := uint32(0)
	for {
		if v.hasLargestDegree() {
			color = v.chooseColor()
		}
		event := ColorEvent{
			Color:  color,
			Vertex: v.Index,
		}
		v.notifyNeighbours(event)
		v.getNotifications()
		v.synchronize(event)
		if color != graphs.Uncolored {
			break
		}
	}
	close(v.PublicChan)
}

func (v *Vertex) notifyNeighbours(event ColorEvent) {
	for idx := range v.Neighbours {
		v.Neighbours[idx].PublicChan <- event
	}
}

func (v *Vertex) getNotifications() {
	neighbours := len(v.Neighbours)
	var event ColorEvent
	for i := 0; i < neighbours; i++ {
		event = <-v.PublicChan
		if event.Color != graphs.Uncolored {
			v.ColorsPool[event.Color] = taken
			delete(v.Neighbours, event.Vertex)
		}
	}
}

func (v *Vertex) hasLargestDegree() bool {
	for idx := range v.Neighbours {
		if v.Neighbours[idx].VertexDegree > v.Degree || (v.Neighbours[idx].VertexDegree == v.Degree && idx > v.Index) {
			return false
		}
	}
	return true
}

func (v *Vertex) chooseColor() uint32 {
	for c := range v.ColorsPool {
		if v.ColorsPool[c] == available {
			return uint32(c)
		}
	}
	panic("Exceeded assumed colors pool")
}

func (v *Vertex) synchronize(event ColorEvent) {
	v.Relay.ColorNotify <- event
	if event.Color == graphs.Uncolored {
		v.Relay.Wait.Wait()
		v.Relay.WaitNotify <- true
		v.Relay.Gate.Wait()
	}
}
