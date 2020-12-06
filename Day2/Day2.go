package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	minFreq int64
	maxFreq int64
	keyChar string
	pw      string
}

func findValidPasswords(fileName string) (int, int, error) {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Failed opening %s file: %s", fileName, err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	validPasswordsCount1 := 0
	validPasswordsCount2 := 0
	var charCount int64
	var interpretedPassword password

	for scanner.Scan() {
		charCount = 0
		passwordString := scanner.Text()
		splitMin := strings.Split(passwordString, "-")
		interpretedPassword.minFreq, err = strconv.ParseInt(splitMin[0], 10, 32)
		if err != nil {
			return 0, 0, err
		}
		splitMax := strings.Split(splitMin[1], " ")
		interpretedPassword.maxFreq, err = strconv.ParseInt(splitMax[0], 10, 32)
		if err != nil {
			return 0, 0, err
		}
		interpretedPassword.keyChar = string(splitMax[1][0])
		interpretedPassword.pw = splitMax[2]
		for i := 0; i < len(interpretedPassword.pw); i++ {
			if interpretedPassword.keyChar == string(interpretedPassword.pw[i]) {
				charCount++
			}
		}
		if (charCount <= interpretedPassword.maxFreq) && (charCount >= interpretedPassword.minFreq) {
			validPasswordsCount1++
		}

		case1 := (string(interpretedPassword.pw[interpretedPassword.minFreq-1]) == interpretedPassword.keyChar)
		case2 := (string(interpretedPassword.pw[interpretedPassword.maxFreq-1]) == interpretedPassword.keyChar)
		if !(case1 && case2) && (case1 || case2) {
			validPasswordsCount2++
		}
	}
	file.Close()
	return validPasswordsCount1, validPasswordsCount2, nil
}

func main() {
	var inputFile = flag.String("i", "", "Input")
	flag.Parse()
	if *inputFile == "" {
		log.Fatal(fmt.Errorf("No file specified"))
	}
	answer1, answer2, err := findValidPasswords(*inputFile)
	if err != nil {
		log.Fatal("Could not produce an answer: %s", err)
	}
	fmt.Printf("Answer 1: %v\nAnswer 2: %v\n", answer1, answer2)
}
