package main

import (
	"strings"
)

type TestSet map[TestCode]struct{}

func (testSet TestSet) toString() []string {
	testCodes := make([]string, len(testSet))
	ind := 0
	for key := range testSet {
		testCodes[ind] = string(key)
		ind++
	}
	return testCodes
}

func (testSet TestSet) add(testCodes []TestCode) []TestCode {
	var notProcessed []TestCode
	for _, testCode := range testCodes {
		if len(testCode) <= 0 {
			continue
		}
		if _, exists := testSet[testCode]; !exists {
			testSet[testCode] = struct{}{}
		} else {
			notProcessed = append(notProcessed, testCode)
		}
	}
	return notProcessed
}

func (testSet TestSet) remove(testCodes []TestCode) []TestCode {
	var notProcessed []TestCode
	for _, testCode := range testCodes {
		if _, exists := testSet[testCode]; exists {
			delete(testSet, testCode)
		} else {
			notProcessed = append(notProcessed, testCode)
		}
	}
	return notProcessed
}

func makeTestSet(testCodes []string) TestSet {
	testSet := make(TestSet, 0)
	for _, testCodeString := range testCodes {
		testCodeString = strings.TrimSpace(testCodeString)
		if len(testCodeString) > 0 {
			testCode := TestCode(testCodeString)
			testSet[testCode] = struct{}{}
		}
	}
	return testSet
}
