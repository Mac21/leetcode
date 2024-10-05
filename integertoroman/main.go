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
	char  byte
	value int
	place int
}

var romanValues = []*Roman{
	{
		char:  'M',
		value: 1000,
		place: 2,
	},
	{
		char:  'D',
		value: 500,
		place: 1,
	},
	{
		char:  'C',
		value: 100,
		place: 1,
	},
	{
		char:  'L',
		value: 50,
		place: 0,
	},
	{
		char:  'X',
		value: 10,
		place: 0,
	},
	{
		char:  'V',
		value: 5,
		place: 0,
	},
	{
		char:  'I',
		value: 1,
		place: 0,
	},
}

func firstDigit(num int) (int, int) {
    if num < 10 {
        return num, 0
    }

	numDivs := 0
	for num > 0 {
		numDivs += 1
		temp := num / 10
		if temp == 0 {
			break
		}
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

		switch fd {
		case 9:
			fc := romanValues[len(romanValues)-divs-1]
			sc := romanValues[len(romanValues)-divs-2]
			res.WriteString(string(fc.char) + string(sc.char))
			num -= sc.value - fc.value
		case 4:
			fc := romanValues[len(romanValues)-divs-1]
			sc := romanValues[len(romanValues)-divs-2]
			res.WriteString(string(fc.char) + string(sc.char))
			num -= sc.value - fc.value
		default:
			count := num / r.value
			num = num % r.value
			// Need to use subtractive form
			if count <= 3 {
				for range count {
					res.WriteByte(r.char)
				}
			} else {
				if count == 4 {
					res.WriteString(string(r.char) + string(romanValues[i-1].char))
				}
			}
		}
	}
	return res.String()
}
