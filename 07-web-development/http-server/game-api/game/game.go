package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	QUIT     = 3
)

type RoundResult struct {
	UserChoose     string `json:"userchoose"`
	ComputerChoose string `json:"computerchoose"`
	Result         string `json:"result"`
	Result_Message string `json:"resultmessage"`
}

func moveToString(move int) string {
	switch move {
	case ROCK:
		return "Rock"
	case PAPER:
		return "Paper"
	case SCISSORS:
		return "Scissors"
	case QUIT:
		return "Quit"
	default:
		return "Invalid Move"
	}
}

func PlayRound(c int) RoundResult {
	// Computer's move can be randomly generated in a complete implementation
	comp_randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	computerMove := comp_randomSeed.Intn(3)
	var win, resultmessage string = " ", " "
	userMove := c
	if (userMove == 0 && computerMove == 2) ||
		(userMove == 2 && computerMove == 1) ||
		(userMove == 1 && computerMove == 0) {
		fmt.Printf("You win! %s crushes %s \n", moveToString(userMove), moveToString(computerMove))
		win = "You won"
		resultmessage = "Wonderful choice you won"
	} else if userMove == computerMove {
		fmt.Println("It's a draw!")
		win = "Draw"
		resultmessage = "Lets win next time"
	} else {
		fmt.Printf("Computer wins! %s crushes %s \n", moveToString(computerMove), moveToString(userMove))
		win = "Computer won"
		resultmessage = "Keep trying lets try once again"
	}
	var result RoundResult
	result.UserChoose = moveToString(userMove)
	result.Result = win
	result.ComputerChoose = moveToString(computerMove)
	result.Result_Message = resultmessage
	return result
}
