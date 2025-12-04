package solutions

import (
	"strconv"
	"strings"
	"testing"
)

func aoc01b(t *testing.T) int {
	f, err := readFile("../inputs/01.txt")
	if err != nil {
		t.Error("could not read file", err)
	}

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
			for num > 0 {
				if current == 99 {
					current = 0
					// means we are ending at or passing 0
					counter++
				} else {
					current += 1
				}
				num--
			}
		} else {
			for num > 0 {
				if current == 0 {
					current = 99
				} else {
					current -= 1

					if current == 0 {
						// means we are ending at or passing 0
						counter++
					}
				}
				num--
			}
		}
	}

	return counter
}

func Test01b(t *testing.T) {
	t.Log("ANSWER:", aoc01b(t))
}
