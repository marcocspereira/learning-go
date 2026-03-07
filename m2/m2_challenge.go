// Desafio 2.1: Sistema de Gestão de Jogadores

/*
*
✅ Definir structs com diferentes tipos de dados
✅ Methods com pointer receivers (AddMatch modifica)
✅ Methods com value receivers (cálculos só lêem)
✅ Validações com múltiplas condições
✅ Multiple returns
✅ Error handling
✅ Trabalhar com time.Time
✅ Evitar divisão por zero
*
*/
package main

import (
	"fmt"
	"time"
)

type Player struct {
	ID       int
	Name     string
	Position string // "GK", "DF", "MF", "FW"
	Number   int    // número da camisola (1-99)
	TeamID   int
}

type MatchStats struct {
	ID            int
	PlayerID      int
	MatchDate     time.Time
	Goals         int
	Assists       int
	MinutesPlayed int
	YellowCards   int
	RedCards      int
}

type PlayerSeason struct {
	PlayerID      int
	Season        string // ex: "2023/2024"
	TotalGoals    int
	TotalAssists  int
	MatchesPlayed int
}

func (p Player) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("nome do jogador não pode ser vazio")
	}
	validPositions := map[string]bool{"GK": true, "DF": true, "MF": true, "FW": true}
	if !validPositions[p.Position] {
		return fmt.Errorf("posição inválida: %s", p.Position)
	}
	if p.Number < 1 || p.Number > 99 {
		return fmt.Errorf("número da camisola deve ser entre 1 e 99")
	}
	if p.TeamID <= 0 {
		return fmt.Errorf("TeamID deve ser um inteiro positivo")
	}
	return nil
}

func (ms MatchStats) Validate() error {
	if ms.Goals < 0 {
		return fmt.Errorf("número de gols não pode ser negativo")
	}
	if ms.Assists < 0 {
		return fmt.Errorf("número de assistências não pode ser negativo")
	}
	if ms.MinutesPlayed < 0 || ms.MinutesPlayed > 120 {
		return fmt.Errorf("minutos jogados devem ser entre 0 e 120")
	}
	if ms.RedCards < 0 || ms.RedCards > 1 {
		return fmt.Errorf("número de cartões vermelhos deve ser 0 ou 1")
	}
	if ms.YellowCards < 0 || ms.YellowCards > 2 {
		return fmt.Errorf("número de cartões amarelos deve ser entre 0 e 2")
	}
	if ms.YellowCards == 2 && ms.RedCards == 0 {
		return fmt.Errorf("se um jogador recebe 2 cartões amarelos, ele deve receber 1 	cartão vermelho")
	}
	if ms.PlayerID <= 0 {
		return fmt.Errorf("PlayerID deve ser um inteiro positivo")
	}
	return nil
}

func (ps PlayerSeason) GoalsPerMatch() float64 {
	if ps.MatchesPlayed == 0 {
		return 0.0
	}
	return float64(ps.TotalGoals) / float64(ps.MatchesPlayed)
}

func (ps *PlayerSeason) AddMatch(stats MatchStats) error {
	if stats.PlayerID != ps.PlayerID {
		return fmt.Errorf("PlayerID do MatchStats (%d) não corresponde ao PlayerID do PlayerSeason (%d)", stats.PlayerID, ps.PlayerID)
	}
	if err := stats.Validate(); err != nil {
		return fmt.Errorf("estatísticas inválidas: %w", err)
	}
	ps.TotalGoals += stats.Goals
	ps.TotalAssists += stats.Assists
	ps.MatchesPlayed++
	return nil
}

func CompareScores(ps1, ps2 PlayerSeason) (better string, difference float64) {
	if ps1.TotalGoals > ps2.TotalGoals {
		better = fmt.Sprintf("PlayerID %d", ps1.PlayerID)
		difference = float64(ps1.TotalGoals) - float64(ps2.TotalGoals)
	} else if ps2.TotalGoals > ps1.TotalGoals {
		better = fmt.Sprintf("PlayerID %d", ps2.PlayerID)
		difference = float64(ps2.TotalGoals) - float64(ps1.TotalGoals)
	} else {
		better = "Empate"
		difference = 0
	}
	return
}

func TopScorer(players []PlayerSeason) (*PlayerSeason, error) {
	if len(players) == 0 {
		return nil, fmt.Errorf("nenhum jogador fornecido")
	}
	topIndex := 0
	for i, ps := range players {
		if i == 0 || ps.GoalsPerMatch() > players[topIndex].GoalsPerMatch() {
			topIndex = i
		}
	}
	return &players[topIndex], nil
}

func main() {
	// criar 3 jogadores
	player1 := Player{ID: 1, Name: "Alice", Position: "FW", Number: 9, TeamID: 1}
	player2 := Player{ID: 2, Name: "Bob", Position: "MF", Number: 8, TeamID: 1}

	// validar
	for _, player := range []Player{player1, player2} {
		if err := player.Validate(); err != nil {
			fmt.Printf("Erro ao validar jogador %s: %v\n", player.Name, err)
		} else {
			fmt.Printf("Jogador %s é válido\n", player.Name)
		}
	}

	// criar época
	season := "2023/2024"
	playerSeason1 := PlayerSeason{PlayerID: player1.ID, Season: season}
	playerSeason2 := PlayerSeason{PlayerID: player2.ID, Season: season}

	// adicionar jogo
	match1 := MatchStats{ID: 1, PlayerID: player1.ID, MatchDate: time.Date(2023, 8, 20, 0, 0, 0, 0, time.UTC), Goals: 2, Assists: 1, MinutesPlayed: 90}
	match2 := MatchStats{ID: 2, PlayerID: player2.ID, MatchDate: time.Date(2023, 8, 20, 0, 0, 0, 0, time.UTC), Goals: 1, Assists: 2, MinutesPlayed: 90, YellowCards: 1}

	for _, match := range []MatchStats{match1, match2} {
		if err := match.Validate(); err != nil {
			fmt.Printf("Erro ao validar match stats para PlayerID %d: %v\n", match.PlayerID, err)
		} else {
			fmt.Printf("Match stats para PlayerID %d são válidos\n", match.PlayerID)
		}
		// adicionar ao player season
		switch match.PlayerID {
		case player1.ID:
			playerSeason1.AddMatch(match)
		case player2.ID:
			playerSeason2.AddMatch(match)
		}
	}

	// imprimir resultados
	fmt.Printf("Player %s - Goals per match: %.2f\n", player1.Name, playerSeason1.GoalsPerMatch())
	fmt.Printf("Player %s - Goals per match: %.2f\n", player2.Name, playerSeason2.GoalsPerMatch())

	better, diff := CompareScores(playerSeason1, playerSeason2)
	fmt.Printf("Comparação de scores: %s é melhor com diferença de %.2f gols por jogo\n", better, float64(diff))
}
