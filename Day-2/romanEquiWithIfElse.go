package main

import (
	"fmt"
	"strings"
)

func romanEqui(inp string) string {

	var curVal int
	res := 0
	prevVal := -1
	flag := false

	for i := len(inp) - 1; i >= 0; i-- {
		if inp[i] == 'I' {
			curVal = 1
		} else if inp[i] == 'V' {
			curVal = 5
		} else if inp[i] == 'X' {
			curVal = 10
		} else if inp[i] == 'L' {
			curVal = 50
		} else if inp[i] == 'C' {
			curVal = 100
		} else if inp[i] == 'D' {
			curVal = 500
		} else if inp[i] == 'M' {
			curVal = 1000
		} else {
			flag = true
			break
		}

		if prevVal > curVal {
			res -= curVal
		} else {
			res += curVal
		}

		prevVal = curVal
	}

	if flag {
		return "Invalid Roman Number !!"
	} else {
		output := fmt.Sprintf("Decimal Equivalent of Given Roman Number is : %d", res)
		return output
	}
}

func main() {

	var inp string
	fmt.Println("Enter Roman Value to Find it's Decimal Value (IN ROMAN,ALL LETTERS ARE IN CAPITAL): ")
	fmt.Scan(&inp)

	inp = strings.ToUpper(inp) //Making it Capital If in case User Enters Inp in Lowercase

	result := romanEqui(inp)
	fmt.Println(result)

}
