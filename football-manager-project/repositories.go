package main

// Interfaces para repositórios (contratos)

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
