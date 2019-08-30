package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var player_turn int
var playing bool
var passed_turn bool

func GetNeighbors(x int, y int) ([4]int, [4][2]int) {
	/*
	Returns neighbors in the form (Ltype,Rtype,Utype,Dtype),((Lx,Ly),(Rx,Ry),(Ux,Uy),(Dx,Dy))
	Type options:
	edge - 5  (This is 5 simply because 5 is not used yet elsewhere, and is less likely to get confused
	player - Players #
	opponent - Opps #
	free - 0
	*/
	var Type [4]int
	var Loc [4][2]int

	itercheck := [2]int{-1,1}
	iterchecky := [2]int{1,-1}

	for indx, elemx := range itercheck{
		if board_dim - 1 >= x + elemx >= 0 {
			Loc[indx][0] = x + elemx
			Loc[indx][1] = y
			Type[indx] = Board[x+elemx][y]
		}
	}
	for indy, elemy := range iterchecky{
		if board_dim - 1 >= y + elemy >= 0 {
			Loc[indy+2][0] = x
			Loc[indy+2][1] = y + elemy
			Type[indy+2] = Board[x][y+elemy]
		}
	}

	}
	// Handle edges/corners
	if x == 0 {
		Type[0] = 5
		Loc[0][0] = 99
		Loc[0][1] = 99
	}
	if x == board_dim - 1 {
		Type[1] = 5
		Loc[1][0] = 99
		Loc[1][1] = 99
	}
	if y == 0 {
		Type[2] = 5
		Loc[2][1] = 99
		Loc[2][1] = 99
	}
	if y == board_dim - 1 {
		Type[3] = 5
		Loc [3][1] = 99
		Loc [3][1] = 99
	}

	return Type, Loc
}

func Check_Captures() {
	//Will check if the given move has captured any spaces and alter board as necessary
}

func Check_Points() (bool, [2] int, [2] int) {
	//Will count points at end of match
	//Returns Draw? (bool), player (victor,loser) and points (victor, loser)
	return
}

func Player_Selection(player int) {
	// Listens for moves, checks captures, switches player

	PrintBoard()
	input_reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Player %d's move (ex: a5): \n",player)
	input := input_reader.ReadString('\n')
	fmt.Println()

	// End game if the opponent quits
	if strings.HasPrefix(strings.ToLower(input), "quit") {
		playing = false
		fmt.Printf("Player %d resigns. Opponent wins!", player)
		return
	}
	// Check if both players have skipped their turn, ending the game
	if strings.HasPrefix(strings.ToLower(input), "pass") {
		if passed_turn == true {
			playing = false
			draw, vic, points := Check_Points()
			if draw == true {
				fmt.Printf("Both Players have passed on their turn. Players 1 and 2 draw with %d points", points[0])
				return
			}
			if draw == false {
				fmt.Printf("Both Players have passed on their turn. Player %d wins with %d points. Player %d loses with %d points")
				return
			}
		}
	}

	xmove, ymove := AlphToBoard(input)
	if ValidMove(xmove,ymove) {
		board[xmove][ymove] = player
		passed_turn = false
		Check_Captures()
	}


	// Switch player state
	if player == 1 {
		player_turn = 2}
	if player == 2 {
		player_turn = 1}
	return
}


func main() {
	fmt.Println("Welcome to Golang Go!")
	fmt.Println("Specify where you wish to place to piece.")
	fmt.Println("If you wish to skip your turn, type 'pass'")
	fmt.Println("Once both players skip their turns or no more moves are available, the game ends")
	fmt.Println("Good luck!")

	playing = true
	PopulateNewBoard()
	for playing {
		Player_Selection()
	}
}