package entity

type User struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Description string `db:"description"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"emupdated_atail"`
}
