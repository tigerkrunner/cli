// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeSetOrgDefaultIsolationSegmentActor struct {
	GetIsolationSegmentByNameStub        func(string) (v3action.IsolationSegment, v3action.Warnings, error)
	getIsolationSegmentByNameMutex       sync.RWMutex
	getIsolationSegmentByNameArgsForCall []struct {
		arg1 string
	}
	getIsolationSegmentByNameReturns struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}
	getIsolationSegmentByNameReturnsOnCall map[int]struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}
	SetOrganizationDefaultIsolationSegmentStub        func(string, string) (v3action.Warnings, error)
	setOrganizationDefaultIsolationSegmentMutex       sync.RWMutex
	setOrganizationDefaultIsolationSegmentArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setOrganizationDefaultIsolationSegmentReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	setOrganizationDefaultIsolationSegmentReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) GetIsolationSegmentByName(arg1 string) (v3action.IsolationSegment, v3action.Warnings, error) {
	fake.getIsolationSegmentByNameMutex.Lock()
	ret, specificReturn := fake.getIsolationSegmentByNameReturnsOnCall[len(fake.getIsolationSegmentByNameArgsForCall)]
	fake.getIsolationSegmentByNameArgsForCall = append(fake.getIsolationSegmentByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetIsolationSegmentByName", []interface{}{arg1})
	fake.getIsolationSegmentByNameMutex.Unlock()
	if fake.GetIsolationSegmentByNameStub != nil {
		return fake.GetIsolationSegmentByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getIsolationSegmentByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) GetIsolationSegmentByNameCallCount() int {
	fake.getIsolationSegmentByNameMutex.RLock()
	defer fake.getIsolationSegmentByNameMutex.RUnlock()
	return len(fake.getIsolationSegmentByNameArgsForCall)
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) GetIsolationSegmentByNameCalls(stub func(string) (v3action.IsolationSegment, v3action.Warnings, error)) {
	fake.getIsolationSegmentByNameMutex.Lock()
	defer fake.getIsolationSegmentByNameMutex.Unlock()
	fake.GetIsolationSegmentByNameStub = stub
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) GetIsolationSegmentByNameArgsForCall(i int) string {
	fake.getIsolationSegmentByNameMutex.RLock()
	defer fake.getIsolationSegmentByNameMutex.RUnlock()
	argsForCall := fake.getIsolationSegmentByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) GetIsolationSegmentByNameReturns(result1 v3action.IsolationSegment, result2 v3action.Warnings, result3 error) {
	fake.getIsolationSegmentByNameMutex.Lock()
	defer fake.getIsolationSegmentByNameMutex.Unlock()
	fake.GetIsolationSegmentByNameStub = nil
	fake.getIsolationSegmentByNameReturns = struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) GetIsolationSegmentByNameReturnsOnCall(i int, result1 v3action.IsolationSegment, result2 v3action.Warnings, result3 error) {
	fake.getIsolationSegmentByNameMutex.Lock()
	defer fake.getIsolationSegmentByNameMutex.Unlock()
	fake.GetIsolationSegmentByNameStub = nil
	if fake.getIsolationSegmentByNameReturnsOnCall == nil {
		fake.getIsolationSegmentByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.IsolationSegment
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getIsolationSegmentByNameReturnsOnCall[i] = struct {
		result1 v3action.IsolationSegment
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) SetOrganizationDefaultIsolationSegment(arg1 string, arg2 string) (v3action.Warnings, error) {
	fake.setOrganizationDefaultIsolationSegmentMutex.Lock()
	ret, specificReturn := fake.setOrganizationDefaultIsolationSegmentReturnsOnCall[len(fake.setOrganizationDefaultIsolationSegmentArgsForCall)]
	fake.setOrganizationDefaultIsolationSegmentArgsForCall = append(fake.setOrganizationDefaultIsolationSegmentArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetOrganizationDefaultIsolationSegment", []interface{}{arg1, arg2})
	fake.setOrganizationDefaultIsolationSegmentMutex.Unlock()
	if fake.SetOrganizationDefaultIsolationSegmentStub != nil {
		return fake.SetOrganizationDefaultIsolationSegmentStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setOrganizationDefaultIsolationSegmentReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) SetOrganizationDefaultIsolationSegmentCallCount() int {
	fake.setOrganizationDefaultIsolationSegmentMutex.RLock()
	defer fake.setOrganizationDefaultIsolationSegmentMutex.RUnlock()
	return len(fake.setOrganizationDefaultIsolationSegmentArgsForCall)
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) SetOrganizationDefaultIsolationSegmentCalls(stub func(string, string) (v3action.Warnings, error)) {
	fake.setOrganizationDefaultIsolationSegmentMutex.Lock()
	defer fake.setOrganizationDefaultIsolationSegmentMutex.Unlock()
	fake.SetOrganizationDefaultIsolationSegmentStub = stub
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) SetOrganizationDefaultIsolationSegmentArgsForCall(i int) (string, string) {
	fake.setOrganizationDefaultIsolationSegmentMutex.RLock()
	defer fake.setOrganizationDefaultIsolationSegmentMutex.RUnlock()
	argsForCall := fake.setOrganizationDefaultIsolationSegmentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) SetOrganizationDefaultIsolationSegmentReturns(result1 v3action.Warnings, result2 error) {
	fake.setOrganizationDefaultIsolationSegmentMutex.Lock()
	defer fake.setOrganizationDefaultIsolationSegmentMutex.Unlock()
	fake.SetOrganizationDefaultIsolationSegmentStub = nil
	fake.setOrganizationDefaultIsolationSegmentReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) SetOrganizationDefaultIsolationSegmentReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.setOrganizationDefaultIsolationSegmentMutex.Lock()
	defer fake.setOrganizationDefaultIsolationSegmentMutex.Unlock()
	fake.SetOrganizationDefaultIsolationSegmentStub = nil
	if fake.setOrganizationDefaultIsolationSegmentReturnsOnCall == nil {
		fake.setOrganizationDefaultIsolationSegmentReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.setOrganizationDefaultIsolationSegmentReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getIsolationSegmentByNameMutex.RLock()
	defer fake.getIsolationSegmentByNameMutex.RUnlock()
	fake.setOrganizationDefaultIsolationSegmentMutex.RLock()
	defer fake.setOrganizationDefaultIsolationSegmentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetOrgDefaultIsolationSegmentActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.SetOrgDefaultIsolationSegmentActor = new(FakeSetOrgDefaultIsolationSegmentActor)
