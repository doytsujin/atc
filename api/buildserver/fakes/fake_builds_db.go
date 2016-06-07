// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/api/buildserver"
	"github.com/concourse/atc/db"
)

type FakeBuildsDB struct {
	GetBuildResourcesStub        func(buildID int) ([]db.BuildInput, []db.BuildOutput, error)
	getBuildResourcesMutex       sync.RWMutex
	getBuildResourcesArgsForCall []struct {
		buildID int
	}
	getBuildResourcesReturns struct {
		result1 []db.BuildInput
		result2 []db.BuildOutput
		result3 error
	}
	GetConfigByBuildIDStub        func(buildID int) (atc.Config, db.ConfigVersion, error)
	getConfigByBuildIDMutex       sync.RWMutex
	getConfigByBuildIDArgsForCall []struct {
		buildID int
	}
	getConfigByBuildIDReturns struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}
}

func (fake *FakeBuildsDB) GetBuildResources(buildID int) ([]db.BuildInput, []db.BuildOutput, error) {
	fake.getBuildResourcesMutex.Lock()
	fake.getBuildResourcesArgsForCall = append(fake.getBuildResourcesArgsForCall, struct {
		buildID int
	}{buildID})
	fake.getBuildResourcesMutex.Unlock()
	if fake.GetBuildResourcesStub != nil {
		return fake.GetBuildResourcesStub(buildID)
	} else {
		return fake.getBuildResourcesReturns.result1, fake.getBuildResourcesReturns.result2, fake.getBuildResourcesReturns.result3
	}
}

func (fake *FakeBuildsDB) GetBuildResourcesCallCount() int {
	fake.getBuildResourcesMutex.RLock()
	defer fake.getBuildResourcesMutex.RUnlock()
	return len(fake.getBuildResourcesArgsForCall)
}

func (fake *FakeBuildsDB) GetBuildResourcesArgsForCall(i int) int {
	fake.getBuildResourcesMutex.RLock()
	defer fake.getBuildResourcesMutex.RUnlock()
	return fake.getBuildResourcesArgsForCall[i].buildID
}

func (fake *FakeBuildsDB) GetBuildResourcesReturns(result1 []db.BuildInput, result2 []db.BuildOutput, result3 error) {
	fake.GetBuildResourcesStub = nil
	fake.getBuildResourcesReturns = struct {
		result1 []db.BuildInput
		result2 []db.BuildOutput
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildsDB) GetConfigByBuildID(buildID int) (atc.Config, db.ConfigVersion, error) {
	fake.getConfigByBuildIDMutex.Lock()
	fake.getConfigByBuildIDArgsForCall = append(fake.getConfigByBuildIDArgsForCall, struct {
		buildID int
	}{buildID})
	fake.getConfigByBuildIDMutex.Unlock()
	if fake.GetConfigByBuildIDStub != nil {
		return fake.GetConfigByBuildIDStub(buildID)
	} else {
		return fake.getConfigByBuildIDReturns.result1, fake.getConfigByBuildIDReturns.result2, fake.getConfigByBuildIDReturns.result3
	}
}

func (fake *FakeBuildsDB) GetConfigByBuildIDCallCount() int {
	fake.getConfigByBuildIDMutex.RLock()
	defer fake.getConfigByBuildIDMutex.RUnlock()
	return len(fake.getConfigByBuildIDArgsForCall)
}

func (fake *FakeBuildsDB) GetConfigByBuildIDArgsForCall(i int) int {
	fake.getConfigByBuildIDMutex.RLock()
	defer fake.getConfigByBuildIDMutex.RUnlock()
	return fake.getConfigByBuildIDArgsForCall[i].buildID
}

func (fake *FakeBuildsDB) GetConfigByBuildIDReturns(result1 atc.Config, result2 db.ConfigVersion, result3 error) {
	fake.GetConfigByBuildIDStub = nil
	fake.getConfigByBuildIDReturns = struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}{result1, result2, result3}
}

var _ buildserver.BuildsDB = new(FakeBuildsDB)
