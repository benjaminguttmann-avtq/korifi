// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
)

type CFRoleRepository struct {
	CreateRoleStub        func(context.Context, repositories.RoleRecord) (repositories.RoleRecord, error)
	createRoleMutex       sync.RWMutex
	createRoleArgsForCall []struct {
		arg1 context.Context
		arg2 repositories.RoleRecord
	}
	createRoleReturns struct {
		result1 repositories.RoleRecord
		result2 error
	}
	createRoleReturnsOnCall map[int]struct {
		result1 repositories.RoleRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFRoleRepository) CreateRole(arg1 context.Context, arg2 repositories.RoleRecord) (repositories.RoleRecord, error) {
	fake.createRoleMutex.Lock()
	ret, specificReturn := fake.createRoleReturnsOnCall[len(fake.createRoleArgsForCall)]
	fake.createRoleArgsForCall = append(fake.createRoleArgsForCall, struct {
		arg1 context.Context
		arg2 repositories.RoleRecord
	}{arg1, arg2})
	stub := fake.CreateRoleStub
	fakeReturns := fake.createRoleReturns
	fake.recordInvocation("CreateRole", []interface{}{arg1, arg2})
	fake.createRoleMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFRoleRepository) CreateRoleCallCount() int {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	return len(fake.createRoleArgsForCall)
}

func (fake *CFRoleRepository) CreateRoleCalls(stub func(context.Context, repositories.RoleRecord) (repositories.RoleRecord, error)) {
	fake.createRoleMutex.Lock()
	defer fake.createRoleMutex.Unlock()
	fake.CreateRoleStub = stub
}

func (fake *CFRoleRepository) CreateRoleArgsForCall(i int) (context.Context, repositories.RoleRecord) {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	argsForCall := fake.createRoleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CFRoleRepository) CreateRoleReturns(result1 repositories.RoleRecord, result2 error) {
	fake.createRoleMutex.Lock()
	defer fake.createRoleMutex.Unlock()
	fake.CreateRoleStub = nil
	fake.createRoleReturns = struct {
		result1 repositories.RoleRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRoleRepository) CreateRoleReturnsOnCall(i int, result1 repositories.RoleRecord, result2 error) {
	fake.createRoleMutex.Lock()
	defer fake.createRoleMutex.Unlock()
	fake.CreateRoleStub = nil
	if fake.createRoleReturnsOnCall == nil {
		fake.createRoleReturnsOnCall = make(map[int]struct {
			result1 repositories.RoleRecord
			result2 error
		})
	}
	fake.createRoleReturnsOnCall[i] = struct {
		result1 repositories.RoleRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRoleRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFRoleRepository) recordInvocation(key string, args []interface{}) {
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

var _ apis.CFRoleRepository = new(CFRoleRepository)
