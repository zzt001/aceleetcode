package entity

import (
	"errors"
	"fmt"
)

type Level int

const (
	Easy Level = iota + 1
	Medium
	Hard
)

type Problem struct {
	ID        int
	Lv        Level
	Url       string
	Title     string
	Frequency float64
}

func NewProblem(id int, title, url string, lv int, freq float64) (*Problem, error) {
	if lv < 1 || lv > 3 {
		return nil, errors.New(fmt.Sprintf("invalid level get %v", lv))
	}
	return &Problem{ID: id, Lv: Level(lv), Url: url, Title: title, Frequency: freq}, nil
}

func (l Level) String() string {
	return []string{"Unknown", "Easy", "Medium", "Hard"}[l]
}
