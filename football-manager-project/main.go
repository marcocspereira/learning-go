package main

import (
	"fmt"
	"time"
)

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

type Player struct {
	ID       int
	Name     string
	Position string // "GK", "DF", "MF", "FW"
	Number   int    // número da camisola (1-99)
	TeamID   int
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

type Team struct {
	ID      int
	Name    string
	Stadium string
	Founded int
}

func (t Team) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("nome da equipa não pode ser vazio")
	}
	if t.Stadium == "" {
		return fmt.Errorf("estádio da equipa não pode ser vazio")
	}
	if t.Founded < 1850 || t.Founded > time.Now().Year() {
		return fmt.Errorf("ano de fundação deve ser entre 1850 e %d", time.Now().Year())
	}
	return nil
}

type Match struct {
	ID         int
	HomeTeamID int // TeamID
	AwayTeamID int // TeamID
	HomeScore  int
	AwayScore  int
	MatchDate  time.Time
	Status     string // "scheduled", "ongoing", "finished"
}

func (m Match) Validate() error {
	if m.HomeTeamID == m.AwayTeamID {
		return fmt.Errorf("HomeTeamID e AwayTeamID não podem ser iguais")
	}
	if m.HomeScore < 0 || m.AwayScore < 0 {
		return fmt.Errorf("HomeScore e AwayScore não podem ser negativos")
	}
	validStatuses := map[string]bool{"scheduled": true, "ongoing": true, "finished": true}
	if !validStatuses[m.Status] {
		return fmt.Errorf("Status inválido: %s", m.Status)
	}
	if m.MatchDate.Before(time.Now()) && m.Status == "scheduled" {
		return fmt.Errorf("Partida agendada para data passada")
	}
	return nil
}

type PlayerMatchStats struct {
	ID            int
	PlayerID      int
	MatchID       int
	Goals         int
	Assists       int
	MinutesPlayed int
	YellowCards   int
	RedCards      int
}

func (pms PlayerMatchStats) Validate() error {
	if pms.PlayerID <= 0 {
		return fmt.Errorf("PlayerID deve ser um inteiro positivo")
	}
	if pms.MatchID <= 0 {
		return fmt.Errorf("MatchID deve ser um inteiro positivo")
	}
	if pms.Goals < 0 {
		return fmt.Errorf("número de gols não pode ser negativo")
	}
	if pms.Assists < 0 {
		return fmt.Errorf("número de assistências não pode ser negativo")
	}
	if pms.MinutesPlayed < 0 || pms.MinutesPlayed > 120 {
		return fmt.Errorf("MinutesPlayed deve ser entre 0 e 120")
	}
	if pms.YellowCards < 0 || pms.YellowCards > 2 {
		return fmt.Errorf("YellowCards deve ser entre 0 e 2")
	}
	if pms.RedCards < 0 || pms.RedCards > 1 {
		return fmt.Errorf("RedCards deve ser entre 0 e 1")
	}
	return nil
}
