package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/emilioziniades/adventofcode2021/fetch"
)

func main() {
	fmt.Println("vim-go")
}

type Board struct {
	board  [][]int
	marked map[int]bool
}

func (b *Board) initialize() {
	b.board = make([][]int, 5)
	for i := 0; i < 5; i++ {
		b.board[i] = make([]int, 0)
	}

	b.marked = make(map[int]bool)
}

func (b *Board) isWinner() bool {
	//horizontal winner
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if n := b.board[i][j]; !b.marked[n] {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	//vertical winner
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if n := b.board[j][i]; !b.marked[n] {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	return false
}

func (b *Board) countScore(num int) int {
	var res int
	for i, _ := range b.board {
		for _, d := range b.board[i] {
			if !b.marked[d] {
				res += d
			}
		}
	}
	return res * num
}

//playBingo goes through each number until someone is a winner. It then calulates the winner's score
func PlayBingo(nums []int, boards []Board) int {
	var score int
Rounds:
	for _, num := range nums {
		for i := 0; i < len(boards); i++ {
			boards[i].marked[num] = true
			if b := boards[i]; b.isWinner() {
				score = b.countScore(num)
				break Rounds
			}
		}
	}
	return score
}

func LoseBingo(nums []int, boards []Board) int {
	winners := make(map[int]bool)
	score := 0
Rounds:
	for _, num := range nums {
		for i := 0; i < len(boards); i++ {
			boards[i].marked[num] = true
			if winners[i] {
				continue
			}
			if b := boards[i]; b.isWinner() {
				winners[i] = true
				if len(boards) == len(winners) {
					score = b.countScore(num)
					break Rounds
				}
			}
		}
	}
	return score
}

func ParseBingo(file string) ([]Board, []int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	input := bufio.NewScanner(f)

	// get numbers
	input.Scan()
	nums, err := fetch.StringToIntSlice(strings.Split(input.Text(), ","))
	if err != nil {
		return nil, nil, err
	}
	input.Scan() //scans first newline before board 1

	// get boards
	nBoards := numBoards(file)
	boards := make([]Board, nBoards)
	for i := 0; i < len(boards); i++ {
		boards[i].initialize()
	}

	n := 0
	nRow := 0
	for input.Scan() {
		if input.Text() == "" {
			n++
			nRow = 0
			continue
		}

		row, err := fetch.StringToIntSlice(strings.Fields(input.Text()))
		if err != nil {
			return nil, nil, err
		}
		boards[n].board[nRow] = append(boards[n].board[nRow], row...)
		nRow++
	}
	return boards, nums, nil

}

func numBoards(file string) int {
	lines, err := fetch.ParseInputString(file)
	if err != nil {
		log.Fatalf("Bingo: numBoards: %s", err)
	}
	var n int
	for _, line := range lines {
		if line == "" {
			n++
		}
	}
	return n
}
