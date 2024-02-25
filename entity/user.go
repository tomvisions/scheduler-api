package entity

type User struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Description string `db:"description"`
	Phone       string `db:"phone"`
	UsherGroup  string `db:"usher_group"`
}
