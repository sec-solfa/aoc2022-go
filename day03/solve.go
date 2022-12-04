package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	set "github.com/deckarep/golang-set/v2"
)

func gena2Z() string {
	alphabet := []byte{}
	var i byte
	for i = 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, i)
	}
	for i = 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, i)
	}
	return string(alphabet)
}

func genSet(input string) set.Set[string] {
	set := set.NewSet[string]()
	for _, c := range input {
		set.Add(string(c))
	}
	return set
}

func part1(scanner *bufio.Scanner) {
	alphabet := gena2Z()
	var total int
	for scanner.Scan() {
		length := len(scanner.Text())
		center := length / 2
		first, latter := genSet(scanner.Text()[:center]), genSet(scanner.Text()[center:])
		duplicate, _ := first.Intersect(latter).Pop()
		total += strings.Index(alphabet, duplicate) + 1
	}
	fmt.Println(total)
}

func part2(scanner *bufio.Scanner) {
	alphabet := gena2Z()
	var total int
	var buf []set.Set[string]
	for scanner.Scan() {
		set := genSet(scanner.Text())
		buf = append(buf, set)
		if len(buf) == 3 {
			duplicate, _ := buf[0].Intersect(buf[1]).Intersect(buf[2]).Pop()
			total += strings.Index(alphabet, duplicate) + 1
			buf = nil
		}
	}
	fmt.Println(total)
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
