package main

import ("fmt"
		"strconv"
		"strings"
)

//TODO: Add more board sizes

var board_dim int = 13
var passed_turn bool = false


type Board [13][13]int

func GenerateBoard() Board {
	board := Board{
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
		{0,0,0,0,0,0,0,0,0,0,0,0,0}}
	return board
}


func (b Board) Str_Form() string {
	piece_dict := map[int]string{
		0: "=", // Empty Spaces
		1: "X", // Player 1 stone
		2: "O"} // Player 2 stone
	board_str := ""
	for i := 0; i < board_dim; i++ {
		board_str += strconv.Itoa(board_dim-i)
		for j := 0; j < board_dim; j++ {
			board_str += " | " + piece_dict[b[i][j]]
		}
		board_str += " |\n"

	}
	board_str += "    A   B   C   D   E   F   G   H   I   J   K   L   M "
	return board_str
}

func (b Board) Print_Str_Form() {
	fmt.Println(b.Str_Form())
	return
}
/*
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
*/
/*
func Boardchar() string {
	piece_dict := map[int]string{
		0: "=" // Empty Spaces
		1: "X" // Player 1 stone
		2: "O" // Player 2 stone
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
*/
/*
func PrintBoard() {
	fmt.Println(Boardchar())
}
*/

func (board Board) ValidMove(x, y int) bool {
	if x > board_dim-1 || x < 0 || y > board_dim-1 || y < 0 {
		return false
	}
	if board[x][y] != 0 {
		return false
	}
	return true
}

func (board *Board) RemoveStone(x, y int) {
	board[x][y] = 0
}


func (board *Board) PlayerMove(player int, x int, y int ) {
	if board.ValidMove(x, y) == false {
		fmt.Println("Choose a Valid Move")
		board.Player_Selection(player)
		return
	}
	board[x][y] = int(player)
	passed_turn = false
	board.Check_Captures(player,x,y)
	return
}

func AlphToBoard(alpha_input string) ([2]int){
	temp := []rune(strings.ToLower(alpha_input))
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
		"m": 12}
	h, _ := horizontal[strings.ToLower(string(temp[0]))]
	v, _ := vertical[strings.ToLower(string(temp[1]))]
	var pos [2]int
	pos[1] = h
	pos[0] = v
	return pos
}