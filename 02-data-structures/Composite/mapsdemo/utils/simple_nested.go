package utils

import (
	"encoding/json"
	"sort"
)

type StudentScores map[string]map[string]int

// NewStudentScores initializes and returns a new StudentScores map.
func NewStudentScores() StudentScores {
	return make(StudentScores)
}

// AddScore safely adds or updates a student's subject score.
func (s StudentScores) AddScore(student, subject string, score int) {
	if s[student] == nil {
		s[student] = make(map[string]int)
	}
	s[student][subject] = score
}

// SortedKeys returns sorted student names.
func (s StudentScores) SortedKeys() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ToJSON returns a JSON representation of the map.
func (s StudentScores) ToJSON() (string, error) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

