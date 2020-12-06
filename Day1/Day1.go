package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func findSum2020(fileName string) (int, error) {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Failed opening %s file: %s", fileName, err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var intLines []int

	for scanner.Scan() {
		lineInt, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return 0, err
		}
		intLines = append(intLines, int(lineInt))
	}

	file.Close()

	sort.Ints(intLines)

	for i := 0; i < len(intLines); i++ {
		int1 := intLines[i]
		int2 := 2020 - int1
		if intLines[sort.SearchInts(intLines, int2)] == int2 {
			return int1 * int2, nil
		}
	}

	return 0, fmt.Errorf("Cannot find :(")

}

func main() {
	var inputFile = flag.String("i", "", "Input")
	flag.Parse()
	if *inputFile == "" {
		log.Fatal(fmt.Errorf("No file specified"))
	}
	answer, err := findSum2020(*inputFile)
	if err != nil {
		log.Fatal("Could not produce an answer: %s", err)
	}
	fmt.Printf("Answer: %v\n", answer)
}
