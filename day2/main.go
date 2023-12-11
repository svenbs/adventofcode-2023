package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	possible := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	ans_p1 := 0
	ans_p2 := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		subsets := strings.Split(line[1], ";")

		gameID, _ := strconv.Atoi(strings.Split(line[0], " ")[1])

		minNeeded := make(map[string]int)
		valid := true
		for _, subset := range subsets {
			for _, set := range strings.Split(subset, ",") {
				s := strings.Split(set, " ")
				cubeCount, _ := strconv.Atoi(s[1])
				colour := s[2]

				minNeeded[colour] = max(cubeCount, minNeeded[colour])

				if cubeCount > possible[colour] {
					valid = false
				}
			}
		}
		if valid {
			ans_p1 += gameID
		}
		ans_p2 += minNeeded["red"] * minNeeded["green"] * minNeeded["blue"]
	}
	fmt.Println("Answer Part 1:", ans_p1)
	fmt.Println("Answer Part 2:", ans_p2)
}