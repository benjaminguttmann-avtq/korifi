// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/korifi/api/handlers"
)

type RequestValidator struct {
	DecodeAndValidateJSONPayloadStub        func(*http.Request, interface{}) error
	decodeAndValidateJSONPayloadMutex       sync.RWMutex
	decodeAndValidateJSONPayloadArgsForCall []struct {
		arg1 *http.Request
		arg2 interface{}
	}
	decodeAndValidateJSONPayloadReturns struct {
		result1 error
	}
	decodeAndValidateJSONPayloadReturnsOnCall map[int]struct {
		result1 error
	}
	DecodeAndValidateURLValuesStub        func(*http.Request, handlers.KeyedPayload) error
	decodeAndValidateURLValuesMutex       sync.RWMutex
	decodeAndValidateURLValuesArgsForCall []struct {
		arg1 *http.Request
		arg2 handlers.KeyedPayload
	}
	decodeAndValidateURLValuesReturns struct {
		result1 error
	}
	decodeAndValidateURLValuesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RequestValidator) DecodeAndValidateJSONPayload(arg1 *http.Request, arg2 interface{}) error {
	fake.decodeAndValidateJSONPayloadMutex.Lock()
	ret, specificReturn := fake.decodeAndValidateJSONPayloadReturnsOnCall[len(fake.decodeAndValidateJSONPayloadArgsForCall)]
	fake.decodeAndValidateJSONPayloadArgsForCall = append(fake.decodeAndValidateJSONPayloadArgsForCall, struct {
		arg1 *http.Request
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.DecodeAndValidateJSONPayloadStub
	fakeReturns := fake.decodeAndValidateJSONPayloadReturns
	fake.recordInvocation("DecodeAndValidateJSONPayload", []interface{}{arg1, arg2})
	fake.decodeAndValidateJSONPayloadMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *RequestValidator) DecodeAndValidateJSONPayloadCallCount() int {
	fake.decodeAndValidateJSONPayloadMutex.RLock()
	defer fake.decodeAndValidateJSONPayloadMutex.RUnlock()
	return len(fake.decodeAndValidateJSONPayloadArgsForCall)
}

func (fake *RequestValidator) DecodeAndValidateJSONPayloadCalls(stub func(*http.Request, interface{}) error) {
	fake.decodeAndValidateJSONPayloadMutex.Lock()
	defer fake.decodeAndValidateJSONPayloadMutex.Unlock()
	fake.DecodeAndValidateJSONPayloadStub = stub
}

func (fake *RequestValidator) DecodeAndValidateJSONPayloadArgsForCall(i int) (*http.Request, interface{}) {
	fake.decodeAndValidateJSONPayloadMutex.RLock()
	defer fake.decodeAndValidateJSONPayloadMutex.RUnlock()
	argsForCall := fake.decodeAndValidateJSONPayloadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *RequestValidator) DecodeAndValidateJSONPayloadReturns(result1 error) {
	fake.decodeAndValidateJSONPayloadMutex.Lock()
	defer fake.decodeAndValidateJSONPayloadMutex.Unlock()
	fake.DecodeAndValidateJSONPayloadStub = nil
	fake.decodeAndValidateJSONPayloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *RequestValidator) DecodeAndValidateJSONPayloadReturnsOnCall(i int, result1 error) {
	fake.decodeAndValidateJSONPayloadMutex.Lock()
	defer fake.decodeAndValidateJSONPayloadMutex.Unlock()
	fake.DecodeAndValidateJSONPayloadStub = nil
	if fake.decodeAndValidateJSONPayloadReturnsOnCall == nil {
		fake.decodeAndValidateJSONPayloadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.decodeAndValidateJSONPayloadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *RequestValidator) DecodeAndValidateURLValues(arg1 *http.Request, arg2 handlers.KeyedPayload) error {
	fake.decodeAndValidateURLValuesMutex.Lock()
	ret, specificReturn := fake.decodeAndValidateURLValuesReturnsOnCall[len(fake.decodeAndValidateURLValuesArgsForCall)]
	fake.decodeAndValidateURLValuesArgsForCall = append(fake.decodeAndValidateURLValuesArgsForCall, struct {
		arg1 *http.Request
		arg2 handlers.KeyedPayload
	}{arg1, arg2})
	stub := fake.DecodeAndValidateURLValuesStub
	fakeReturns := fake.decodeAndValidateURLValuesReturns
	fake.recordInvocation("DecodeAndValidateURLValues", []interface{}{arg1, arg2})
	fake.decodeAndValidateURLValuesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *RequestValidator) DecodeAndValidateURLValuesCallCount() int {
	fake.decodeAndValidateURLValuesMutex.RLock()
	defer fake.decodeAndValidateURLValuesMutex.RUnlock()
	return len(fake.decodeAndValidateURLValuesArgsForCall)
}

func (fake *RequestValidator) DecodeAndValidateURLValuesCalls(stub func(*http.Request, handlers.KeyedPayload) error) {
	fake.decodeAndValidateURLValuesMutex.Lock()
	defer fake.decodeAndValidateURLValuesMutex.Unlock()
	fake.DecodeAndValidateURLValuesStub = stub
}

func (fake *RequestValidator) DecodeAndValidateURLValuesArgsForCall(i int) (*http.Request, handlers.KeyedPayload) {
	fake.decodeAndValidateURLValuesMutex.RLock()
	defer fake.decodeAndValidateURLValuesMutex.RUnlock()
	argsForCall := fake.decodeAndValidateURLValuesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *RequestValidator) DecodeAndValidateURLValuesReturns(result1 error) {
	fake.decodeAndValidateURLValuesMutex.Lock()
	defer fake.decodeAndValidateURLValuesMutex.Unlock()
	fake.DecodeAndValidateURLValuesStub = nil
	fake.decodeAndValidateURLValuesReturns = struct {
		result1 error
	}{result1}
}

func (fake *RequestValidator) DecodeAndValidateURLValuesReturnsOnCall(i int, result1 error) {
	fake.decodeAndValidateURLValuesMutex.Lock()
	defer fake.decodeAndValidateURLValuesMutex.Unlock()
	fake.DecodeAndValidateURLValuesStub = nil
	if fake.decodeAndValidateURLValuesReturnsOnCall == nil {
		fake.decodeAndValidateURLValuesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.decodeAndValidateURLValuesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *RequestValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decodeAndValidateJSONPayloadMutex.RLock()
	defer fake.decodeAndValidateJSONPayloadMutex.RUnlock()
	fake.decodeAndValidateURLValuesMutex.RLock()
	defer fake.decodeAndValidateURLValuesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RequestValidator) recordInvocation(key string, args []interface{}) {
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

var _ handlers.RequestValidator = new(RequestValidator)
