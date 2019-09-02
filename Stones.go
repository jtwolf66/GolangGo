package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var playing bool

var player int

type Neighborhood struct{
	Type []int
	Pos [][2]int
}

type PosList [][2]int

func Opponent(player int) int {
	var opponent int
	if player == 1 {
		opponent = 2
	}
	if player == 2 {
		opponent = 1
	}
	return opponent
}

func SwitchPlayer(player_turn int) int {
	var newplayer int
	if player_turn == 1 {
		newplayer = 2
	}
	if player_turn == 2 {
		newplayer = 1
	}
	return newplayer
}

func (board *Board) GetNeighbors(x int, y int) Neighborhood {
	/*
	Returns neighbors in the form (Ltype,Rtype,Utype,Dtype),((Lx,Ly),(Rx,Ry),(Ux,Uy),(Dx,Dy))
	Type options:
	edge - 3  (This is 3 simply because 3 is not used yet elsewhere, and is less likely to get confused
	player - Players #
	opponent - Opps #
	free - 0
	*/
	var neighbors Neighborhood
	itercheck := [2]int{-1,1}
	iterchecky := [2]int{1,-1}

	for _, elemx := range itercheck{
		if (board_dim - 1) >= (x + elemx) && (x + elemx) >= 0 {
			neighbors.Type = append(neighbors.Type, board[x+elemx][y])
			var pos [2]int
			pos[0] = x+elemx
			pos[1] = y
			neighbors.Pos = append(neighbors.Pos, pos)
		}
	}
	for _, elemy := range iterchecky{
		if (board_dim - 1) >= (y + elemy) && (y + elemy) >= 0 {
			neighbors.Type = append(neighbors.Type, board[x][y+elemy])
			var pos [2]int
			pos[0] = x
			pos[1] = y + elemy
			neighbors.Pos = append(neighbors.Pos, pos)
		}
	}
	/*
	// Handle edges/corners
	if x == 0 {
		Type[0] = 3
		Loc[0][0] = 99
		Loc[0][1] = 99
	}
	if x == board_dim - 1 {
		Type[1] = 3
		Loc[1][0] = 99
		Loc[1][1] = 99
	}
	if y == 0 {
		Type[2] = 3
		Loc[2][1] = 99
		Loc[2][1] = 99
	}
	if y == board_dim - 1 {
		Type[3] = 3
		Loc [3][1] = 99
		Loc [3][1] = 99
	}
	*/
	return neighbors
}
func (board *Board) Get_Liberties(player int, x int, y int, nodes PosList) (PosList,PosList) {
	var posl PosList
	var pos [2]int

	// Add node to collection.
	var currnode [2]int
	currnode[0] = x
	currnode[1] = y
	nodes = append(nodes, currnode)

	nbrs := board.GetNeighbors(x,y)

	for i := 0; i < len(nbrs.Type); i++ {
		if nbrs.Type[i] == 0 {
			pos[0] = x
			pos[1] = y
			posl = append(posl, pos)
			}
		if nbrs.Type[i] == int(player) {
			sol, nds := board.Get_Liberties(player, x, y,nodes)
			for i := 0; i<len(sol); i++ {
				posl = append(posl, sol[i])
				nodes = append(nodes,nds[i])
			}

		}
	}
	return posl, nodes
}

func (board *Board) Check_Captures(player int ,x int ,y int) {
	nbrs := board.GetNeighbors(x,y)
	for i := 0; i < len(nbrs.Type); i++ {
		if nbrs.Type[i] == int(Opponent(player)) {
			var nodes PosList
			lib, connected := board.Get_Liberties(Opponent(player),nbrs.Pos[i][0],nbrs.Pos[i][1], nodes)
			if len(lib) == 0 {
				for i := 0; i < len(nodes); i++ {
					board.RemoveStone(connected[i][0],connected[i][1])
				}

			}
		}
	}

	//Will check if the given move has captured any spaces and alter board as necessary
}

/* Will implement scoring later
func (board Board) Check_Points() (bool, [2] int, [2] int) {
	//Will count points at end of match
	//Returns Draw? (bool), player (victor,loser) and points (victor, loser)
	return
}
*/
func (board *Board) Player_Selection(player int) int {
	// Listens for moves, checks captures, switches player

	board.Print_Str_Form()
	input_reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Player %d's move (ex: a5): \n",player)
	input, _ := input_reader.ReadString('\n')
	fmt.Println()

	// End game if the opponent quits
	if strings.HasPrefix(strings.ToLower(input), "quit") {
		playing = false
		fmt.Printf("Player %d resigns. Opponent wins!", player)
		return 3
	}
	// Check if both players have skipped their turn, ending the game
	if strings.HasPrefix(strings.ToLower(input), "pass") {
		if passed_turn == true {
			playing = false
			/* draw, vic, points := Check_Points()
			if draw == true {
				fmt.Printf("Both Players have passed on their turn. Players 1 and 2 draw with %d points", points[0])
				return
			}
			if draw == false {
				fmt.Printf("Both Players have passed on their turn. Player %d wins with %d points. Player %d loses with %d points",)
				return
			}
			*/

		}
	}
	move := AlphToBoard(input)
	board.PlayerMove(player, move[0], move[1])


	// Switch player state
	player = SwitchPlayer(player)
	return player
}


func main() {
	fmt.Println("Welcome to Golang Go!")
	fmt.Println("Specify where you wish to place to piece.")
	fmt.Println("If you wish to skip your turn, type 'pass'")
	fmt.Println("Once both players skip their turns or no more moves are available, the game ends")
	fmt.Println("Good luck!")

	var player int = 1
	board := GenerateBoard()
	playing = true
	for playing {
		player = board.Player_Selection(player)
	}
}