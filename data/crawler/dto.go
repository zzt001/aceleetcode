package crawler

type Dto struct {
	Pairs []Pair `json:"stat_status_pairs"`
}

type Pair struct {
	Stat       `json:"stat"`
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
	Freq float64 `json:"frequency"`
}

type Stat struct {
	QuestionId int    `json:"question_id"`
	Title      string `json:"question__title"`
}
