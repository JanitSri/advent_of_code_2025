package solutions

import (
	"strconv"
	"strings"
	"testing"
)

func aoc01a(t *testing.T) int {
	f := ReadFile("../inputs/01.txt")
	lines := strings.Split(f, "\n")

	current := 50
	counter := 0
	for _, line := range lines {
		dir := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			t.Error("could not convert to int", num)
		}

		if dir == 'R' {
			current = (current + num) % 100
		} else {
			current -= num % 100

			if current < 0 {
				current = 100 - (current * -1)
			}
		}

		if current == 0 {
			counter++
		}
	}

	return counter
}

func Test01a(t *testing.T) {
	t.Log("ANSWER:", aoc01a(t))
}
