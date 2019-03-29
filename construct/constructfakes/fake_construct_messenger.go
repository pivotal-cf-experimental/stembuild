// Code generated by counterfeiter. DO NOT EDIT.
package constructfakes

import (
	sync "sync"

	construct "github.com/cloudfoundry-incubator/stembuild/construct"
)

type FakeConstructMessenger struct {
	EnableWinRMStartedStub        func()
	enableWinRMStartedMutex       sync.RWMutex
	enableWinRMStartedArgsForCall []struct {
	}
	EnableWinRMSucceededStub        func()
	enableWinRMSucceededMutex       sync.RWMutex
	enableWinRMSucceededArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConstructMessenger) EnableWinRMStarted() {
	fake.enableWinRMStartedMutex.Lock()
	fake.enableWinRMStartedArgsForCall = append(fake.enableWinRMStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("EnableWinRMStarted", []interface{}{})
	fake.enableWinRMStartedMutex.Unlock()
	if fake.EnableWinRMStartedStub != nil {
		fake.EnableWinRMStartedStub()
	}
}

func (fake *FakeConstructMessenger) EnableWinRMStartedCallCount() int {
	fake.enableWinRMStartedMutex.RLock()
	defer fake.enableWinRMStartedMutex.RUnlock()
	return len(fake.enableWinRMStartedArgsForCall)
}

func (fake *FakeConstructMessenger) EnableWinRMStartedCalls(stub func()) {
	fake.enableWinRMStartedMutex.Lock()
	defer fake.enableWinRMStartedMutex.Unlock()
	fake.EnableWinRMStartedStub = stub
}

func (fake *FakeConstructMessenger) EnableWinRMSucceeded() {
	fake.enableWinRMSucceededMutex.Lock()
	fake.enableWinRMSucceededArgsForCall = append(fake.enableWinRMSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("EnableWinRMSucceeded", []interface{}{})
	fake.enableWinRMSucceededMutex.Unlock()
	if fake.EnableWinRMSucceededStub != nil {
		fake.EnableWinRMSucceededStub()
	}
}

func (fake *FakeConstructMessenger) EnableWinRMSucceededCallCount() int {
	fake.enableWinRMSucceededMutex.RLock()
	defer fake.enableWinRMSucceededMutex.RUnlock()
	return len(fake.enableWinRMSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) EnableWinRMSucceededCalls(stub func()) {
	fake.enableWinRMSucceededMutex.Lock()
	defer fake.enableWinRMSucceededMutex.Unlock()
	fake.EnableWinRMSucceededStub = stub
}

func (fake *FakeConstructMessenger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.enableWinRMStartedMutex.RLock()
	defer fake.enableWinRMStartedMutex.RUnlock()
	fake.enableWinRMSucceededMutex.RLock()
	defer fake.enableWinRMSucceededMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConstructMessenger) recordInvocation(key string, args []interface{}) {
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

var _ construct.ConstructMessenger = new(FakeConstructMessenger)