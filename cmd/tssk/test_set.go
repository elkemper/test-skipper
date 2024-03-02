package main

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

func (testSet TestSet) add(testCodes []string) []string {
	var notProcessed []string
	for _, stringCode := range testCodes {
		testCode := TestCode(stringCode)
		if _, exists := testSet[testCode]; !exists {
			testSet[testCode] = struct{}{}
		} else {
			notProcessed = append(notProcessed, stringCode)
		}
	}
	return notProcessed
}

func (testSet TestSet) remove(testCodes []string) []string {
	var notProcessed []string
	for _, stringCode := range testCodes {
		testCode := TestCode(stringCode)
		if _, exists := testSet[testCode]; exists {
			delete(testSet, testCode)
		} else {
			notProcessed = append(notProcessed, stringCode)
		}
	}
	return notProcessed
}

func makeTestSet(testCodes []string) TestSet {
	testSet := make(TestSet, len(testCodes))
	for ind := range testCodes {
		testCode := TestCode(testCodes[ind])
		testSet[testCode] = struct{}{}
	}
	return testSet
}
