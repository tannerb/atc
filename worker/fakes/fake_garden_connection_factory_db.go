// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeGardenConnectionFactoryDB struct {
	GetWorkerStub        func(string) (db.WorkerInfo, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		arg1 string
	}
	getWorkerReturns struct {
		result1 db.WorkerInfo
		result2 bool
		result3 error
	}
}

func (fake *FakeGardenConnectionFactoryDB) GetWorker(arg1 string) (db.WorkerInfo, bool, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(arg1)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2, fake.getWorkerReturns.result3
	}
}

func (fake *FakeGardenConnectionFactoryDB) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeGardenConnectionFactoryDB) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].arg1
}

func (fake *FakeGardenConnectionFactoryDB) GetWorkerReturns(result1 db.WorkerInfo, result2 bool, result3 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 db.WorkerInfo
		result2 bool
		result3 error
	}{result1, result2, result3}
}

var _ worker.GardenConnectionFactoryDB = new(FakeGardenConnectionFactoryDB)