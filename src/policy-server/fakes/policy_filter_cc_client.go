// This file was generated by counterfeiter
package fakes

import "sync"

type PolicyFilterCCClient struct {
	GetAppSpacesStub        func(token string, appGUIDs []string) (map[string]string, error)
	getAppSpacesMutex       sync.RWMutex
	getAppSpacesArgsForCall []struct {
		token    string
		appGUIDs []string
	}
	getAppSpacesReturns struct {
		result1 map[string]string
		result2 error
	}
	GetUserSpacesStub        func(token, userGUID string) (map[string]struct{}, error)
	getUserSpacesMutex       sync.RWMutex
	getUserSpacesArgsForCall []struct {
		token    string
		userGUID string
	}
	getUserSpacesReturns struct {
		result1 map[string]struct{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyFilterCCClient) GetAppSpaces(token string, appGUIDs []string) (map[string]string, error) {
	var appGUIDsCopy []string
	if appGUIDs != nil {
		appGUIDsCopy = make([]string, len(appGUIDs))
		copy(appGUIDsCopy, appGUIDs)
	}
	fake.getAppSpacesMutex.Lock()
	fake.getAppSpacesArgsForCall = append(fake.getAppSpacesArgsForCall, struct {
		token    string
		appGUIDs []string
	}{token, appGUIDsCopy})
	fake.recordInvocation("GetAppSpaces", []interface{}{token, appGUIDsCopy})
	fake.getAppSpacesMutex.Unlock()
	if fake.GetAppSpacesStub != nil {
		return fake.GetAppSpacesStub(token, appGUIDs)
	} else {
		return fake.getAppSpacesReturns.result1, fake.getAppSpacesReturns.result2
	}
}

func (fake *PolicyFilterCCClient) GetAppSpacesCallCount() int {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return len(fake.getAppSpacesArgsForCall)
}

func (fake *PolicyFilterCCClient) GetAppSpacesArgsForCall(i int) (string, []string) {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return fake.getAppSpacesArgsForCall[i].token, fake.getAppSpacesArgsForCall[i].appGUIDs
}

func (fake *PolicyFilterCCClient) GetAppSpacesReturns(result1 map[string]string, result2 error) {
	fake.GetAppSpacesStub = nil
	fake.getAppSpacesReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *PolicyFilterCCClient) GetUserSpaces(token string, userGUID string) (map[string]struct{}, error) {
	fake.getUserSpacesMutex.Lock()
	fake.getUserSpacesArgsForCall = append(fake.getUserSpacesArgsForCall, struct {
		token    string
		userGUID string
	}{token, userGUID})
	fake.recordInvocation("GetUserSpaces", []interface{}{token, userGUID})
	fake.getUserSpacesMutex.Unlock()
	if fake.GetUserSpacesStub != nil {
		return fake.GetUserSpacesStub(token, userGUID)
	} else {
		return fake.getUserSpacesReturns.result1, fake.getUserSpacesReturns.result2
	}
}

func (fake *PolicyFilterCCClient) GetUserSpacesCallCount() int {
	fake.getUserSpacesMutex.RLock()
	defer fake.getUserSpacesMutex.RUnlock()
	return len(fake.getUserSpacesArgsForCall)
}

func (fake *PolicyFilterCCClient) GetUserSpacesArgsForCall(i int) (string, string) {
	fake.getUserSpacesMutex.RLock()
	defer fake.getUserSpacesMutex.RUnlock()
	return fake.getUserSpacesArgsForCall[i].token, fake.getUserSpacesArgsForCall[i].userGUID
}

func (fake *PolicyFilterCCClient) GetUserSpacesReturns(result1 map[string]struct{}, result2 error) {
	fake.GetUserSpacesStub = nil
	fake.getUserSpacesReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *PolicyFilterCCClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	fake.getUserSpacesMutex.RLock()
	defer fake.getUserSpacesMutex.RUnlock()
	return fake.invocations
}

func (fake *PolicyFilterCCClient) recordInvocation(key string, args []interface{}) {
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
