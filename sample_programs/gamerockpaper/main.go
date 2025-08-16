package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	QUIT     = 3
)

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

func main() {

	// Computer's move can be randomly generated in a complete implementation
	comp_randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize scores
	var userWin, computerWin, draw = 0, 0, 0

	// Welcome message and instructions
	fmt.Println("Welcome to the Rock-Paper-Scissors Game with Computer!")
	fmt.Println("Please select your move:")
	fmt.Println("0: Rock")
	fmt.Println("1: Paper")
	fmt.Println("2: Scissors")
	fmt.Println("3 to quit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your move: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		userMove, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid move. Please select 0, 1, 2, or 3 to quit. (%v)\n", err)
			continue
		}

		fmt.Printf("You selected: %s\n", moveToString(userMove))

		if userMove < 0 || userMove > 3 {
			fmt.Println("Invalid move. Please select 0, 1, 2, or 3 to quit.")
		} else if userMove != 3 {
			computerMove := comp_randomSeed.Intn(3)
			fmt.Printf("Computer's move:%s\n", moveToString(computerMove))
			if (userMove == 0 && computerMove == 2) ||
				(userMove == 2 && computerMove == 1) ||
				(userMove == 1 && computerMove == 0) {
				fmt.Printf("You win! %s crushes %s \n", moveToString(userMove), moveToString(computerMove))
				userWin++
			} else if userMove == computerMove {
				fmt.Println("It's a draw!")
				draw++
			} else {
				fmt.Printf("Computer wins! %s crushes %s \n", moveToString(computerMove), moveToString(userMove))
				computerWin++
			}
		} else if userMove == 3 {
			fmt.Println("Exiting the game.")
			break
		}
	}

	fmt.Println("Game Over ! and the final score is:")
	fmt.Printf("User Wins: %d, Computer Wins: %d, Draws: %d\n", userWin, computerWin, draw)
}
