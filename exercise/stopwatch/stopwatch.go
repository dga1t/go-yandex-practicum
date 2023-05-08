package stopwatch

import (
	"time"
)

type Stopwatch struct {
	start time.Time
	splits []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.start = time.Now()
	sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
	t := time.Now()
	elapsed := t.Sub(sw.start)
	sw.splits = append(sw.splits, elapsed)
}

func (sw Stopwatch) GetResults() []time.Duration {
	return sw.splits
}