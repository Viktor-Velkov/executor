// Code generated by counterfeiter. DO NOT EDIT.
package containerstorefakes

import (
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/executor/depot/containerstore"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
)

type FakeCredentialHandler struct {
	CloseStub        func(containerstore.Credential, executor.Container) error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 containerstore.Credential
		arg2 executor.Container
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CreateDirStub        func(lager.Logger, executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error)
	createDirMutex       sync.RWMutex
	createDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	createDirReturns struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}
	createDirReturnsOnCall map[int]struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}
	RemoveDirStub        func(lager.Logger, executor.Container) error
	removeDirMutex       sync.RWMutex
	removeDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	removeDirReturns struct {
		result1 error
	}
	removeDirReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(containerstore.Credential, executor.Container) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 containerstore.Credential
		arg2 executor.Container
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredentialHandler) Close(arg1 containerstore.Credential, arg2 executor.Container) error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 containerstore.Credential
		arg2 executor.Container
	}{arg1, arg2})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{arg1, arg2})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCredentialHandler) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeCredentialHandler) CloseCalls(stub func(containerstore.Credential, executor.Container) error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeCredentialHandler) CloseArgsForCall(i int) (containerstore.Credential, executor.Container) {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	argsForCall := fake.closeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCredentialHandler) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) CreateDir(arg1 lager.Logger, arg2 executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error) {
	fake.createDirMutex.Lock()
	ret, specificReturn := fake.createDirReturnsOnCall[len(fake.createDirArgsForCall)]
	fake.createDirArgsForCall = append(fake.createDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	stub := fake.CreateDirStub
	fakeReturns := fake.createDirReturns
	fake.recordInvocation("CreateDir", []interface{}{arg1, arg2})
	fake.createDirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCredentialHandler) CreateDirCallCount() int {
	fake.createDirMutex.RLock()
	defer fake.createDirMutex.RUnlock()
	return len(fake.createDirArgsForCall)
}

func (fake *FakeCredentialHandler) CreateDirCalls(stub func(lager.Logger, executor.Container) ([]garden.BindMount, []executor.EnvironmentVariable, error)) {
	fake.createDirMutex.Lock()
	defer fake.createDirMutex.Unlock()
	fake.CreateDirStub = stub
}

func (fake *FakeCredentialHandler) CreateDirArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.createDirMutex.RLock()
	defer fake.createDirMutex.RUnlock()
	argsForCall := fake.createDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCredentialHandler) CreateDirReturns(result1 []garden.BindMount, result2 []executor.EnvironmentVariable, result3 error) {
	fake.createDirMutex.Lock()
	defer fake.createDirMutex.Unlock()
	fake.CreateDirStub = nil
	fake.createDirReturns = struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCredentialHandler) CreateDirReturnsOnCall(i int, result1 []garden.BindMount, result2 []executor.EnvironmentVariable, result3 error) {
	fake.createDirMutex.Lock()
	defer fake.createDirMutex.Unlock()
	fake.CreateDirStub = nil
	if fake.createDirReturnsOnCall == nil {
		fake.createDirReturnsOnCall = make(map[int]struct {
			result1 []garden.BindMount
			result2 []executor.EnvironmentVariable
			result3 error
		})
	}
	fake.createDirReturnsOnCall[i] = struct {
		result1 []garden.BindMount
		result2 []executor.EnvironmentVariable
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCredentialHandler) RemoveDir(arg1 lager.Logger, arg2 executor.Container) error {
	fake.removeDirMutex.Lock()
	ret, specificReturn := fake.removeDirReturnsOnCall[len(fake.removeDirArgsForCall)]
	fake.removeDirArgsForCall = append(fake.removeDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	stub := fake.RemoveDirStub
	fakeReturns := fake.removeDirReturns
	fake.recordInvocation("RemoveDir", []interface{}{arg1, arg2})
	fake.removeDirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCredentialHandler) RemoveDirCallCount() int {
	fake.removeDirMutex.RLock()
	defer fake.removeDirMutex.RUnlock()
	return len(fake.removeDirArgsForCall)
}

func (fake *FakeCredentialHandler) RemoveDirCalls(stub func(lager.Logger, executor.Container) error) {
	fake.removeDirMutex.Lock()
	defer fake.removeDirMutex.Unlock()
	fake.RemoveDirStub = stub
}

func (fake *FakeCredentialHandler) RemoveDirArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.removeDirMutex.RLock()
	defer fake.removeDirMutex.RUnlock()
	argsForCall := fake.removeDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCredentialHandler) RemoveDirReturns(result1 error) {
	fake.removeDirMutex.Lock()
	defer fake.removeDirMutex.Unlock()
	fake.RemoveDirStub = nil
	fake.removeDirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) RemoveDirReturnsOnCall(i int, result1 error) {
	fake.removeDirMutex.Lock()
	defer fake.removeDirMutex.Unlock()
	fake.RemoveDirStub = nil
	if fake.removeDirReturnsOnCall == nil {
		fake.removeDirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeDirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) Update(arg1 containerstore.Credential, arg2 executor.Container) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 containerstore.Credential
		arg2 executor.Container
	}{arg1, arg2})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCredentialHandler) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeCredentialHandler) UpdateCalls(stub func(containerstore.Credential, executor.Container) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeCredentialHandler) UpdateArgsForCall(i int) (containerstore.Credential, executor.Container) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCredentialHandler) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredentialHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.createDirMutex.RLock()
	defer fake.createDirMutex.RUnlock()
	fake.removeDirMutex.RLock()
	defer fake.removeDirMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredentialHandler) recordInvocation(key string, args []interface{}) {
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

var _ containerstore.CredentialHandler = new(FakeCredentialHandler)
