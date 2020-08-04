package main

import (
	"fmt"
	"os"
)

var cnt int
var ok bool

func main() {
	if len(os.Args[1:]) != 9 {
		fmt.Println("Error")
		return
	}

	for _, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}

		for _, ch := range arg {
			if ch <= '0' || ch > '9' {
				if ch != '.' {
					fmt.Println("Error")
					return
				}
			}
		}
	}

	arr := os.Args[1:]
	StrFromArr := ""
	for _, valueArr := range arr {
		for i := 0; i <= 8; i++ {
			if valueArr[i] == '.' {
				StrFromArr = StrFromArr + "0"
			} else {
				StrFromArr = StrFromArr + string(valueArr[i])
			}

		}
	}

	table := parseInput(StrFromArr)

	if solving(&table) {
		if cnt == 1 {
			ok = true
			solving(&table)
			if proverkaNaUnikalnostTab(&table) {
				printTable(table)
			} else {
				fmt.Println("Error")
			}
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}

}

func solving(table *[9][9]int) bool {
	if cnt > 1 && !ok {
		return false
	}
	if !poiskPustoiStroki(table) {
		cnt++

		if ok {
			return !poiskPustoiStroki(table)
		}
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				ok2 := false
				for candidate := 9; candidate >= 1; candidate-- {
					table[i][j] = candidate

					if proverkaNaUnikalnostTab(table) {
						if solving(table) {
							if !ok {
								ok2 = true
							} else {
								return true
							}
						}
						table[i][j] = 0
					} else {
						table[i][j] = 0
					}
				}
				if ok2 {
					return true
				}
				return false
			}
		}
	}
	return false
}

func poiskPustoiStroki(table *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func proverkaNaUnikalnostTab(table *[9][9]int) bool {

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[table[row][col]]++
		}
		if proverkaNaUnikalnostNomera(counter) {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[table[col][row]]++
		}
		if proverkaNaUnikalnostNomera(counter) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[table[row][col]]++
				}
				if proverkaNaUnikalnostNomera(counter) {
					return false
				}
			}
		}
	}

	return true
}

func proverkaNaUnikalnostNomera(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printTable(table [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col != 8 {
				fmt.Printf("%d ", table[row][col])
			} else {
				fmt.Printf("%d", table[row][col])
			}
		}
		fmt.Println()
	}
}

func parseInput(input string) [9][9]int {
	table := [9][9]int{}
	i := 0
	s := []rune(input)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			i1 := charVInt(s[i])
			i++
			table[row][col] = i1
		}
	}
	return table
}

func charVInt(ch rune) int {
	count := 0
	if ch < 48 && ch > 58 {
		return 0
	}

	for i := '1'; i <= ch; i++ {
		count++
	}

	return count
}
