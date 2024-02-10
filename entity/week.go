package entity

type Week struct {
	ID     string `db:"id"`
	Hour   string `db:"hour"`
	Minute string `db:"minute"`
	Day    string `db:"day"`
	Month  string `db:"month"`
	Year   string `db:"year"`
}

type Analytics struct {
	Visitors  int `json:"visitors"`
	PageViews int `json:"page_views"`
}

type CreateWeek struct {
	Begin  []Start
	Finish []End
	Mass   []MassTime
}

// Error implements error.
func (CreateWeek) Error() string {
	panic("unimplemented")
}

type Start struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type End struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
type MassTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Day    int `json:"day"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}
