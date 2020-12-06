package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type arrayFlag []int

func (i *arrayFlag) String() string {
	return ""
}

func (i *arrayFlag) Set(value string) error {
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*i = append(*i, valueInt)
	return nil
}

var slopesHorizontal arrayFlag
var slopesVertical arrayFlag

func findTrees(fileName string, slope [2]int) (int, error) {

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
		currentCoords[0] = (currentCoords[0] + slope[0]) % width
		currentCoords[1] = (currentCoords[1] + slope[1])
		if string(forestMap[currentCoords[1]][currentCoords[0]]) == "#" {
			treeCount++

		}
	}
	return treeCount, nil

}

func main() {
	var inputFile = flag.String("i", "", "Input")
	flag.Var(&slopesHorizontal, "h", "Slope (horizontal part).")
	flag.Var(&slopesVertical, "v", "Slope (vertical part).")

	flag.Parse()
	if *inputFile == "" {
		log.Fatal(fmt.Errorf("No file specified"))
	}
	if len(slopesHorizontal) != len(slopesVertical) {
		log.Fatal(fmt.Errorf("The number of horizontal and vertical parts (of the slopes) must be equal."))
	}
	product := 1
	for i, h := range slopesHorizontal {
		answer, err := findTrees(*inputFile, [2]int{h, slopesVertical[i]})
		if err != nil {
			log.Fatal("Could not produce an answer: %s", err)
		}
		product *= answer
		fmt.Printf("Answer %d: %v\n", i+1, answer)
	}
	fmt.Printf("Product: %v\n", product)
}
