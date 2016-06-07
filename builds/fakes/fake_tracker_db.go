// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/builds"
	"github.com/concourse/atc/db"
)

type FakeTrackerDB struct {
	GetAllStartedBuildsStub        func() ([]db.Build, error)
	getAllStartedBuildsMutex       sync.RWMutex
	getAllStartedBuildsArgsForCall []struct{}
	getAllStartedBuildsReturns     struct {
		result1 []db.Build
		result2 error
	}
}

func (fake *FakeTrackerDB) GetAllStartedBuilds() ([]db.Build, error) {
	fake.getAllStartedBuildsMutex.Lock()
	fake.getAllStartedBuildsArgsForCall = append(fake.getAllStartedBuildsArgsForCall, struct{}{})
	fake.getAllStartedBuildsMutex.Unlock()
	if fake.GetAllStartedBuildsStub != nil {
		return fake.GetAllStartedBuildsStub()
	} else {
		return fake.getAllStartedBuildsReturns.result1, fake.getAllStartedBuildsReturns.result2
	}
}

func (fake *FakeTrackerDB) GetAllStartedBuildsCallCount() int {
	fake.getAllStartedBuildsMutex.RLock()
	defer fake.getAllStartedBuildsMutex.RUnlock()
	return len(fake.getAllStartedBuildsArgsForCall)
}

func (fake *FakeTrackerDB) GetAllStartedBuildsReturns(result1 []db.Build, result2 error) {
	fake.GetAllStartedBuildsStub = nil
	fake.getAllStartedBuildsReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

var _ builds.TrackerDB = new(FakeTrackerDB)
