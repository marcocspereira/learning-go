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

// Intercaces para repositórios (contratos)
type PlayerRepository interface {
	Create(player Player) (Player, error) // retorna o player criado com ID preenchido
	FindByID(id int) (Player, error)
	FindByTeam(teamID int) ([]Player, error)
	Update(player Player) error
	Delete(id int) error
}

type TeamRepository interface {
	Create(team Team) (Team, error)
	FindByID(id int) (Team, error)
	FindAll() ([]Team, error)
	Update(team Team) error
	Delete(id int) error
}

type MatchRepository interface {
	Create(match Match) (Match, error)
	FindByID(id int) (Match, error)
	FindByTeam(teamID int) ([]Match, error)
	Update(match Match) error
}

type StatusRepository interface {
	Create(stats PlayerMatchStats) (PlayerMatchStats, error)
	FindByMatch(matchID int) ([]PlayerMatchStats, error)
	FindByPlayer(playerID int) ([]PlayerMatchStats, error)
}

// Mock implementations em memória (para testes)

type InMemoryPlayerRepo struct {
	players map[int]Player
	nextID  int
}

func NewInMemoryPlayerRepo() *InMemoryPlayerRepo {
	return &InMemoryPlayerRepo{
		players: make(map[int]Player),
		nextID:  1,
	}
}

func (repo *InMemoryPlayerRepo) Create(player Player) (Player, error) {
	if err := player.Validate(); err != nil {
		return Player{}, err
	}
	player.ID = repo.nextID
	repo.players[repo.nextID] = player
	repo.nextID++
	return player, nil
}

func (repo *InMemoryPlayerRepo) FindByID(id int) (Player, error) {
	player, exists := repo.players[id]
	if !exists {
		return Player{}, fmt.Errorf("Player with ID %d not found", id)
	}
	return player, nil
}

func (repo *InMemoryPlayerRepo) FindByTeam(teamID int) ([]Player, error) {
	var result []Player
	for _, player := range repo.players {
		if player.TeamID == teamID {
			result = append(result, player)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("No players found for TeamID %d", teamID)
	}
	return result, nil
}

func (repo *InMemoryPlayerRepo) Update(player Player) error {
	if err := player.Validate(); err != nil {
		return err
	}
	_, exists := repo.players[player.ID]
	if !exists {
		return fmt.Errorf("Player with ID %d not found", player.ID)
	}
	repo.players[player.ID] = player
	return nil
}

func (repo *InMemoryPlayerRepo) Delete(id int) error {
	_, exists := repo.players[id]
	if !exists {
		return fmt.Errorf("Player with ID %d not found", id)
	}
	delete(repo.players, id)
	return nil
}

// InMemoryTeamRepo

type InMemoryTeamRepo struct {
	teams  map[int]Team
	nextID int
}

func NewInMemoryTeamRepo() *InMemoryTeamRepo {
	return &InMemoryTeamRepo{
		teams:  make(map[int]Team),
		nextID: 1,
	}
}

func (repo *InMemoryTeamRepo) Create(team Team) (Team, error) {
	if err := team.Validate(); err != nil {
		return Team{}, err
	}
	team.ID = repo.nextID
	repo.teams[repo.nextID] = team
	repo.nextID++
	return team, nil
}

func (repo *InMemoryTeamRepo) FindByID(id int) (Team, error) {
	team, exists := repo.teams[id]
	if !exists {
		return Team{}, fmt.Errorf("Team with ID %d not found", id)
	}
	return team, nil
}

func (repo *InMemoryTeamRepo) FindAll() ([]Team, error) {
	result := make([]Team, 0, len(repo.teams))
	for _, team := range repo.teams {
		result = append(result, team)
	}
	return result, nil
}

func (repo *InMemoryTeamRepo) Update(team Team) error {
	if err := team.Validate(); err != nil {
		return err
	}
	_, exists := repo.teams[team.ID]
	if !exists {
		return fmt.Errorf("Team with ID %d not found", team.ID)
	}
	repo.teams[team.ID] = team
	return nil
}

func (repo *InMemoryTeamRepo) Delete(id int) error {
	_, exists := repo.teams[id]
	if !exists {
		return fmt.Errorf("Team with ID %d not found", id)
	}
	delete(repo.teams, id)
	return nil
}

// InMemoryMatchRepo

type InMemoryMatchRepo struct {
	matches map[int]Match
	nextID  int
}

func NewInMemoryMatchRepo() *InMemoryMatchRepo {
	return &InMemoryMatchRepo{
		matches: make(map[int]Match),
		nextID:  1,
	}
}

func (repo *InMemoryMatchRepo) Create(match Match) (Match, error) {
	if err := match.Validate(); err != nil {
		return Match{}, err
	}
	match.ID = repo.nextID
	repo.matches[repo.nextID] = match
	repo.nextID++
	return match, nil
}

func (repo *InMemoryMatchRepo) FindByID(id int) (Match, error) {
	match, exists := repo.matches[id]
	if !exists {
		return Match{}, fmt.Errorf("Match with ID %d not found", id)
	}
	return match, nil
}

func (repo *InMemoryMatchRepo) FindByTeam(teamID int) ([]Match, error) {
	var result []Match
	for _, match := range repo.matches {
		if match.HomeTeamID == teamID || match.AwayTeamID == teamID {
			result = append(result, match)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("No matches found for TeamID %d", teamID)
	}
	return result, nil
}

func (repo *InMemoryMatchRepo) Update(match Match) error {
	if err := match.Validate(); err != nil {
		return err
	}
	_, exists := repo.matches[match.ID]
	if !exists {
		return fmt.Errorf("Match with ID %d not found", match.ID)
	}
	repo.matches[match.ID] = match
	return nil
}

// InMemoryStatsRepo

type InMemoryStatsRepo struct {
	stats  map[int]PlayerMatchStats
	nextID int
}

func NewInMemoryStatsRepo() *InMemoryStatsRepo {
	return &InMemoryStatsRepo{
		stats:  make(map[int]PlayerMatchStats),
		nextID: 1,
	}
}

func (repo *InMemoryStatsRepo) Create(stats PlayerMatchStats) (PlayerMatchStats, error) {
	if err := stats.Validate(); err != nil {
		return PlayerMatchStats{}, err
	}
	stats.ID = repo.nextID
	repo.stats[repo.nextID] = stats
	repo.nextID++
	return stats, nil
}

func (repo *InMemoryStatsRepo) FindByMatch(matchID int) ([]PlayerMatchStats, error) {
	var result []PlayerMatchStats
	for _, s := range repo.stats {
		if s.MatchID == matchID {
			result = append(result, s)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("No stats found for MatchID %d", matchID)
	}
	return result, nil
}

func (repo *InMemoryStatsRepo) FindByPlayer(playerID int) ([]PlayerMatchStats, error) {
	var result []PlayerMatchStats
	for _, s := range repo.stats {
		if s.PlayerID == playerID {
			result = append(result, s)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("No stats found for PlayerID %d", playerID)
	}
	return result, nil
}

func main() {
	fmt.Println("Football Manager - sistema iniciado")
}
