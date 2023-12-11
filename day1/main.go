package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	ans_p1 := 0
	ans_p2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits_p1 := make([]int, 0)
		digits_p2 := make([]int, 0)

		for i, r := range line {
			if unicode.IsDigit(r) {
				digits_p1 = append(digits_p1, int(r - '0'))
				digits_p2 = append(digits_p2, int(r - '0'))
			}
			for d, val := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], val) {
					digits_p2 = append(digits_p2, d)
				}
			}
		}

		score_p1, _ := strconv.Atoi(fmt.Sprintf("%d%d", digits_p1[0], digits_p1[len(digits_p1)-1]))
		score_p2, _ := strconv.Atoi(fmt.Sprintf("%d%d", digits_p2[0], digits_p2[len(digits_p2)-1]))
		ans_p1 += score_p1
		ans_p2 += score_p2
	}

	fmt.Println(ans_p1)
	fmt.Println(ans_p2)
}