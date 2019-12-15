package crawler

import (
	"errors"
	"fmt"
	"net/url"
	"path"

	"github.com/zzt001/aceleetcode/data/entity"
)

type Dto struct {
	Pairs []Pair `json:"stat_status_pairs"`
}

type Pair struct {
	Stat       `json:"stat"`
	Difficulty `json:"difficulty"`
}

type Difficulty struct {
	Level int `json:"level"`
}

type Stat struct {
	QuestionId int    `json:"question_id"`
	Title      string `json:"question__title"`
	Slug       string `json:"question__title_slug"`
}

func (dto *Dto) ToProblems() (pbs []*entity.Problem, err error) {
	if dto == nil {
		return nil, errors.New("dto should not be nil")
	}
	for _, pair := range dto.Pairs {
		link, err := problemUrl(pair.Slug)
		if err != nil {
			fmt.Printf("could not get entity link of %v, will set to empty\n", pair)
		}
		pb, err := entity.NewProblem(pair.QuestionId, pair.Title, link, pair.Level)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pbs = append(pbs, pb)
	}
	return pbs, nil
}

func problemUrl(slug string) (string, error) {
	u, err := url.Parse(BASE_URL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, slug)
	return u.String(), nil
}
