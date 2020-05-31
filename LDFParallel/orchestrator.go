package LDFParallel

import (
	"sync"
)

func NewOrchestrator(vertices int, stop *sync.WaitGroup) *Orchestrator {
	o := &Orchestrator{
		Vertices:    vertices,
		colorNotify: make(chan ColorEvent, vertices),
		wait:        &sync.WaitGroup{},
		waitNotify:  make(chan bool, vertices),
		gate:        &sync.WaitGroup{},
		gateNotify:  make(chan bool, vertices),
		stop:        stop,
	}
	o.wait.Add(1)
	return o
}

type OrchestratorRelay struct {
	ColorNotify chan ColorEvent
	WaitNotify  chan bool
	GateNotify  chan bool
	Wait        *sync.WaitGroup
	Gate        *sync.WaitGroup
}

type Orchestrator struct {
	Vertices    int
	colorNotify chan ColorEvent
	wait        *sync.WaitGroup
	waitNotify  chan bool
	gate        *sync.WaitGroup
	gateNotify  chan bool
	Max         uint32
	stop        *sync.WaitGroup
}

func (o *Orchestrator) run() {
	for {
		colored := 0
		for i := 0; i < o.Vertices; i++ {
			cE := <-o.colorNotify
			//fmt.Println(cE)
			if cE.Color > 0 {
				colored++
				if cE.Color > o.Max {
					o.Max = cE.Color
				}
			}
		}
		o.gate.Add(1)
		o.wait.Done()
		if colored == 0 {
			panic("deadlock, no vertices were colored")
		}
		o.Vertices -= colored
		if o.Vertices == 0 {
			break
		}
		for i := 0; i < o.Vertices; i++ {
			<-o.waitNotify
		}
		o.wait.Add(1)
		o.gate.Done()
	}
	o.stop.Done()
}

func (o *Orchestrator) relay() OrchestratorRelay {
	return OrchestratorRelay{
		ColorNotify: o.colorNotify,
		WaitNotify:  o.waitNotify,
		GateNotify:  o.gateNotify,
		Wait:        o.wait,
		Gate:        o.gate,
	}
}