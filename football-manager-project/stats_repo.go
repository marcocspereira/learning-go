package main

import "fmt"

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
