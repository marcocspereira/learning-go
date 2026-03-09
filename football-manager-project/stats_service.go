package main

import "fmt"

type StatsService struct {
	statsRepo  StatsRepository
	playerRepo PlayerRepository
	matchRepo  MatchRepository
}

func NewStatsService(statsRepo StatsRepository, playerRepo PlayerRepository, matchRepo MatchRepository) *StatsService {
	return &StatsService{statsRepo: statsRepo, playerRepo: playerRepo, matchRepo: matchRepo}
}

func (s *StatsService) AddStats(stats PlayerMatchStats) (PlayerMatchStats, error) {
	if err := stats.Validate(); err != nil {
		return PlayerMatchStats{}, err
	}
	if _, err := s.playerRepo.FindByID(stats.PlayerID); err != nil {
		return PlayerMatchStats{}, fmt.Errorf("player with ID %d not found", stats.PlayerID)
	}
	match, err := s.matchRepo.FindByID(stats.MatchID)
	if err != nil {
		return PlayerMatchStats{}, fmt.Errorf("match with ID %d not found", stats.MatchID)
	}
	if match.Status == "scheduled" {
		return PlayerMatchStats{}, fmt.Errorf("cannot add stats to a scheduled match")
	}
	return s.statsRepo.Create(stats)
}

func (s *StatsService) GetMatchStats(matchID int) ([]PlayerMatchStats, error) {
	if _, err := s.matchRepo.FindByID(matchID); err != nil {
		return nil, fmt.Errorf("match with ID %d not found", matchID)
	}
	return s.statsRepo.FindByMatch(matchID)
}

func (s *StatsService) GetPlayerStats(playerID int) ([]PlayerMatchStats, error) {
	if _, err := s.playerRepo.FindByID(playerID); err != nil {
		return nil, fmt.Errorf("player with ID %d not found", playerID)
	}
	return s.statsRepo.FindByPlayer(playerID)
}
