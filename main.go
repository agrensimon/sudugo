package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type suduko struct {
	board            [][]int
	possibleRowMoves map[int][]int
}

func main() {}

func (s *suduko) solve(cursor int) bool {
	if cursor >= 9*9 {
		return true
	}

	// skip already occupied cells
	for s.board[cursor/9][cursor%9] != 0 {
		cursor++
	}

	toTry := len(s.possibleRowMoves[cursor/9])
	i := 0
	for i < toTry {
		n, ok := s.popMove(cursor / 9)
		if !ok {
			return false
		}

		s.board[cursor/9][cursor%9] = n
		if s.check(cursor/9, cursor%9) && s.solve(cursor+1) {
			return true
		}
		s.board[cursor/9][cursor%9] = 0
		s.appendMove(cursor/9, n)
		i++
	}
	return false
}

func (s *suduko) check(row, col int) bool {
	return s.checkRow(row) && s.checkCol(col) && s.checkSubGrid(row, col)
}

func (s *suduko) checkSubGrid(row, col int) bool {
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3

	occured := map[int]bool{}
	for r := rowStart; r < rowStart+3; r++ {
		for c := colStart; c < colStart+3; c++ {
			if s.board[r][c] == 0 {
				continue
			}
			if s.board[r][c] != 0 && occured[s.board[r][c]] {
				return false
			}
			occured[s.board[r][c]] = true
		}
	}

	return true
}

func (s *suduko) checkRow(row int) bool {
	occured := map[int]bool{}
	for c := 0; c < 9; c++ {
		if s.board[row][c] == 0 {
			continue
		}
		if occured[s.board[row][c]] {
			return false
		}
		occured[s.board[row][c]] = true
	}
	return true
}

func (s *suduko) checkCol(col int) bool {
	occured := map[int]bool{}
	for r := 0; r < 9; r++ {
		if s.board[r][col] == 0 {
			continue
		}
		if occured[s.board[r][col]] {
			return false
		}
		occured[s.board[r][col]] = true
	}
	return true
}

func newSuduko() suduko {
	s := suduko{}
	s.board = make([][]int, 9)
	s.possibleRowMoves = make(map[int][]int)
	for i := 0; i < 9; i++ {
		s.board[i] = make([]int, 9)
		s.possibleRowMoves[i] = append(s.possibleRowMoves[i], []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	}
	return s
}

func (s *suduko) dump() {
	for y, row := range s.board {
		if y%3 == 0 {
			fmt.Println("+---------+---------+---------+")
		}
		for x, s := range row {
			if x%3 == 0 {
				fmt.Print("|")
			}

			switch s {
			case 0:
				fmt.Print(" . ")
			default:
				fmt.Printf(" %v ", s)
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+---------+---------+---------+")
}

func parse(path string) (begin suduko, goal suduko, err error) {
	file, err := os.Open(path)
	if err != nil {
		return suduko{}, suduko{}, err
	}

	scanner := bufio.NewScanner(file)

	begin = newSuduko()
	goal = newSuduko()
	row, col := 0, 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		if len(split) == 0 {
			break
		}
		for _, r := range split {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				return suduko{}, suduko{}, err
			}
			begin.board[row][col] = i
			begin.possibleRowMoves[row] = remove(begin.possibleRowMoves[row], i)
			col++
		}
		col = 0
		row++
	}

	row, col = 0, 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		if len(split) == 0 {
			break
		}
		for _, r := range split {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				return suduko{}, suduko{}, err
			}
			goal.board[row][col] = i
			col++
		}
		col = 0
		row++
	}
	return
}

func (s *suduko) popMove(row int) (int, bool) {
	if len(s.possibleRowMoves[row]) == 0 {
		return -1, false
	}
	move := s.possibleRowMoves[row][0]
	s.possibleRowMoves[row][0] = 0
	s.possibleRowMoves[row] = s.possibleRowMoves[row][1:]
	return move, true
}

func (s *suduko) appendMove(row, move int) {
	s.possibleRowMoves[row] = append(s.possibleRowMoves[row], move)
}

// remove removes an integer from an integer slice if there is a matching value
func remove(arr []int, val int) []int {
	for i, elem := range arr {
		if elem == val {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
