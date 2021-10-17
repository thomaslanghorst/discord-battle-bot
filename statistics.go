package main

import "sort"

type Statistics struct {
	wins  map[string]uint64
	loses map[string]uint64
}

type Statistic struct {
	Username string
	Wins     uint64
	Loses    uint64
}

func NewStatistics() *Statistics {
	return &Statistics{
		wins:  make(map[string]uint64),
		loses: make(map[string]uint64),
	}
}

func (s *Statistics) AddStatistic(winner, loser string) {

	if val, ok := s.wins[winner]; ok {
		s.wins[winner] = val + 1
	} else {
		s.wins[winner] = 1
	}

	if val, ok := s.loses[loser]; ok {
		s.loses[loser] = val + 1
	} else {
		s.loses[loser] = 1
	}
}

func (s *Statistics) GetStatistics() []Statistic {

	stats := make(map[string]Statistic, 0)

	for username, winCount := range s.wins {

		if _, ok := stats[username]; ok {
			stats[username] = Statistic{
				Username: username,
				Wins:     winCount,
				Loses:    0,
			}
		} else {
			stat := stats[username]
			stat.Username = username
			stat.Wins = winCount
			stats[username] = stat
		}
	}

	for username, loseCount := range s.loses {

		if _, ok := stats[username]; ok {
			stats[username] = Statistic{
				Username: username,
				Wins:     0,
				Loses:    loseCount,
			}
		} else {
			stat := stats[username]
			stat.Loses = loseCount
			stat.Username = username
			stats[username] = stat
		}
	}

	// only return values from stats, not the keys
	result := make([]Statistic, 0)

	for _, v := range stats {
		result = append(result, v)
	}

	// sort result for highest wins
	sort.Slice(result, func(i, j int) bool {
		return result[i].Wins > result[j].Wins
	})

	return result
}
