package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(teamA, teamB string, scoreA, scoreB int) {
	switch {
	case scoreA > scoreB:
		l.Wins[teamA]++
	case scoreB > scoreA:
		l.Wins[teamB]++
	}
}

func (l *League) Ranking() []string {
	teams := []string{}
	for _, team := range l.Teams {
		teams = append(teams, team.name)
	}
	sort.Slice(teams, func(i, j int) bool {
		return l.Wins[teams[i]] > l.Wins[teams[j]]
	})

	return teams
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()
	for _, team := range ranking {
		w.Write([]byte(team + "\n"))
	}
}

func main() {
	league := League{
		Teams: []Team{
			{
				name:    "Mumbai Indians",
				players: []string{"p1", "p2", "p3", "p4", "p5", "p6"},
			},
			{
				name:    "Chennai Super Kings",
				players: []string{"p1", "p2", "p3", "p4", "p5", "p6"},
			},
			{
				name:    "Royal Challengers Bengaluru",
				players: []string{"p1", "p2", "p3", "p4", "p5", "p6"},
			},
			{
				name:    "Gujarat Titans",
				players: []string{"p1", "p2", "p3", "p4", "p5", "p6"},
			},
			{
				name:    "Sunrisers Hyderabad",
				players: []string{"p1", "p2", "p3", "p4", "p5", "p6"},
			},
			{
				name:    "Delhi Daredevils",
				players: []string{"p1", "p2", "p3", "p4", "p5", "p6"},
			},
		},
		Wins: map[string]int{
			"Mumbai Indians":              2,
			"Delhi Daredevils":            1,
			"Sunrisers Hyderabad":         5,
			"Gujarat Titans":              8,
			"Royal Challengers Bengaluru": 3,
			"Chennai Super Kings":         2,
		},
	}

	RankPrinter(&league, os.Stdout)

	animalFunc()
}
