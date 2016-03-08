// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/exec"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeFactory struct {
	GetStub        func(lager.Logger, exec.StepMetadata, exec.SourceName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.Version, atc.ResourceTypes) exec.StepFactory
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  exec.SourceName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  atc.Params
		arg10 atc.Version
		arg11 atc.ResourceTypes
	}
	getReturns struct {
		result1 exec.StepFactory
	}
	PutStub        func(lager.Logger, exec.StepMetadata, worker.Identifier, worker.Metadata, exec.PutDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.ResourceTypes) exec.StepFactory
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 lager.Logger
		arg2 exec.StepMetadata
		arg3 worker.Identifier
		arg4 worker.Metadata
		arg5 exec.PutDelegate
		arg6 atc.ResourceConfig
		arg7 atc.Tags
		arg8 atc.Params
		arg9 atc.ResourceTypes
	}
	putReturns struct {
		result1 exec.StepFactory
	}
	DependentGetStub        func(lager.Logger, exec.StepMetadata, exec.SourceName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.ResourceTypes) exec.StepFactory
	dependentGetMutex       sync.RWMutex
	dependentGetArgsForCall []struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  exec.SourceName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  atc.Params
		arg10 atc.ResourceTypes
	}
	dependentGetReturns struct {
		result1 exec.StepFactory
	}
	TaskStub        func(lager.Logger, exec.SourceName, worker.Identifier, worker.Metadata, exec.TaskDelegate, exec.Privileged, atc.Tags, exec.TaskConfigSource, atc.ResourceTypes, map[string]string, map[string]string) exec.StepFactory
	taskMutex       sync.RWMutex
	taskArgsForCall []struct {
		arg1  lager.Logger
		arg2  exec.SourceName
		arg3  worker.Identifier
		arg4  worker.Metadata
		arg5  exec.TaskDelegate
		arg6  exec.Privileged
		arg7  atc.Tags
		arg8  exec.TaskConfigSource
		arg9  atc.ResourceTypes
		arg10 map[string]string
		arg11 map[string]string
	}
	taskReturns struct {
		result1 exec.StepFactory
	}
}

func (fake *FakeFactory) Get(arg1 lager.Logger, arg2 exec.StepMetadata, arg3 exec.SourceName, arg4 worker.Identifier, arg5 worker.Metadata, arg6 exec.GetDelegate, arg7 atc.ResourceConfig, arg8 atc.Tags, arg9 atc.Params, arg10 atc.Version, arg11 atc.ResourceTypes) exec.StepFactory {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  exec.SourceName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  atc.Params
		arg10 atc.Version
		arg11 atc.ResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeFactory) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFactory) GetArgsForCall(i int) (lager.Logger, exec.StepMetadata, exec.SourceName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.Version, atc.ResourceTypes) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3, fake.getArgsForCall[i].arg4, fake.getArgsForCall[i].arg5, fake.getArgsForCall[i].arg6, fake.getArgsForCall[i].arg7, fake.getArgsForCall[i].arg8, fake.getArgsForCall[i].arg9, fake.getArgsForCall[i].arg10, fake.getArgsForCall[i].arg11
}

func (fake *FakeFactory) GetReturns(result1 exec.StepFactory) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Put(arg1 lager.Logger, arg2 exec.StepMetadata, arg3 worker.Identifier, arg4 worker.Metadata, arg5 exec.PutDelegate, arg6 atc.ResourceConfig, arg7 atc.Tags, arg8 atc.Params, arg9 atc.ResourceTypes) exec.StepFactory {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 lager.Logger
		arg2 exec.StepMetadata
		arg3 worker.Identifier
		arg4 worker.Metadata
		arg5 exec.PutDelegate
		arg6 atc.ResourceConfig
		arg7 atc.Tags
		arg8 atc.Params
		arg9 atc.ResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	} else {
		return fake.putReturns.result1
	}
}

func (fake *FakeFactory) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeFactory) PutArgsForCall(i int) (lager.Logger, exec.StepMetadata, worker.Identifier, worker.Metadata, exec.PutDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.ResourceTypes) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3, fake.putArgsForCall[i].arg4, fake.putArgsForCall[i].arg5, fake.putArgsForCall[i].arg6, fake.putArgsForCall[i].arg7, fake.putArgsForCall[i].arg8, fake.putArgsForCall[i].arg9
}

func (fake *FakeFactory) PutReturns(result1 exec.StepFactory) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) DependentGet(arg1 lager.Logger, arg2 exec.StepMetadata, arg3 exec.SourceName, arg4 worker.Identifier, arg5 worker.Metadata, arg6 exec.GetDelegate, arg7 atc.ResourceConfig, arg8 atc.Tags, arg9 atc.Params, arg10 atc.ResourceTypes) exec.StepFactory {
	fake.dependentGetMutex.Lock()
	fake.dependentGetArgsForCall = append(fake.dependentGetArgsForCall, struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  exec.SourceName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  atc.Params
		arg10 atc.ResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.dependentGetMutex.Unlock()
	if fake.DependentGetStub != nil {
		return fake.DependentGetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	} else {
		return fake.dependentGetReturns.result1
	}
}

func (fake *FakeFactory) DependentGetCallCount() int {
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	return len(fake.dependentGetArgsForCall)
}

func (fake *FakeFactory) DependentGetArgsForCall(i int) (lager.Logger, exec.StepMetadata, exec.SourceName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.ResourceTypes) {
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	return fake.dependentGetArgsForCall[i].arg1, fake.dependentGetArgsForCall[i].arg2, fake.dependentGetArgsForCall[i].arg3, fake.dependentGetArgsForCall[i].arg4, fake.dependentGetArgsForCall[i].arg5, fake.dependentGetArgsForCall[i].arg6, fake.dependentGetArgsForCall[i].arg7, fake.dependentGetArgsForCall[i].arg8, fake.dependentGetArgsForCall[i].arg9, fake.dependentGetArgsForCall[i].arg10
}

func (fake *FakeFactory) DependentGetReturns(result1 exec.StepFactory) {
	fake.DependentGetStub = nil
	fake.dependentGetReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Task(arg1 lager.Logger, arg2 exec.SourceName, arg3 worker.Identifier, arg4 worker.Metadata, arg5 exec.TaskDelegate, arg6 exec.Privileged, arg7 atc.Tags, arg8 exec.TaskConfigSource, arg9 atc.ResourceTypes, arg10 map[string]string, arg11 map[string]string) exec.StepFactory {
	fake.taskMutex.Lock()
	fake.taskArgsForCall = append(fake.taskArgsForCall, struct {
		arg1  lager.Logger
		arg2  exec.SourceName
		arg3  worker.Identifier
		arg4  worker.Metadata
		arg5  exec.TaskDelegate
		arg6  exec.Privileged
		arg7  atc.Tags
		arg8  exec.TaskConfigSource
		arg9  atc.ResourceTypes
		arg10 map[string]string
		arg11 map[string]string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11})
	fake.taskMutex.Unlock()
	if fake.TaskStub != nil {
		return fake.TaskStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	} else {
		return fake.taskReturns.result1
	}
}

func (fake *FakeFactory) TaskCallCount() int {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return len(fake.taskArgsForCall)
}

func (fake *FakeFactory) TaskArgsForCall(i int) (lager.Logger, exec.SourceName, worker.Identifier, worker.Metadata, exec.TaskDelegate, exec.Privileged, atc.Tags, exec.TaskConfigSource, atc.ResourceTypes, map[string]string, map[string]string) {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return fake.taskArgsForCall[i].arg1, fake.taskArgsForCall[i].arg2, fake.taskArgsForCall[i].arg3, fake.taskArgsForCall[i].arg4, fake.taskArgsForCall[i].arg5, fake.taskArgsForCall[i].arg6, fake.taskArgsForCall[i].arg7, fake.taskArgsForCall[i].arg8, fake.taskArgsForCall[i].arg9, fake.taskArgsForCall[i].arg10, fake.taskArgsForCall[i].arg11
}

func (fake *FakeFactory) TaskReturns(result1 exec.StepFactory) {
	fake.TaskStub = nil
	fake.taskReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

var _ exec.Factory = new(FakeFactory)
