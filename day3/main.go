package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	// test-reddit from https://www.reddit.com/r/adventofcode/comments/189q9wv/2023_day_3_another_sample_grid_to_use/
	// Part 1: 413
	// Part 2: 6756
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var ans_p1, ans_p2, i int
	// (G)rid
	var G [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		G = append(G, []rune{})
		for _, r := range line {
			G[i] = append(G[i], r)
		}
		i++
	}
	// (C)olumns + (R)ows
	R := len(G)
	C := len(G[0])

	type gear struct {
		X int
		Y int
	}
	gears := make(map[gear]bool)
	nums := make(map[gear][]int)

	for r, line := range G {
		n := 0
		has_part := false
		for c := 0; c < len(line)+1; c++ {
			rune := '.'
			if c < C {
				rune = line[c]
			}
			if c < C && unicode.IsDigit(rune) {
				n = n*10 + int(rune-'0')
				for _, rr := range []int{-1, 0, 1} {
					for _, cc := range []int{-1, 0, 1} {
						if 0 <= r+rr && r+rr < R && 0 <= c+cc && c+cc < C {
							ch := G[r+rr][c+cc]
							if !unicode.IsDigit(ch) && ch != '.' {
								has_part = true
							}
							if ch == '*' {
								gears[gear{r + rr, c + cc}] = true
							}
						}
					}
				}
			} else if n > 0 {
				for gear, _ := range gears {
					nums[gear] = append(nums[gear], n)
				}
				if has_part {
					ans_p1 += n
				}
				n = 0
				has_part = false
				gears = make(map[gear]bool)
			}
		}
	}

	// fmt.Println(nums)
	for _, n := range nums {
		if len(n) == 2 {
			ans_p2 += n[0] * n[1]
		}
	}

	fmt.Println("Answer Part1:", ans_p1)
	fmt.Println("Answer Part2:", ans_p2)
}
