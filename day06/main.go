package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(scanner *bufio.Scanner) {
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		for i := 3; i <= len(arr); i++ {
			a, b, c, d := arr[i-3], arr[i-2], arr[i-1], arr[i]
			if a != b && a != c && a != d && b != c && b != d && c != d {
				fmt.Println(i + 1)
				break
			}
		}
	}
}

func part1_beta(scanner *bufio.Scanner) {
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		detected := false
		for i := 4; i < len(arr); i++ {
			encountered := map[string]bool{}
			tmp := arr[i-4 : i]
			for j := 0; j < len(tmp); j++ {
				if !encountered[tmp[j]] {
					encountered[tmp[j]] = true
				} else {
					break
				}
				if j == 3 {
					detected = true
					break
				}
			}
			if detected {
				fmt.Println(i)
				break
			}
		}
	}
}

func part2(scanner *bufio.Scanner) {
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		detected := false
		for i := 14; i < len(arr); i++ {
			encountered := map[string]bool{}
			tmp := arr[i-14 : i]
			for j := 0; j < len(tmp); j++ {
				if !encountered[tmp[j]] {
					encountered[tmp[j]] = true
				} else {
					break
				}
				if j == 13 {
					detected = true
					break
				}
			}
			if detected {
				fmt.Println(i)
				break
			}
		}
	}
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	// part1(scanner)
	// part1_beta(scanner)
	part2(scanner)
}
