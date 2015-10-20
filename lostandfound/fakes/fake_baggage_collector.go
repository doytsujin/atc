// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/lostandfound"
)

type FakeBaggageCollector struct {
	CollectStub        func() error
	collectMutex       sync.RWMutex
	collectArgsForCall []struct{}
	collectReturns     struct {
		result1 error
	}
}

func (fake *FakeBaggageCollector) Collect() error {
	fake.collectMutex.Lock()
	fake.collectArgsForCall = append(fake.collectArgsForCall, struct{}{})
	fake.collectMutex.Unlock()
	if fake.CollectStub != nil {
		return fake.CollectStub()
	} else {
		return fake.collectReturns.result1
	}
}

func (fake *FakeBaggageCollector) CollectCallCount() int {
	fake.collectMutex.RLock()
	defer fake.collectMutex.RUnlock()
	return len(fake.collectArgsForCall)
}

func (fake *FakeBaggageCollector) CollectReturns(result1 error) {
	fake.CollectStub = nil
	fake.collectReturns = struct {
		result1 error
	}{result1}
}

var _ lostandfound.BaggageCollector = new(FakeBaggageCollector)