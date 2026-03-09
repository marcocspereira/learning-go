package main

import "fmt"

type PlayerService struct {
	playerRepo PlayerRepository
	teamRepo   TeamRepository
}

func NewPlayerService(playerRepo PlayerRepository, teamRepo TeamRepository) *PlayerService {
	return &PlayerService{playerRepo: playerRepo, teamRepo: teamRepo}
}

func (s *PlayerService) CreatePlayer(player Player) (Player, error) {
	if _, err := s.teamRepo.FindByID(player.TeamID); err != nil {
		return Player{}, fmt.Errorf("team with ID %d not found", player.TeamID)
	}
	return s.playerRepo.Create(player)
}

func (s *PlayerService) GetTeamPlayers(teamID int) ([]Player, error) {
	if _, err := s.teamRepo.FindByID(teamID); err != nil {
		return nil, fmt.Errorf("team with ID %d not found", teamID)
	}
	return s.playerRepo.FindByTeam(teamID)
}
