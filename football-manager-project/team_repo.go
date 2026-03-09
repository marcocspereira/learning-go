package main

import "fmt"

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
