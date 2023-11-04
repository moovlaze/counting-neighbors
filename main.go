package main

import (
	"bufio"
	"fmt"
	"os"
)

type World struct {
	Height, Width int
	Cells         [][]bool
}

func (w *World) SaveState(filename string) error {

	if filename == "" {
		return fmt.Errorf("error")
	}

	file, err := os.Create(filename)
	defer file.Close()

	for i, arr := range w.Cells {
		str := ""
		for _, el := range arr {
			if el {
				str += "1"
			} else {
				str += "0"
			}
		}

		if i != len(w.Cells)-1 {
			str += "\n"
		}

		file.WriteString(str)
	}

	return err
}

func (w *World) LoadState(filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cells := [][]bool{}

	countHeight := 0
	for scanner.Scan() {

		cells = append(cells, []bool{})

		for _, num := range scanner.Text() {
			cells[countHeight] = append(cells[countHeight], string(num) == "1")
		}

		countHeight++
	}

	for i := len(cells) - 1; i >= 1; i-- {
		if len(cells[i]) != len(cells[i-1]) {
			return fmt.Errorf("error: cells is empty")
		}
	}

	w.Cells = cells

	return nil
}

func (w World) String() string {
	resultStr := ""

	for _, arr := range w.Cells {

		for _, el := range arr {
			if el {
				resultStr += "\xF0\x9F\x9F\xA9"
			} else {
				resultStr += "\xF0\x9F\x94\xA5"
			}
		}

		resultStr += "\n"
	}

	return resultStr
}

func (w *World) Neighbors(x, y int) int {
	n := 0

	for i := y - 1; i <= y+1; i++ {
		if i < 0 {
			i = len(w.Cells) - 1
		}

		for j := x - 1; j <= x+1; j++ {
			if j < 0 {
				j = len(w.Cells[i]) - 1
			}
			if w.Cells[i][j] {
				n++
			}
			if j == len(w.Cells[i])-1 {
				j = -1
			}
			if j == 0 && x == len(w.Cells[i])-1 {
				break
			}
		}

		if i == len(w.Cells)-1 {
			i = -1
		}

		if i == 0 && y == len(w.Cells)-1 {
			break
		}
	}

	return n
}

func main() {
	w := World{}

	w.LoadState("statefile.txt")
	// str := w.String()
	n := w.Neighbors(0, 0)

	// fmt.Println(str)
	fmt.Println(n)
}

// for i := y - 1; i <= y+1; i++ {
// 	if i < 0 {
// 		continue
// 	}
// 	if i > len(w.Cells)-1 {
// 		return n
// 	}
// 	for j := x - 1; j <= x+1; j++ {
// 		if j < 0 || j > len(w.Cells[i])-1 || i == y && j == x {
// 			continue
// 		}
// 		if w.Cells[i][j] {
// 			n++
// 		}
// 	}
// }
