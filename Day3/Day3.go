package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func findTrees(fileName string) (int, error) {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Failed opening %s file: %s", fileName, err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var forestMap []string

	for scanner.Scan() {
		forestMap = append(forestMap, scanner.Text())
	}

	file.Close()

	width := len(forestMap[0])
	currentCoords := [2]int{0, 0}
	treeCount := 0
	for currentCoords[1] < len(forestMap)-1 {
		currentCoords[0] = (currentCoords[0] + 3) % width
		currentCoords[1] = (currentCoords[1] + 1)
		if string(forestMap[currentCoords[1]][currentCoords[0]]) == "#" {
			treeCount++
		}
	}
	return treeCount, nil

}

func main() {
	var inputFile = flag.String("i", "", "Input")
	flag.Parse()
	if *inputFile == "" {
		log.Fatal(fmt.Errorf("No file specified"))
	}
	answer, err := findTrees(*inputFile)
	if err != nil {
		log.Fatal("Could not produce an answer: %s", err)
	}
	fmt.Printf("Answer: %v\n", answer)
}
