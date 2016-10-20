// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResourceFactoryFactory struct {
	FactoryForStub        func(worker.Client) resource.ResourceFactory
	factoryForMutex       sync.RWMutex
	factoryForArgsForCall []struct {
		arg1 worker.Client
	}
	factoryForReturns struct {
		result1 resource.ResourceFactory
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactoryFactory) FactoryFor(arg1 worker.Client) resource.ResourceFactory {
	fake.factoryForMutex.Lock()
	fake.factoryForArgsForCall = append(fake.factoryForArgsForCall, struct {
		arg1 worker.Client
	}{arg1})
	fake.recordInvocation("FactoryFor", []interface{}{arg1})
	fake.factoryForMutex.Unlock()
	if fake.FactoryForStub != nil {
		return fake.FactoryForStub(arg1)
	} else {
		return fake.factoryForReturns.result1
	}
}

func (fake *FakeResourceFactoryFactory) FactoryForCallCount() int {
	fake.factoryForMutex.RLock()
	defer fake.factoryForMutex.RUnlock()
	return len(fake.factoryForArgsForCall)
}

func (fake *FakeResourceFactoryFactory) FactoryForArgsForCall(i int) worker.Client {
	fake.factoryForMutex.RLock()
	defer fake.factoryForMutex.RUnlock()
	return fake.factoryForArgsForCall[i].arg1
}

func (fake *FakeResourceFactoryFactory) FactoryForReturns(result1 resource.ResourceFactory) {
	fake.FactoryForStub = nil
	fake.factoryForReturns = struct {
		result1 resource.ResourceFactory
	}{result1}
}

func (fake *FakeResourceFactoryFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.factoryForMutex.RLock()
	defer fake.factoryForMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResourceFactoryFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.ResourceFactoryFactory = new(FakeResourceFactoryFactory)