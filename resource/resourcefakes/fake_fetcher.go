// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeFetcher struct {
	FetchStub        func(logger lager.Logger, session resource.Session, tags atc.Tags, teamID int, resourceTypes creds.VersionedResourceTypes, resourceInstance resource.ResourceInstance, metadata resource.Metadata, imageFetchingDelegate worker.ImageFetchingDelegate, signals <-chan os.Signal, ready chan<- struct{}) (resource.VersionedSource, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		logger                lager.Logger
		session               resource.Session
		tags                  atc.Tags
		teamID                int
		resourceTypes         creds.VersionedResourceTypes
		resourceInstance      resource.ResourceInstance
		metadata              resource.Metadata
		imageFetchingDelegate worker.ImageFetchingDelegate
		signals               <-chan os.Signal
		ready                 chan<- struct{}
	}
	fetchReturns struct {
		result1 resource.VersionedSource
		result2 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 resource.VersionedSource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetcher) Fetch(logger lager.Logger, session resource.Session, tags atc.Tags, teamID int, resourceTypes creds.VersionedResourceTypes, resourceInstance resource.ResourceInstance, metadata resource.Metadata, imageFetchingDelegate worker.ImageFetchingDelegate, signals <-chan os.Signal, ready chan<- struct{}) (resource.VersionedSource, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		logger                lager.Logger
		session               resource.Session
		tags                  atc.Tags
		teamID                int
		resourceTypes         creds.VersionedResourceTypes
		resourceInstance      resource.ResourceInstance
		metadata              resource.Metadata
		imageFetchingDelegate worker.ImageFetchingDelegate
		signals               <-chan os.Signal
		ready                 chan<- struct{}
	}{logger, session, tags, teamID, resourceTypes, resourceInstance, metadata, imageFetchingDelegate, signals, ready})
	fake.recordInvocation("Fetch", []interface{}{logger, session, tags, teamID, resourceTypes, resourceInstance, metadata, imageFetchingDelegate, signals, ready})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(logger, session, tags, teamID, resourceTypes, resourceInstance, metadata, imageFetchingDelegate, signals, ready)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchReturns.result1, fake.fetchReturns.result2
}

func (fake *FakeFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeFetcher) FetchArgsForCall(i int) (lager.Logger, resource.Session, atc.Tags, int, creds.VersionedResourceTypes, resource.ResourceInstance, resource.Metadata, worker.ImageFetchingDelegate, <-chan os.Signal, chan<- struct{}) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].logger, fake.fetchArgsForCall[i].session, fake.fetchArgsForCall[i].tags, fake.fetchArgsForCall[i].teamID, fake.fetchArgsForCall[i].resourceTypes, fake.fetchArgsForCall[i].resourceInstance, fake.fetchArgsForCall[i].metadata, fake.fetchArgsForCall[i].imageFetchingDelegate, fake.fetchArgsForCall[i].signals, fake.fetchArgsForCall[i].ready
}

func (fake *FakeFetcher) FetchReturns(result1 resource.VersionedSource, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) FetchReturnsOnCall(i int, result1 resource.VersionedSource, result2 error) {
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 resource.VersionedSource
			result2 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetcher) recordInvocation(key string, args []interface{}) {
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

var _ resource.Fetcher = new(FakeFetcher)
