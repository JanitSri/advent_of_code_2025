package solutions

import (
	"os"
	"strconv"
)

type coord [2]int

var directionMap = map[string][2]int{
	"up":        {-1, 0},
	"upRight":   {-1, 1},
	"right":     {0, 1},
	"downRight": {1, 1},
	"down":      {1, 0},
	"downLeft":  {1, -1},
	"left":      {0, -1},
	"upLeft":    {-1, -1},
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func resetGrid[T comparable](g [][]T, o, t T) {
	for r := range g {
		for c := range g[r] {
			if g[r][c] == o {
				g[r][c] = t
			}
		}
	}
}

func contains[T comparable](a []T, x T) bool {
	for _, i := range a {
		if i == x {
			return true
		}
	}

	return false
}

func countDigits(i int) int {
	var r int

	for i > 0 {
		i /= 10
		r++
	}

	return r
}

func splitDigitsInHalf(i int) []int {
	s := strconv.Itoa(i)
	mid := len(s) / 2
	l, _ := strconv.Atoi(s[:mid])
	r, _ := strconv.Atoi(s[mid:])

	return []int{l, r}
}
