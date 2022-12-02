package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toInt_1(a string) int {
	if a == "A" || a == "X" {
		return 0
	} else if a == "B" || a == "Y" {
		return 1
	} else if a == "C" || a == "Z" {
		return 2
	} else {
		panic("Error")
	}
}

func jadge_1(a, b string) int {
	a_hand, b_hand := toInt_1(a), toInt_1(b)
	point := b_hand + 1
	result := (a_hand - b_hand + 3) % 3
	if result == 2 { // win
		return 6 + point
	} else if result == 0 { // draw
		return 3 + point
	} else if result == 1 { // lose
		return 0 + point
	} else {
		panic("Error")
	}
}

func toInt_2(a string) int {
	if a == "A" {
		return 0
	} else if a == "B" {
		return 1
	} else if a == "C" {
		return 2
	} else {
		panic("Error")
	}
}

func determineHands(a_hand int, result string) (int, int) {
	if result == "X" { // lose
		return (a_hand + 2) % 3, 0
	} else if result == "Y" { // draw
		return a_hand, 3
	} else if result == "Z" { // win
		return (a_hand + 1) % 3, 6
	} else {
		panic("Error")
	}
}

func calc_point(a_hand int, result string) int {
	b_hand, score := determineHands(a_hand, result)
	fmt.Println(b_hand, score, (b_hand+1)+score)
	return (b_hand + 1) + score
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	// var total_1 int
	var total_2 int
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		// total_1 += jadge_1(arr[0], arr[1])
		total_2 += calc_point(toInt_2(arr[0]), arr[1])
	}
	fmt.Println(total_2)
}
