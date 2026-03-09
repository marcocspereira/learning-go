package main

import "fmt"

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
