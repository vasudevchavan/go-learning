package game

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

// Converts move integer to string
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
		return "Invalid"
	}
}

// Game struct
type Game struct {
	DisplayChan chan string
	Round       Round
}

// Round struct to store scores and current round
type Round struct {
	UserWin     int
	CompWin     int
	Draw        int
	RoundNumber int
}

// Displays intro message
func (g *Game) DisplayIntro() {
	g.DisplayChan <- "Welcome to the Rock-Paper-Scissors Game with Computer!"
	g.DisplayChan <- "Best of 3 rounds wins the game."
	g.DisplayChan <- "Please select your move:"
	g.DisplayChan <- "0: Rock"
	g.DisplayChan <- "1: Paper"
	g.DisplayChan <- "2: Scissors"
	g.DisplayChan <- "3: Quit"
}

// Plays one round of the game
func (g *Game) PlayRound() {
	compRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	scanner := bufio.NewScanner(os.Stdin)

	g.DisplayChan <- fmt.Sprintf("\n--- Round %d ---", g.Round.RoundNumber)
	fmt.Print("Enter your move: ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	userMove, err := strconv.Atoi(input)
	if err != nil {
		g.DisplayChan <- fmt.Sprintf("Invalid input: %v. Please enter 0, 1, 2, or 3.", err)
		return
	}

	if userMove < 0 || userMove > 3 {
		g.DisplayChan <- "Invalid move. Please select 0, 1, 2, or 3 to quit."
		return
	}

	if userMove == QUIT {
		g.DisplayChan <- "You chose to quit the game. Goodbye!"
		os.Exit(0)
	}

	computerMove := compRand.Intn(3)

	g.DisplayChan <- fmt.Sprintf("You selected: %s", moveToString(userMove))
	g.DisplayChan <- fmt.Sprintf("Computer selected: %s", moveToString(computerMove))

	// Determine outcome
	if userMove == computerMove {
		g.DisplayChan <- "It's a draw!"
		g.Round.Draw++
	} else if (userMove == ROCK && computerMove == SCISSORS) ||
		(userMove == PAPER && computerMove == ROCK) ||
		(userMove == SCISSORS && computerMove == PAPER) {
		g.DisplayChan <- fmt.Sprintf("You win this round! %s beats %s.", moveToString(userMove), moveToString(computerMove))
		g.Round.UserWin++
	} else {
		g.DisplayChan <- fmt.Sprintf("Computer wins this round! %s beats %s.", moveToString(computerMove), moveToString(userMove))
		g.Round.CompWin++
	}
}
