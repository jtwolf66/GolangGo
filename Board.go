package main

import ("fmt"
		"strconv"
		"strings"
)

var board_dim int = 13
var board [ board_dim ][ board_dim ] int

func ClearBoard() {
	board = [ board_dim ][ board_dim ] int {
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0} ,
		{0,0,0,0,0,0,0,0,0,0,0,0,0}
	}
	
}

func Boardchar() string {
	piece_dict := map[int]string{
		0: "=" // Empty Spaces
		1: "X" // Player 1 stone
		2: "O" // Player 2 stone
		3: "x"  // Player 1 capture
		4: "o" // Player 2 capture
	}
	board_str := ""
	for i := 0; i < board_dim; i++ {
		board_str += strconv.Itoa(board_dim-i)
		for j := 0; j < board_dim; j++ {
			board_str += " | " + piece_dict[board[i][j]]
		}
		board_str += " |\n"

	}
	board_str += "    A   B   C   D   E   F   G   H   I   J   K   L   M "
	return board_str

}

func ValidMove(x, y int) bool {
	if x > board_dim-1 || x < 0 || y > board_dim-1 || hoz < 0 {
		return false
	}
	if board[x][y] != 0 {
		return false
	}
	return true
}

func PlayerMove(player, x, y int ) {
	if ValidMove(x, y) == false {
		fmt.Println("Choose a Valid Move")
		return
	}
	if player == 1 {
		board[x][y] = 1
	}
	if player == 2 {
		board[x][y] = 2
	}
	return
}

func AlphToBoard(alpha_input string) ([2]int){
	vertical := map[string]int{
		"13": 0,
		"12": 1,
		"11": 2,
		"10": 3,
		"9": 4,
		"8": 5,
		"7": 6,
		"6": 7,
		"5": 8,
		"4": 9,
		"3": 10,
		"2": 11,
		"1": 12,
	}

	horizontal := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12
	}
	return horizontal[alpha_input[0]], vertical[alpha_input[1]]
}