// Code generated by counterfeiter. DO NOT EDIT.
package commandparserfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/stembuild/commandparser"
)

type FakeConstructCmdValidator struct {
	LGPOInDirectoryStub        func() bool
	lGPOInDirectoryMutex       sync.RWMutex
	lGPOInDirectoryArgsForCall []struct {
	}
	lGPOInDirectoryReturns struct {
		result1 bool
	}
	lGPOInDirectoryReturnsOnCall map[int]struct {
		result1 bool
	}
	PopulatedArgsStub        func(...string) bool
	populatedArgsMutex       sync.RWMutex
	populatedArgsArgsForCall []struct {
		arg1 []string
	}
	populatedArgsReturns struct {
		result1 bool
	}
	populatedArgsReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConstructCmdValidator) LGPOInDirectory() bool {
	fake.lGPOInDirectoryMutex.Lock()
	ret, specificReturn := fake.lGPOInDirectoryReturnsOnCall[len(fake.lGPOInDirectoryArgsForCall)]
	fake.lGPOInDirectoryArgsForCall = append(fake.lGPOInDirectoryArgsForCall, struct {
	}{})
	fake.recordInvocation("LGPOInDirectory", []interface{}{})
	fake.lGPOInDirectoryMutex.Unlock()
	if fake.LGPOInDirectoryStub != nil {
		return fake.LGPOInDirectoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lGPOInDirectoryReturns
	return fakeReturns.result1
}

func (fake *FakeConstructCmdValidator) LGPOInDirectoryCallCount() int {
	fake.lGPOInDirectoryMutex.RLock()
	defer fake.lGPOInDirectoryMutex.RUnlock()
	return len(fake.lGPOInDirectoryArgsForCall)
}

func (fake *FakeConstructCmdValidator) LGPOInDirectoryCalls(stub func() bool) {
	fake.lGPOInDirectoryMutex.Lock()
	defer fake.lGPOInDirectoryMutex.Unlock()
	fake.LGPOInDirectoryStub = stub
}

func (fake *FakeConstructCmdValidator) LGPOInDirectoryReturns(result1 bool) {
	fake.lGPOInDirectoryMutex.Lock()
	defer fake.lGPOInDirectoryMutex.Unlock()
	fake.LGPOInDirectoryStub = nil
	fake.lGPOInDirectoryReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConstructCmdValidator) LGPOInDirectoryReturnsOnCall(i int, result1 bool) {
	fake.lGPOInDirectoryMutex.Lock()
	defer fake.lGPOInDirectoryMutex.Unlock()
	fake.LGPOInDirectoryStub = nil
	if fake.lGPOInDirectoryReturnsOnCall == nil {
		fake.lGPOInDirectoryReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.lGPOInDirectoryReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConstructCmdValidator) PopulatedArgs(arg1 ...string) bool {
	fake.populatedArgsMutex.Lock()
	ret, specificReturn := fake.populatedArgsReturnsOnCall[len(fake.populatedArgsArgsForCall)]
	fake.populatedArgsArgsForCall = append(fake.populatedArgsArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("PopulatedArgs", []interface{}{arg1})
	fake.populatedArgsMutex.Unlock()
	if fake.PopulatedArgsStub != nil {
		return fake.PopulatedArgsStub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.populatedArgsReturns
	return fakeReturns.result1
}

func (fake *FakeConstructCmdValidator) PopulatedArgsCallCount() int {
	fake.populatedArgsMutex.RLock()
	defer fake.populatedArgsMutex.RUnlock()
	return len(fake.populatedArgsArgsForCall)
}

func (fake *FakeConstructCmdValidator) PopulatedArgsCalls(stub func(...string) bool) {
	fake.populatedArgsMutex.Lock()
	defer fake.populatedArgsMutex.Unlock()
	fake.PopulatedArgsStub = stub
}

func (fake *FakeConstructCmdValidator) PopulatedArgsArgsForCall(i int) []string {
	fake.populatedArgsMutex.RLock()
	defer fake.populatedArgsMutex.RUnlock()
	argsForCall := fake.populatedArgsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConstructCmdValidator) PopulatedArgsReturns(result1 bool) {
	fake.populatedArgsMutex.Lock()
	defer fake.populatedArgsMutex.Unlock()
	fake.PopulatedArgsStub = nil
	fake.populatedArgsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConstructCmdValidator) PopulatedArgsReturnsOnCall(i int, result1 bool) {
	fake.populatedArgsMutex.Lock()
	defer fake.populatedArgsMutex.Unlock()
	fake.PopulatedArgsStub = nil
	if fake.populatedArgsReturnsOnCall == nil {
		fake.populatedArgsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.populatedArgsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConstructCmdValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lGPOInDirectoryMutex.RLock()
	defer fake.lGPOInDirectoryMutex.RUnlock()
	fake.populatedArgsMutex.RLock()
	defer fake.populatedArgsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConstructCmdValidator) recordInvocation(key string, args []interface{}) {
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

var _ commandparser.ConstructCmdValidator = new(FakeConstructCmdValidator)
