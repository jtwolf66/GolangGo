package main

import (
	"fmt"
	"bufio"
)

var player_turn int
var playing bool

func Check_Captures() {
	//Will check if the given move has captured any spaces and alter board as necessary
}

func Check_Points() {
	//Will count points at end of match
}

func Player_Selection(player int, loc string ) {
	PrintBoard()
	// Listens for moves, checks captures, switches player
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