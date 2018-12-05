package executor

import (
	"time"
)

type periodic struct {
	refreshPeriod time.Duration
	executionFunc func(stopCh <-chan struct{})
	tickStop      tickerStopper
}

// NewPeriodic creates a periodic executor, which calls given executionFunc periodically.
func NewPeriodic(period time.Duration, executionFunc func(stopCh <-chan struct{})) *periodic {
	t := time.NewTicker(period)
	w := tickerWrapper{internal:t}
	return &periodic{
		refreshPeriod: period,
		executionFunc: executionFunc,
		tickStop:      w,
	}
}

type tickerStopper interface {
	Tick() chan<- time.Time
	Stopp()
}

type tickerWrapper struct {
	internal *time.Ticker
}

func (tw *tickerWrapper) Tick() <-chan time.Time {
	return tw.internal.C
}

func (tw *tickerWrapper) Stopp() {
	tw.internal.Stop()
}

// Run starts the periodic work
func (e *periodic) Run(stopCh <-chan struct{}) {
	go func() {
		for {
			e.executionFunc(stopCh)
			select {
			case <-stopCh:
				ticker.Stop()
				return
			case <-ticker.C:
			}
		}
	}()
}
