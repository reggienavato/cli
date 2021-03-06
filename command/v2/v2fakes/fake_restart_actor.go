// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeRestartActor struct {
	GetApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.ApplicationSummary, v2action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}
	RestartApplicationStub        func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan v2action.ApplicationState, <-chan string, <-chan error)
	restartApplicationMutex       sync.RWMutex
	restartApplicationArgsForCall []struct {
		app    v2action.Application
		client v2action.NOAAClient
		config v2action.Config
	}
	restartApplicationReturns struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan v2action.ApplicationState
		result4 <-chan string
		result5 <-chan error
	}
	restartApplicationReturnsOnCall map[int]struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan v2action.ApplicationState
		result4 <-chan string
		result5 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRestartActor) GetApplicationByNameAndSpace(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeRestartActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeRestartActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].name, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeRestartActor) GetApplicationByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRestartActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRestartActor) GetApplicationSummaryByNameAndSpace(name string, spaceGUID string) (v2action.ApplicationSummary, v2action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationSummaryByNameAndSpaceReturns.result1, fake.getApplicationSummaryByNameAndSpaceReturns.result2, fake.getApplicationSummaryByNameAndSpaceReturns.result3
}

func (fake *FakeRestartActor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeRestartActor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].name, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeRestartActor) GetApplicationSummaryByNameAndSpaceReturns(result1 v2action.ApplicationSummary, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRestartActor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v2action.ApplicationSummary, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.ApplicationSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRestartActor) RestartApplication(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan v2action.ApplicationState, <-chan string, <-chan error) {
	fake.restartApplicationMutex.Lock()
	ret, specificReturn := fake.restartApplicationReturnsOnCall[len(fake.restartApplicationArgsForCall)]
	fake.restartApplicationArgsForCall = append(fake.restartApplicationArgsForCall, struct {
		app    v2action.Application
		client v2action.NOAAClient
		config v2action.Config
	}{app, client, config})
	fake.recordInvocation("RestartApplication", []interface{}{app, client, config})
	fake.restartApplicationMutex.Unlock()
	if fake.RestartApplicationStub != nil {
		return fake.RestartApplicationStub(app, client, config)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4, ret.result5
	}
	return fake.restartApplicationReturns.result1, fake.restartApplicationReturns.result2, fake.restartApplicationReturns.result3, fake.restartApplicationReturns.result4, fake.restartApplicationReturns.result5
}

func (fake *FakeRestartActor) RestartApplicationCallCount() int {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return len(fake.restartApplicationArgsForCall)
}

func (fake *FakeRestartActor) RestartApplicationArgsForCall(i int) (v2action.Application, v2action.NOAAClient, v2action.Config) {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return fake.restartApplicationArgsForCall[i].app, fake.restartApplicationArgsForCall[i].client, fake.restartApplicationArgsForCall[i].config
}

func (fake *FakeRestartActor) RestartApplicationReturns(result1 <-chan *v2action.LogMessage, result2 <-chan error, result3 <-chan v2action.ApplicationState, result4 <-chan string, result5 <-chan error) {
	fake.RestartApplicationStub = nil
	fake.restartApplicationReturns = struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan v2action.ApplicationState
		result4 <-chan string
		result5 <-chan error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeRestartActor) RestartApplicationReturnsOnCall(i int, result1 <-chan *v2action.LogMessage, result2 <-chan error, result3 <-chan v2action.ApplicationState, result4 <-chan string, result5 <-chan error) {
	fake.RestartApplicationStub = nil
	if fake.restartApplicationReturnsOnCall == nil {
		fake.restartApplicationReturnsOnCall = make(map[int]struct {
			result1 <-chan *v2action.LogMessage
			result2 <-chan error
			result3 <-chan v2action.ApplicationState
			result4 <-chan string
			result5 <-chan error
		})
	}
	fake.restartApplicationReturnsOnCall[i] = struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan v2action.ApplicationState
		result4 <-chan string
		result5 <-chan error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeRestartActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRestartActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.RestartActor = new(FakeRestartActor)
