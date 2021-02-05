package main

/*
There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

Example

There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .

Function Description

Complete the sockMerchant function in the editor below.

sockMerchant has the following parameter(s):

int n: the number of socks in the pile
int ar[n]: the colors of each sock
Returns

int: the number of pairs
Input Format

The first line contains an integer , the number of socks represented in .
The second line contains  space-separated integers, , the colors of the socks in the pile.

Constraints

 where
Sample Input

STDIN                       Function
-----                       --------
9                           n = 9
10 20 20 10 10 30 50 10 20  ar = [10, 20, 20, 10, 10, 30, 50, 10, 20]
Sample Output

3
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func pairs(ar []int) int {
	matches := make(map[int]int)
	for i := 0; i < len(ar); i++ {
		matches[ar[i]]++
	}
	// final results
	results := 0
	// loop over map key/values
	// divide count by 2 ( because pairs are even numbers )
	for _, cnt := range matches {
		results += cnt / 2 // take whole number
	}
	return results
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	strNumber := readLine(reader)
	number, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		panic(err)
	}
	strValues := readLine(reader)
	strValuesArry := strings.Split(strValues, " ")
	values := make([]int, 0)
	for i := 0; i < int(number); i++ {
		strVal, err := strconv.ParseInt(strValuesArry[i], 10, 64)
		if err != nil {
			panic(err)
		}
		intVal := int(strVal)
		values = append(values, intVal)
	}
	result := pairs(values)
	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
}
