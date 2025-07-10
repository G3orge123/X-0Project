package main

import (
	"fmt"
)

// Structure for the X&0 Game
type Game struct {
	board   [3][3]byte
	current byte
	moves   int
}

// Initialise a new X&0 Game
func NewGame() *Game {
	g := &Game{current: 'X'}
	for i := range g.board {
		for j := range g.board[i] {
			g.board[i][j] = '*'
		}
	}
	return g
}

// Displays the board on the screen
func (g *Game) printBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
		for j := 0; j < 3; j++ {
			fmt.Printf("%c", g.board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println(" ------")
		}
	}
}

// Check if the moves are valid
func (g *Game) makeMoves(row, col int) bool {
	if g.board[row][col] != '*' {
		return false
	}
	if row < 0 || row > 2 || col > 2 || col < 0 {
		return false
	}
	g.board[row][col] = g.current
	g.moves++
	return true
}

// Check if the player Won the Game
func (g *Game) checkWin() bool {
	c := g.current
	for i := 0; i < 3; i++ {
		if g.board[i][0] == c && g.board[i][1] == c && g.board[i][2] == c {
			return true
		}
		if g.board[0][i] == c && g.board[1][i] == c && g.board[2][i] == c {
			return true
		}
	}
	if g.board[0][0] == c && g.board[1][1] == c && g.board[2][2] == c {
		return true
	}
	if g.board[0][2] == c && g.board[1][1] == c && g.board[2][0] == c {
		return true
	}
	return false
}

// Check if the Game is a Draw
func (g *Game) isDraw() bool {
	if g.moves == 9 {
		return true
	} else {
		return false
	}
}

// Switches the players
func (g *Game) switchPlayer() {
	if g.current == 'X' {
		g.current = '0'
	} else {
		g.current = 'X'
	}
}

func main() {
	g := NewGame()
	var row, col int
	for {
		g.printBoard()
		fmt.Println()
		fmt.Printf("Player %c, moves (row,col):  ", g.current)
		n, err := fmt.Scan(&row, &col)
		fmt.Println()
		if err != nil || n != 2 || row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println()
			fmt.Println("Invalid Input, Introduce a number between 0 and 2")
			fmt.Println()
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if !g.makeMoves(row, col) {
			fmt.Println()
			fmt.Println("The move is invalid,please try again")
			fmt.Println()
			continue
		}
		if g.checkWin() {
			g.printBoard()
			fmt.Println()
			fmt.Printf("Player %c has won the game\n", g.current)
			fmt.Println()
			break
		}
		if g.isDraw() {
			g.printBoard()
			fmt.Println()
			fmt.Println("Draw")
			fmt.Println()
			break
		}
		g.switchPlayer()
	}
}
