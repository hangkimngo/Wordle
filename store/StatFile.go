package store

import (
	"encoding/csv"
	"os"
	"strconv"
)

type GameStat struct {
	Username string
	Secret   string
	Attempts int
	Outcome  string // "win" or "loss"
}

type UserStats struct {
	Played      int
	Won         int
	AvgAttempts float64
}

// Append one row: username,secret word,attempts,win/loss
func AppendGameStat(path string, s GameStat) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	return w.Write([]string{
		s.Username,
		s.Secret,
		strconv.Itoa(s.Attempts),
		s.Outcome,
	})
}

// Read stats for a username and compute totals
func ReadUserStats(path, username string) (UserStats, error) {
	f, err := os.Open(path)
	if err != nil {
		// If file doesn't exist yet, treat as 0 stats (don’t crash tests)
		if os.IsNotExist(err) {
			return UserStats{Played: 0, Won: 0, AvgAttempts: 0}, nil
		}
		return UserStats{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	played := 0
	won := 0
	sumAttempts := 0

	for {
		rec, err := r.Read()
		if err != nil {
			break
		}
		if len(rec) < 4 {
			continue
		}
		if rec[0] != username {
			continue
		}
		played++
		a, _ := strconv.Atoi(rec[2])
		sumAttempts += a
		if rec[3] == "win" {
			won++
		}
	}

	if played == 0 {
		return UserStats{Played: 0, Won: 0, AvgAttempts: 0}, nil
	}
	return UserStats{
		Played:      played,
		Won:         won,
		AvgAttempts: float64(sumAttempts) / float64(played),
	}, nil
}
