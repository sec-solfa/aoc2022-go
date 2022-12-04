package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	key   string
	value int
}

func part1(scanner *bufio.Scanner) {
	var count int
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		a := strings.Split(arr[0], "-")
		b := strings.Split(arr[1], "-")
		a_start, _ := strconv.Atoi(a[0])
		a_end, _ := strconv.Atoi(a[1])
		b_start, _ := strconv.Atoi(b[0])
		b_end, _ := strconv.Atoi(b[1])
		if a_start <= b_start && b_end <= a_end {
			count += 1
		} else if b_start <= a_start && a_end <= b_end {
			count += 1
		}
	}
	fmt.Println(count)
}

func part2(scanner *bufio.Scanner) {
	var count int
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		a := strings.Split(arr[0], "-")
		b := strings.Split(arr[1], "-")
		a_start, _ := strconv.Atoi(a[0])
		a_end, _ := strconv.Atoi(a[1])
		b_start, _ := strconv.Atoi(b[0])
		b_end, _ := strconv.Atoi(b[1])
		if a_end < b_start || b_end < a_start {
			continue
		}
		count += 1
	}
	fmt.Println(count)
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	// part1(scanner)
	part2(scanner)
}
