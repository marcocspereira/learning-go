package main

import (
	"fmt"
	"time"
)

type PlayerStats struct {
	PlayerID   int
	TotalGoals int
	Error      error
}

func calculatePlayerStats(playerID int) PlayerStats {
	// Simula cálculo demorado
	time.Sleep(1 * time.Second)

	// Calcula stats...
	return PlayerStats{
		PlayerID:   playerID,
		TotalGoals: 10, // exemplo
	}
}

func main() {
	playerIDs := []int{1, 2, 3, 4, 5}

	// Channel para receber resultados
	results := make(chan PlayerStats, len(playerIDs)) // buffered channel

	// Lançar goroutines
	for _, id := range playerIDs {
		go func(playerID int) {
			stats := calculatePlayerStats(playerID)
			results <- stats // envia resultado
		}(id)
	}

	// Recolher resultados
	for i := 0; i < len(playerIDs); i++ {
		stats := <-results
		fmt.Printf("Player %d: %d golos\n", stats.PlayerID, stats.TotalGoals)
	}

	// Sequencial: 5 segundos (1s × 5)
	// Concorrente: 1 segundo (todos ao mesmo tempo!)
}
