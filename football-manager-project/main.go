package main

import (
	"fmt"
	"time"
)

// Run: go run .

// 🎯 Mini-Projeto: Sistema de Gestão de Futebol

/**
O que vais construir:
Um sistema para gerir:

Players (jogadores)
Teams (equipas)
Matches (jogos)
Stats (estatísticas de jogos)

Scope desta sessão:
✅ Structs (data models)
✅ Interfaces (contracts para repositories)
✅ Validações (methods de validação)
✅ Mock repositories (implementação em memória)
✅ Services (lógica de negócio básica)
❌ NÃO vais fazer (ainda):

Database real (Fase 6/7)
HTTP endpoints (Fase 7)
Concorrência (Fase 5)
Error handling avançado (Fase 4)
*/

func main() {
	// Criar repositórios em memória
	playerRepo := NewInMemoryPlayerRepo()
	teamRepo := NewInMemoryTeamRepo()
	matchRepo := NewInMemoryMatchRepo()
	statsRepo := NewInMemoryStatsRepo()

	fmt.Println("Football Manager - sistema iniciado")
	fmt.Printf("Repos inicializados: players=%T, teams=%T, matches=%T, stats=%T\n",
		playerRepo, teamRepo, matchRepo, statsRepo)

	// Criar serviços
	playerService := NewPlayerService(playerRepo, teamRepo)
	matchService := NewMatchService(matchRepo, teamRepo)

	// Teste 1: criar equipas
	fmt.Println("=== Criar Equipas ===")
	benfica, _ := teamRepo.Create(Team{Name: "Benfica", Stadium: "Estádio da Luz", Founded: 1904})
	porto, _ := teamRepo.Create(Team{Name: "Porto", Stadium: "Estádio do Dragão", Founded: 1893})
	fmt.Printf("Equipas criadas: %v, %v\n", benfica, porto)

	// Teste 2: criar jogadores
	fmt.Println("=== Criar Jogadores ===")
	jogador1, _ := playerService.CreatePlayer(Player{Name: "João Silva", Position: "MF", Number: 8, TeamID: benfica.ID})
	jogador2, _ := playerService.CreatePlayer(Player{Name: "Carlos Pereira", Position: "FW", Number: 9, TeamID: porto.ID})
	fmt.Printf("Jogadores criados: %v, %v\n", jogador1, jogador2)

	// Teste 3: criar jogadores em equipa inexistente (deve falhar)
	fmt.Println("=== Criar Jogador em Equipa Inexistente (deve falhar) ===")
	_, err := playerService.CreatePlayer(Player{Name: "Miguel Santos", Position: "DF", Number: 5, TeamID: 999})
	if err != nil {
		fmt.Printf("Erro esperado ao criar jogador em equipa inexistente: %v\n", err)
	} else {
		fmt.Println("Erro: jogador criado com TeamID inexistente (deveria falhar)")
	}

	// Teste 4: criar jogo
	fmt.Println("=== Criar Jogo ===")
	jogo, error := matchService.CreateMatch(Match{ID: 1, HomeTeamID: benfica.ID, AwayTeamID: porto.ID, HomeScore: 0, AwayScore: 0, MatchDate: time.Now().Add(24 * time.Hour), Status: "scheduled"})
	if error != nil {
		fmt.Printf("Erro ao criar jogo: %v\n", error)
	}
	fmt.Printf("Jogo criado: %v\n", jogo)
}
