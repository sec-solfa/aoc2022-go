package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func genStack() [][]string {
	var slice [][]string
	fp, err := os.Open("stack.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		slice = append(slice, arr)
	}
	return slice
}

func part1(scanner *bufio.Scanner, stack [][]string) {
	for scanner.Scan() {
		rep := regexp.MustCompile(`\d+`)
		res := rep.FindAllString(scanner.Text(), -1)
		quantity, _ := strconv.Atoi(res[0])
		src, _ := strconv.Atoi(res[1])
		dst, _ := strconv.Atoi(res[2])
		for i := 0; i < quantity; i++ {
			stack[dst-1] = append(stack[dst-1], stack[src-1][len(stack[src-1])-1])
			stack[src-1] = stack[src-1][:len(stack[src-1])-1]
		}
	}
	for _, arr := range stack {
		fmt.Print(arr[len(arr)-1])
	}
	fmt.Println()
}

func part2(scanner *bufio.Scanner, stack [][]string) {
	for scanner.Scan() {
		rep := regexp.MustCompile(`\d+`)
		res := rep.FindAllString(scanner.Text(), -1)
		quantity, _ := strconv.Atoi(res[0])
		src, _ := strconv.Atoi(res[1])
		dst, _ := strconv.Atoi(res[2])
		var tmp []string
		for i := 0; i < quantity; i++ {
			tmp = append(tmp, stack[src-1][len(stack[src-1])-1])
			stack[src-1] = stack[src-1][:len(stack[src-1])-1]
		}
		// reverse tmp slice
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
		stack[dst-1] = append(stack[dst-1], tmp...)
	}
	for _, arr := range stack {
		fmt.Print(arr[len(arr)-1])
	}
	fmt.Println()
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	stack := genStack()
	scanner := bufio.NewScanner(fp)
	// part1(scanner, stack)
	part2(scanner, stack)
}
