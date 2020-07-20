package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type tictactoeboard [3][3]rune

const (
	playerShape   = 'O'
	computerShape = 'X'
)

func main() {

	var whoWon string
	var win bool
	var retry bool
	var board tictactoeboard
	retry = true
	for retry {
		board.clear()
		for i := 0; i < 9; i++ {
			board.displayBoard()
			if i%2 == 0 {
				board.player()
			} else {
				board.computer()
			}
			_, win = board.check()
			if win {
				break
			}
		}
		whoWon, win = board.check()
		if win {
			fmt.Printf("---------  %v  won  ---------\nFinal board:\n", whoWon)
		} else {
			fmt.Println("--------- draw ---------\nFinal board:")
		}
		board.displayBoard()

		retry = wantRetry()
	}

}

func wantRetry() bool {
	var valuePlay string
	fmt.Println("do you want retry this game ?(Y/N)")
	if _, err := fmt.Scan(&valuePlay); err == nil {
		if strings.ToUpper(valuePlay) == "Y" {
			return true
		} else if strings.ToUpper(valuePlay) == "N" {
			return false
		} else {
			fmt.Println("error input")
			wantRetry()
		}
	} else {
		fmt.Println("error input")
		wantRetry()
	}

	return false
}

func (t *tictactoeboard) displayBoard() {
	fmt.Print("-------------")
	for i := 0; i < 3; i++ {
		fmt.Print("\n|")
		for j := 0; j < 3; j++ {
			fmt.Printf(" %c |", t[i][j])
		}
		fmt.Print("\n------------")
	}
	fmt.Println()
}

func (t *tictactoeboard) clear() {
	for x, v := range t {
		for y := range v {
			t[x][y] = 0
		}
	}
}

func (t *tictactoeboard) player() {
	//set value of user play and check if is correct and not printed
	var column int32
	var line int32

	for {
		fmt.Println("Choose Row and column (between 1 and 3):")
		if _, err := fmt.Scan(&line, &column); err == nil {
			line--
			column--
			if line >= 0 && line <= 3 && column >= 0 && column <= 3 && t[line][column] == 0 {
				t[line][column] = playerShape
				break
			} else {
				if line >= 0 && line <= 3 && column >= 0 && column <= 3 {
					fmt.Printf("You input value is already use on line %v and column %v , value => %v\n", line+1, column+1, string(t[line][column]))
				} else {
					fmt.Printf("You input value is not corrected, line : %v , column : %v \n", line+1, column+1)
				}

			}
		} else {
			fmt.Println("Your input value is not correct")
			t.player()
		}
	}
}

func (t *tictactoeboard) computer() {
	// random value set for
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		line := r1.Intn(3)
		column := r1.Intn(3)
		if line >= 0 && line <= 3 && column >= 0 && column <= 3 && t[line][column] == 0 {
			t[line][column] = computerShape
			break
		}
	}

}

func (t *tictactoeboard) check() (string, bool) {
	//check if as a winner
	for x, v := range t {
		for y := range v {
			if t.checkArround(x, y) {
				if t[x][y] == playerShape {
					return "Player", true
				} else {
					return "Computer", true
				}
			}
		}
	}

	return "", false
}

func (t *tictactoeboard) checkArround(x, y int) bool {
	test := true
	symbole := t[x][y]
	if symbole == 0 {
		return false
	}
	for line := 0; line < 3; line++ {
		if t[line][y] != symbole {
			test = false
			break
		}
	}

	if !test {
		test = true
		for column := 0; column < 3; column++ {
			if t[x][column] != symbole {
				test = false
				break
			}
		}
		if !test {
			test = true
			if x == 1 && y == 1 {
				if t[0][0] != symbole || t[2][2] != symbole {
					test = false
				} else {
					test = true
				}
				if !test || (t[0][2] != symbole || t[2][0] != symbole) {
					test = false
				}
			} else if x == 0 && y == 0 {
				if t[1][1] != symbole || t[2][2] != symbole {
					test = false
				}
			} else if x == 0 && y == 2 {
				if t[1][1] != symbole || t[2][0] != symbole {
					test = false
				}
			} else if x == 2 && y == 0 {
				if t[1][1] != symbole || t[0][2] != symbole {
					test = false
				}
			} else if x == 2 && y == 2 {
				if t[1][1] != symbole || t[0][0] != symbole {
					test = false
				}
			} else {
				test = false
			}
		}
	}

	return test
}
