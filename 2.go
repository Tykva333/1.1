package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		
		for _, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return
	}

	sort.Ints(numbers)

	outFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer outFile.Close()

	for i, num := range numbers {
		if i > 0 {
			_, err = outFile.WriteString(" ")
			if err != nil {
				return
			}
		}
		_, err = outFile.WriteString(strconv.Itoa(num))
		if err != nil {
			return
		}
	}

	fmt.Println("Done")
}