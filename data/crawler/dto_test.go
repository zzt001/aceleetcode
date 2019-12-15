package crawler_test

import (
	"reflect"
	"testing"

	"github.com/zzt001/aceleetcode/data/crawler"
	"github.com/zzt001/aceleetcode/data/entity"
)

import "github.com/stretchr/testify/require"

var testTable = []struct {
	in  crawler.Dto
	out []*entity.Problem
}{
	{
		in: crawler.Dto{
			Pairs: []crawler.Pair{
				{crawler.Stat{
					QuestionId: 1,
					Title:      "problem1",
					Slug:       "problme-1",
				},
					crawler.Difficulty{
						Level: 1,
					},
				},
				{crawler.Stat{
					QuestionId: 2,
					Title:      "problem2",
					Slug:       "problme-2",
				},
					crawler.Difficulty{
						Level: 2,
					},
				},
				{crawler.Stat{
					QuestionId: 3,
					Title:      "problem3",
					Slug:       "problme-3",
				},
					crawler.Difficulty{
						Level: 3,
					},
				},
			},
		},
		out: []*entity.Problem{
			{
				ID:    1,
				Lv:    entity.Easy,
				Title: "problem1",
				Url:   crawler.BASE_URL + "problme-1",
			},
			{
				ID:    2,
				Lv:    entity.Medium,
				Title: "problem2",
				Url:   crawler.BASE_URL + "problme-2",
			},
			{
				ID:    3,
				Lv:    entity.Hard,
				Title: "problem3",
				Url:   crawler.BASE_URL + "problme-3",
			},
		},
	},
}

func TestDtoToProblems(t *testing.T) {
	ast := require.New(t)
	for _, tt := range testTable {
		pbs, err := tt.in.ToProblems()
		ast.Nil(err)
		ast.Equal(len(pbs), len(tt.out))
		for i, _ := range pbs {
			ast.Equal(*tt.out[i], *pbs[i])
		}
		ast.True(reflect.DeepEqual(tt.out, pbs))
	}
}
