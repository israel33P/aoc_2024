package main

import (
	s "strings"
	"fmt"
	"os"
	"strconv"
	"slices"
)

func eCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func textReader(filePath string) string {
	dat, err := os.ReadFile(filePath)
	eCheck(err)

	return string(dat)
}

func converter(inputText string) ([]int, []int) {
	lines := s.Split(inputText, "\n")
	var lCount int = 0
	var rCount int = 0
	var leftArr []int
	var rightArr []int
	leftArr = make([]int, len(lines))
	rightArr = make([]int, len(lines))
	for _, l := range lines {
		vals := s.Split(l, " ")
		lVal := s.TrimSpace(vals[0])
		rVal := s.TrimSpace(vals[len(vals)-1])
		if lVal != ""{
			i, err := strconv.Atoi(lVal)
			eCheck(err)
			leftArr[lCount] = i
			lCount++
		}
		if rVal != ""{
			i, err := strconv.Atoi(rVal)
			eCheck(err)
			rightArr[rCount] = i
			rCount++
		}
	}
	slices.Sort(leftArr)
	slices.Sort(rightArr)
	return leftArr, rightArr
}

func doTheMath(left []int, right []int) int {
	var output int = 0
	for i, l := range left {
		r := right[i]
		val := l-r
		if val < 0 {
			val *= -1
		}
		//fmt.Println(output)
		output += val 
	}
	return output
}

func main() {
	inputText := textReader("/mnt/c/Users/isrsanchez/Documents/_Scripting/AdventOfCode_2024/inputs/input_day1.txt")
	leftArr, rightArr := converter(inputText)
	output := doTheMath(leftArr, rightArr)

	fmt.Print(output)
	//fmt.Print(leftArr)
	//fmt.Println("\n=========================================")
	//fmt.Println(rightArr)
}
