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
	if strings.HasPrefix(strings.ToLower(input), "pass") {
		if passed_turn == true {
			draw, vic, points := Check_Points()
			if draw == true {
				fmt.Printf("Both Players have passed on their turn. Players 1 and 2 draw with %d points",points[0])
				return
			}
			if draw == false {
				fmt.Printf("Both Players have passed on their turn. Player %d wins with %d points. Player %d loses with %d points")
				return
			}
		}


	// Switch player state
	if player == 1 {
		player_turn = 2}
	if player == 2 {
		player_turn = 1}
	}
	return

}

func main() {
	fmt.Println("Welcome to Golang Go!")
	fmt.Println("Specify where you wish to place to piece.")
	fmt.Println("If you wish to skip your turn, type 'pass'")
	fmt.Println("Once both players skip their turns or no more moves are available, the game ends")
	fmt.Println("Good luck!")
	
	PopulateNewBoard()
	for playing {
		Player_Selection()
	}
}