package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func genSlice(scanner *bufio.Scanner) (rows, columns [][]string) {
	var tmpRows [][]string
	for scanner.Scan() {
		tmpRows = append(tmpRows, strings.Split(scanner.Text(), ""))
	}
	tmpColumns := make([][]string, len(tmpRows[0]))
	for i := 0; i < len(tmpRows); i++ {
		row := tmpRows[i]
		for i, s := range row {
			tmpColumns[i] = append(tmpColumns[i], s)
		}
	}
	return tmpRows, tmpColumns
}

func part1(rows, columns [][]string) {
	var count int
	for row_idx, row := range rows {
		for column_idx, target_s := range row {
			var views [][]string
			target_s_int, _ := strconv.Atoi(target_s)
			column := columns[column_idx]
			topView, bottomView := column[:row_idx], column[row_idx+1:]
			leftView, rightView := row[:column_idx], row[column_idx+1:]
			// fmt.Println(target_s, topView, bottomView, leftView, rightView)
			if len(topView) == 0 || len(bottomView) == 0 || len(leftView) == 0 || len(rightView) == 0 {
				count += 1
				continue
			}
			views = append(views, topView, bottomView, leftView, rightView)
			var found bool
			for _, view := range views {
				for idx, s := range view {
					if found {
						break
					}
					s_int, _ := strconv.Atoi(s)
					if s_int >= target_s_int {
						break
					}
					if idx+1 == len(view) {
						count += 1
						found = true
						break
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func part2(rows, columns [][]string) {
	var score int
	for row_idx, row := range rows {
		for column_idx, target_s := range row {
			target_s_int, _ := strconv.Atoi(target_s)
			column := columns[column_idx]
			topView, bottomView := column[:row_idx], column[row_idx+1:]
			leftView, rightView := row[:column_idx], row[column_idx+1:]
			// fmt.Println(target_s, topView, bottomView, leftView, rightView)
			var top, bottom, left, right int
			var topTmp, bottomTmp, leftTmp, rightTmp int
			if len(topView) > 0 {
				topTmp = target_s_int
				for i := len(topView) - 1; i >= 0; i-- {
					top += 1
					tmp, _ := strconv.Atoi(topView[i])
					if tmp >= topTmp {
						break
					}
				}
			}
			if len(leftView) > 0 {
				leftTmp = target_s_int
				for i := len(leftView) - 1; i >= 0; i-- {
					left += 1
					tmp, _ := strconv.Atoi(leftView[i])
					if tmp >= leftTmp {
						break
					}
				}
			}
			if len(bottomView) > 0 {
				bottomTmp = target_s_int
				for i := 0; i <= len(bottomView)-1; i++ {
					bottom += 1
					tmp, _ := strconv.Atoi(bottomView[i])
					if tmp >= bottomTmp {
						break
					}
				}
			}
			if len(rightView) > 0 {
				rightTmp = target_s_int
				for i := 0; i <= len(rightView)-1; i++ {
					right += 1
					tmp, _ := strconv.Atoi(rightView[i])
					if tmp >= rightTmp {
						break
					}
				}

			}
			// fmt.Println(target_s, "|", top, bottom, left, right)
			// fmt.Println(top * bottom * left * right)
			if score < top*bottom*left*right {
				score = top * bottom * left * right
			}
		}
	}
	fmt.Println(score)
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	rows, columns := genSlice(scanner)
	// part1(rows, columns)
	part2(rows, columns)
}
