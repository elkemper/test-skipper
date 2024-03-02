package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const skippedFileName = "skipped.cfg"
const delimiter = ", "

type Command uint8
type FileFormat uint8
type TestCode string
type SkippedTests struct {
	Tests TestSet
}

const (
	PlainText FileFormat = 1 + iota
)

const (
	Skip Command = 1 + iota
	Unskip
)

func (command Command) Name() string {
	var name = "invalid"
	switch command {
	case Skip:
		name = "skip"
	case Unskip:
		name = "unskip"
	}
	return name
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		command, testCodes, err := parseArgs(args)
		if err != nil {
			log.Fatalf("Error:%s.", err)
			return
		}
		possible := possibleCommands(skippedFileName)
		if !slices.Contains(possible, command) {
			log.Fatalf("Provided command isn't possible.")
		}
		err = PerformAction(command, skippedFileName, testCodes)
		if err != nil {
			log.Fatalf("Error:%s", err)
			return
		}
	} else {
		log.Fatal("Error: Please specify a command. Supported commands are skip and unskip.")
		return
	}
}

func parseArgs(args []string) (Command, []string, error) {
	commandString := args[0]
	command, err := parseCommand(commandString)
	testCodes := args[1:]
	if err != nil {
		return command, testCodes, err
	}

	if len(testCodes) <= 0 {
		return command, testCodes, errors.New("invalid arguments. No test codes passed")
	}
	return command, testCodes, nil
}

func possibleCommands(skippedFileName string) []Command {
	if checkFileExists(skippedFileName) {
		return []Command{Skip, Unskip}
	}
	return []Command{Skip}
}

// PerformAction is the base function that receives all the arguments
// and decides what to do.
func PerformAction(command Command, skippedFileName string, testCodes []string) error {
	var skippedTests SkippedTests
	testSet := make(TestSet, 0)
	skippedTests.Tests = testSet
	var err error
	if !bypassFileReading(skippedFileName, command) {
		skippedTests, err = readSkippedTests(skippedFileName)
		if err != nil {
			return errors.New("error occurred in reading the file")
		}
	}

	var notProcessed []string
	var notProcessingReason string
	if command == Skip {
		notProcessed = skippedTests.Tests.add(testCodes)
		notProcessingReason = "Already existed."
	} else {
		notProcessed = skippedTests.Tests.remove(testCodes)
		notProcessingReason = "Didn't exist."
	}
	if len(notProcessed) > 0 {
		fmt.Printf("Couldn't %s: %s. Reason: %s\n", command.Name(), notProcessed, notProcessingReason)
	}
	fmt.Printf("Saving %d code(s): %s\n", len(skippedTests.Tests), skippedTests.Tests.toString())
	err = saveSkippedTests(skippedFileName, PlainText, skippedTests)
	if err != nil {
		return errors.New("error occurred in writing to file")
	}
	return nil
}

func parseCommand(commandString string) (Command, error) {
	var command Command
	switch commandString {
	case "skip":
		command = Skip
	case "unskip":
		command = Unskip
	default:
		return command, errors.New("supported commands are skip and unskip")
	}
	return command, nil
}

func checkFileExists(skippedFileName string) bool {
	_, err := os.Stat(skippedFileName)
	return !os.IsNotExist(err)
}

func bypassFileReading(skippedFileName string, command Command) bool {
	fileExists := checkFileExists(skippedFileName)
	return !fileExists && command == Skip
}

func saveSkippedTests(fileName string, format FileFormat, skippedTests SkippedTests) error {
	var testCodesSingleString string
	if format == PlainText {
		testCodesSingleString = strings.Join(skippedTests.Tests.toString(), delimiter)
	} else {
		return errors.New("only plain text format is supported")
	}
	
	err := os.WriteFile(skippedFileName, []byte(testCodesSingleString), 0644)
	if err != nil {
		return err
	}
	return nil
}

func readSkippedTests(fileName string) (SkippedTests, error) {
	var skippedTests SkippedTests
	contents, err := os.ReadFile(skippedFileName)
	if err != nil {
		return skippedTests, errors.New("error occurred in reading the file")
	}

	contentString := string(contents)
	var testSlice []string
	if len(strings.TrimSpace(contentString)) > 0 {
		testSlice = strings.Split(contentString, delimiter)
	}
	testSet := makeTestSet(testSlice)
	skippedTests.Tests = testSet
	return skippedTests, nil
}
