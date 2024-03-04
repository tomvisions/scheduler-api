package entity

type UsherGroup struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Day         string `db:"day"`
	Hour        int    `db:"hour"`
	Minute      int    `db:"minute"`
}

type UsherGroupKV struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
