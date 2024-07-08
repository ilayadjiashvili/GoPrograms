package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BOARD
type Board struct {
	board [][]string
}

func NewBoard(size int) *Board {
	b := new(Board)
	for i := 0; i < size; i++ {
		var curLine []string
		for j := 0; j < size; j++ {
			curLine = append(curLine, "_")
		}
		b.board = append(b.board, curLine)
	}
	return b
}

func (b *Board) printBoard() {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			fmt.Printf(" %s ", b.board[i][j])
		}
		fmt.Println()
	}
}

func (b *Board) isFull() bool {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if b.board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

// PLAYER

type Player struct {
	sign string
}

func newPlayer(sign string) *Player {
	p := new(Player)
	p.sign = sign
	return p
}

func (p *Player) playMove(b *Board, row int, col int) bool {
	if row > len(b.board) || col > len(b.board[0]) || row < 0 || col < 0 || b.board[row][col] != "_" {
		fmt.Println("Not A Valid Coordinate. Please Enter A New Coordinate")
		return false
	} else {
		b.board[row][col] = p.sign
		return true
	}
}

func runTicTacToe(size int) {
	b := NewBoard(size)
	scanner := bufio.NewScanner(os.Stdin)
	xPlayer := newPlayer("x")
	oPlayer := newPlayer("o")
	var players [2]Player
	players[0] = *xPlayer
	players[1] = *oPlayer
	curPlayerIndex := 0
	var row, col int
	for !b.isFull() {
		b.printBoard()
		fmt.Println("Hi there ", players[curPlayerIndex].sign, " player! whats your move? ")
		scanner.Scan()
		input := scanner.Text()
		coordinates := strings.Split(input, " ")
		row, _ = strconv.Atoi(coordinates[0])
		col, _ = strconv.Atoi(coordinates[1])
		for !players[curPlayerIndex].playMove(b, row, col) {
			scanner.Scan()
			input := scanner.Text()
			coordinates := strings.Split(input, " ")
			row, _ = strconv.Atoi(coordinates[0])
			col, _ = strconv.Atoi(coordinates[1])
		}
		curPlayerIndex = (curPlayerIndex + 1) % 2
	}
}

func main() {
	runTicTacToe(3)
}
