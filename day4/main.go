package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	p1 := 0
	p2 := 0
	N := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		card := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Fields(card[0])[1])
		N[id] += 1
		numbers := strings.Split(card[1], "|")
		card_nums := strings.Fields(numbers[0])
		rest_nums := strings.Fields(numbers[1])

		count := 0
		for _, cn := range card_nums {
			for _, cr := range rest_nums {
				d, err := strconv.Atoi(cn)
				if err != nil {
					continue
				}
				wd, err := strconv.Atoi(cr)
				if err != nil {
					continue
				}

				if d == wd {
					count++
				}
			}
		}
		if count > 0 {
			p1 += int(math.Pow(2, float64(count-1)))
		}
		for j := 1; j <= count; j++ {
			N[id+j] += N[id]
		}
	}
	for _, v := range N {
		p2 += v
	}
	fmt.Println("Answer Part1: ", p1)
	fmt.Println("Answer Part2: ", p2)
}
