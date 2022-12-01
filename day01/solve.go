package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readFile(fileName string) {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var total = 0
	var candidates []int
	for scanner.Scan() {
		if scanner.Text() != "" {
			var cal, _ = strconv.Atoi(scanner.Text())
			total += cal
		} else {
			candidates = append(candidates, total)
			total = 0
		}
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	fmt.Println(candidates)
	fmt.Println(70116 + 68706 + 67760)
}

func main() {
	readFile("input.txt")
}
