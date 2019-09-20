package art

import (
	"fmt"
	"sort"
	"strings"
)

// String returns an art string
func String(text string) string {
	if len(text) == 0 {
		return ""
	}
	chars := []byte(text)
	s := getArt(chars[0])
	for i := 1; i < len(chars); i++ {
		s = fuse(s, getArt(chars[i]))
	}
	return strings.Replace(strings.Join(s, "\n"), "$", " ", -1)
}

// Print prints the art string
func Print(text string) {
	fmt.Println(String(text))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fuse(left, right []string) []string {
	var step []int
	c := min(len(left[0]), len(right[0]))
	for j := 0; j < 6; j++ {
		k := len(left[j]) + len(right[j]) -
			len(strings.TrimRight(left[j], " ")) -
			len(strings.TrimLeft(right[j], " "))
		step = append(step, k)
	}

	sort.Ints(step)
	c = step[0]

	for j := 0; j < 6; j++ {
		cc := ""
		for i := 1; i <= c; i++ {
			lc := left[j][len(left[j])-(c-i+1) : len(left[j])-(c-i)]
			rc := right[j][i-1 : i]
			if lc == " " {
				cc += rc
			}else{
				cc += lc
			}
		}
		left[j] = left[j][:len(left[j])-c] + cc + right[j][c:]
	}
	return left
}