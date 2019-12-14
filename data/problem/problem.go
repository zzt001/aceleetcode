package problem

type Level int

const (
	Easy Level = iota + 1
	Medium
	Hard
)

type Problem struct {
	ID string
	Lv Level
	Url string
	Tag string
	Title string
}

func NewProblem(lv Level, url, tag, title string) *Problem {
	return &Problem{Lv: lv, Url: url, Tag: tag, Title: title}
}

func (l Level) String() string {
	return []string{"Unknown", "Easy", "Medium", "Hard"}[l]
}
