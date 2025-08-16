package main

import (
	"fmt"
	"game/game"
	"sync"
)

func main() {
	// Create channels
	displayChan := make(chan string)
	var wg sync.WaitGroup

	// Initialize game
	g := game.Game{
		DisplayChan: displayChan,
		Round: game.Round{
			UserWin:     0,
			CompWin:     0,
			Draw:        0,
			RoundNumber: 1,
		},
	}

	// Start a goroutine to continuously display messages sent via DisplayChan
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range displayChan {
			fmt.Println(msg)
		}
	}()

	// Show game instructions
	g.DisplayIntro()

	// Game loop - best of 3 rounds
	for g.Round.RoundNumber < 4 {
		g.PlayRound()
		g.Round.RoundNumber++
	}

	// Final result
	displayChan <- "\n--- Final Result ---"
	if g.Round.UserWin < g.Round.CompWin {
		displayChan <- fmt.Sprintf("You lost the game! 😢 Final Score - You: %d, Computer: %d, Draws: %d", g.Round.UserWin, g.Round.CompWin, g.Round.Draw)
	} else if g.Round.UserWin == g.Round.CompWin {
		displayChan <- fmt.Sprintf("It's a draw! 🤝 Final Score - You: %d, Computer: %d, Draws: %d", g.Round.UserWin, g.Round.CompWin, g.Round.Draw)
	} else {
		displayChan <- fmt.Sprintf("You won the game! 🎉 Final Score - You: %d, Computer: %d, Draws: %d", g.Round.UserWin, g.Round.CompWin, g.Round.Draw)
	}
	close(displayChan)
	wg.Wait()
}
