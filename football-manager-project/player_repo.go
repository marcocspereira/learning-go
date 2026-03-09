package main

import "fmt"

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
