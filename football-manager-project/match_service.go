package main

import "fmt"

type MatchService struct {
	matchRepo MatchRepository
	teamRepo  TeamRepository
}

func NewMatchService(matchRepo MatchRepository, teamRepo TeamRepository) *MatchService {
	return &MatchService{matchRepo: matchRepo, teamRepo: teamRepo}
}

func (s *MatchService) CreateMatch(match Match) (Match, error) {
	if err := match.Validate(); err != nil {
		return Match{}, err
	}
	if _, err := s.teamRepo.FindByID(match.HomeTeamID); err != nil {
		return Match{}, fmt.Errorf("home team with ID %d not found", match.HomeTeamID)
	}
	if _, err := s.teamRepo.FindByID(match.AwayTeamID); err != nil {
		return Match{}, fmt.Errorf("away team with ID %d not found", match.AwayTeamID)
	}
	return s.matchRepo.Create(match)
}

func (s *MatchService) GetTeamMatches(teamID int) ([]Match, error) {
	if _, err := s.teamRepo.FindByID(teamID); err != nil {
		return nil, fmt.Errorf("team with ID %d not found", teamID)
	}
	return s.matchRepo.FindByTeam(teamID)
}

func (s *MatchService) UpdateScore(matchID, homeScore, awayScore int) (Match, error) {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return Match{}, fmt.Errorf("match with ID %d not found", matchID)
	}
	if match.Status == "finished" {
		return Match{}, fmt.Errorf("cannot update score of a finished match")
	}
	if homeScore < 0 || awayScore < 0 {
		return Match{}, fmt.Errorf("scores cannot be negative")
	}
	match.HomeScore = homeScore
	match.AwayScore = awayScore
	match.Status = "ongoing"
	if err := s.matchRepo.Update(match); err != nil {
		return Match{}, err
	}
	return match, nil
}
