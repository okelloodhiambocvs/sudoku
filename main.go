package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	b := make([][]rune, 9)
	for i, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		row := []rune(arg)
		for _, v := range row {
			if v != '.' && (v < '1' || v > '9') {
				fmt.Println("Error")
				return
			}
		}
		b[i] = row
	}

	for i := range b {
		for j, v := range b[i] {
			if v != '.' {
				b[i][j] = '.'
				if !safe(b, i, j, v) {
					fmt.Println("Error")
					return
				}
				b[i][j] = v
			}
		}
	}

	if !solve(b) {
		fmt.Println("Error")
		return
	}

	for _, r := range b {
		for j, v := range r {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(v))
		}
		fmt.Println()
	}
}

func solve(b [][]rune) bool {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == '.' {
				for d := '1'; d <= '9'; d++ {
					if safe(b, i, j, d) {
						b[i][j] = d
						if solve(b) {
							return true
						}
						b[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func safe(b [][]rune, r, c int, v rune) bool {
	sr, sc := r/3*3, c/3*3
	for i := 0; i < 9; i++ {
		if b[r][i] == v || b[i][c] == v || b[sr+i/3][sc+i%3] == v {
			return false
		}
	}
	return true
}