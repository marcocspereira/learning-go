package main

import "fmt"

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
	playerRepo := NewInMemoryPlayerRepo()
	teamRepo := NewInMemoryTeamRepo()
	matchRepo := NewInMemoryMatchRepo()
	statsRepo := NewInMemoryStatsRepo()

	fmt.Println("Football Manager - sistema iniciado")
	fmt.Printf("Repos inicializados: players=%T, teams=%T, matches=%T, stats=%T\n",
		playerRepo, teamRepo, matchRepo, statsRepo)
}
