
package main

import (
	s "strings"
	"fmt"
	"os"
	"math"
	"strconv"
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

func converter(inputText string) (int) {
	lines := s.Split(inputText, "\n")
	var safe int = 0
	for _, l := range lines {
		line := s.TrimSpace(l)
		if line != ""{
			numbs := s.Split(l, " ")
			var temp float64 = 0
			var isSafe bool = true
			for i:=1; i < len(numbs); i++{
				s0 := numbs[i-1]
				s1 := numbs[i]
				n0, err := strconv.ParseFloat(s0, 64)
				eCheck(err)
				n1, err := strconv.ParseFloat(s1, 64)
				eCheck(err)
				nDiff := n0 - n1
				if i == 1 {
					if math.Abs(nDiff) > 3 {
						isSafe = false
						break
					}
					temp = nDiff
					continue
				}
				if (temp*nDiff) <= 0 {
					isSafe = false
					break
				}
				if math.Abs(nDiff) > 3 {
					isSafe = false
					break
				}
				temp = nDiff
			}
			//fmt.Printf("-- %s: %t\n", line, isSafe)
			if isSafe {
				safe++
			}
		}
	}
	return safe
}

func main() {
	inputText := textReader("/mnt/c/Users/isrsanchez/Documents/_Scripting/AdventOfCode_2024/inputs/input_day2.txt")
	safe := converter(inputText)

	fmt.Println(safe)
	//fmt.Print(leftArr)
	//fmt.Println("\n=========================================")
	//fmt.Println(rightArr)
}
