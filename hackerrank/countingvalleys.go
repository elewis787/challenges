package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func calcValleys(s string) int {
	num := 0
	valleys := 0
	inValley := false
	for _, v := range s {
		if strings.ToLower(string(v)) == "u" {
			num++
		} else if strings.ToLower(string(v)) == "d" {
			num--
		}
		if num < 0 {
			//in vallye
			inValley = true
		}
		if inValley && num == 0 {
			// out of valley
			valleys++
			inValley = false
		}
	}
	return valleys
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		panic(err)
	}
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	strNumber := readLine(reader)
	_, err = strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		panic(err)
	}
	strValues := readLine(reader)
	/*	strValuesArry := strings.Split(strValues, "")
		values := make([]int, 0)
		for i := 0; i < int(number); i++ {
			strVal, err := strconv.ParseInt(strValuesArry[i], 10, 64)
			if err != nil {
				panic(err)
			}
			intVal := int(strVal)
			values = append(values, intVal)
		}*/
	result := calcValleys(string(strValues))
	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
}

// simple version
/*
package main
import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)

    var s string
    fmt.Scanln(&s)

    l, v := 0, 0
    for i, _ := range s {
        if s[i] == 'U' {
            if l < 0 && l+1 == 0 {
                v++
            }

            l++
        } else if s[i] == 'D' {
            l--
        }
    }

    fmt.Println(v)
}
*/
