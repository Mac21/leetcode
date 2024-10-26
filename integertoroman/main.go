package main

import (
	"fmt"
	"strings"
)

/*
12. Integer to Roman Medium
Seven different symbols represent Roman numerals with the following values:

Symbol	Value
I	1
V	5
X	10
L	50
C	100
D	500
M	1000
Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
Given an integer, convert it to a Roman numeral.



Example 1:

Input: num = 3749

Output: "MMMDCCXLIX"

Explanation:

3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
 700 = DCC as 500 (D) + 100 (C) + 100 (C)
  40 = XL as 10 (X) less of 50 (L)
   9 = IX as 1 (I) less of 10 (X)
Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
Example 2:

Input: num = 58

Output: "LVIII"

Explanation:

50 = L
 8 = VIII
Example 3:

Input: num = 1994

Output: "MCMXCIV"

Explanation:

1000 = M
 900 = CM
  90 = XC
   4 = IV


Constraints:

1 <= num <= 3999
*/

type Roman struct {
	char  string
	value int
}

var romanValues = []*Roman{
	{
		char:  "M",
		value: 1000,
	},
	{
		char:  "D",
		value: 500,
	},
	{
		char:  "C",
		value: 100,
	},
	{
		char:  "L",
		value: 50,
	},
	{
		char:  "X",
		value: 10,
	},
	{
		char:  "V",
		value: 5,
	},
	{
		char:  "I",
		value: 1,
	},
}

// Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
func getSubtractiveForm(cd *Roman) (string, int) {
	switch cd.char {
	case "M":
		return "CM", 900
	case "D":
		return "CD", 400
	case "C":
		return "XC", 90
	case "L":
		return "XL", 40
	case "X":
		return "IX", 9
	case "V":
		return "IV", 4
	default:
		return "", 0
	}
}

// firstDigit returns the leftmost number and how many times the number was divided by 10
// e.g. num = 49
// firstDigit(49) -> (4, 1)
func firstDigit(num int) (int, int) {
	if num < 10 {
		return num, 0
	}

	numDivs := 0
	for num > 0 {
		temp := num / 10
		if temp == 0 {
			break
		}
		numDivs += 1
		num = temp
	}

	return num, numDivs
}

func intToRoman(num int) string {
	if num < 1 {
		return ""
	}

	var res strings.Builder
	for i, r := range romanValues {
		if num < 1 {
			break
		}
		fd, divs := firstDigit(num)
		fmt.Println(num, r, fd, divs)

		ch, val := getSubtractiveForm(r)
		if val <= num && (fd == 9 || fd == 4) {
			res.WriteString(ch)
			num -= val
		} else {
			count := num / r.value
			num = num % r.value
			// Need to use subtractive form
			if count <= 3 {
				for range count {
					res.WriteString(r.char)
				}
			} else {
				if count == 4 {
					res.WriteString(r.char + romanValues[i-1].char)
				}
			}
		}
	}
	return res.String()
}
