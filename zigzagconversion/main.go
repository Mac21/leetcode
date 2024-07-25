package main

import (
    "fmt"
)

func cheatConvert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    result := ""
    increment := max(1, numRows*2-2)
    i := 0
    for ;i < numRows; i++{
        for x := 0; x+i < len(s); x += increment {
            result += string(s[x+i])
            if i != 0 && i != numRows-1 && x+increment - i < len(s) {
                result += string(s[x+increment-i])
            }
        }
    }

    return result
}

func convertDebug(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    numColumns := (len(s) / numRows) + numRows

    temp := make([][]string, numRows)
    for i := 0; i < numRows; i++ {
        temp[i] = make([]string, numColumns)
    }

    fmt.Printf("%#v\n", temp)

    temp[0][0] = string(s[0])


    currRow := 0
    currColumn := 0
    for i := 0; i < len(s); i++ {
        currRow = i % numRows
        currColumn = i / numRows


        temp[currRow][currColumn] = string(s[i])

        // for j := 0; j < currColumn; j++ {
            // columns.WriteString(" ")
        // }

        // for k := 0; k < currRow; k++ {
            // newLines.WriteString("\n")
        // }

        // fmt.Printf("%s%s%s", columns.String(), string(s[i]), newLines.String())
        // columns.Reset()
        // newLines.Reset()
    }
    fmt.Printf("%#v\n", temp[0])
    fmt.Printf("%#v\n", temp[1])
    fmt.Printf("%#v\n", temp[2])

    return ""
}

func main() {
    input := "PAYPALISHIRING"
    answer := "PAHNAPLSIIGYIR"
    rows := 3
    result := cheatConvert(input, rows)
    fmt.Printf("Test success: %v? Expected %s got %s\n\n", answer == result, answer, result)

    answer = "PINALSIGYAHRPI"
    rows = 4
    result = cheatConvert(input, rows)
    fmt.Printf("Test success: %v? Expected %s got %s\n\n", answer == result, answer, result)

    input = "A"
    answer = "A"
    rows = 1
    result = cheatConvert(input, rows)
    fmt.Printf("Test success: %v? Expected %s got %s\n", answer == result, answer, result)

    input = "PAYPALISHIRING"
    answer = "PAYPALISHIRING"
    rows = 1
    result = cheatConvert(input, rows)
    fmt.Printf("Test success: %v? Expected %s got %s\n", answer == result, answer, result)

    input = "PAYPALISHIRING"
    answer = "PAYPALISHIRING"
    rows = 15
    result = cheatConvert(input, rows)
    fmt.Printf("Test success: %v? Expected %s got %s\n", answer == result, answer, result)
}
