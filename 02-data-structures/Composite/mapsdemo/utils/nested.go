package utils

import "encoding/json"

type DeepScores map[string]map[string]map[string]int

// NewDeepScores creates a new DeepScores structure.
func NewDeepScores() DeepScores {
	return make(DeepScores)
}

// AddScore safely inserts/updates a deeply nested score.
func (d DeepScores) AddScore(student, term, subject string, score int) {
	if d[student] == nil {
		d[student] = make(map[string]map[string]int)
	}
	if d[student][term] == nil {
		d[student][term] = make(map[string]int)
	}
	d[student][term][subject] = score
}

// ToJSON returns the nested structure as JSON.
func (d DeepScores) ToJSON() (string, error) {
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

