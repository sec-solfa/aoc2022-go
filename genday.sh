#!/bin/bash


day=${1##+(0)}
if ((day < 1 || day > 25)); then
    return
fi
    
project=$(printf "day%02d" $1)
session="53616c7465645f5f0b0901e18eae495316dca7b37ca0340fdd3cb205e8afcfad786743fb9f5a0297128c3850635b39e849b447b1f362d4d0c36d7178d091cc67"
mkdir $project
cd ${project}

curl -s "https://adventofcode.com/2022/day/${day}/input" --cookie "session=${session}" -o input.txt



echo -n "package main
import (
	\"bufio\"
	\"fmt\"
	\"os\"
)

func readFile(fileName string) {
    fp, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}


func main() {
    readFile(\"input.txt\")
}" > solve.go
