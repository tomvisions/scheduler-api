package model

import (
	"scheduler-api/db"
	e "scheduler-api/entity"
)

func AddWeek(week *e.Week) error {
	//now := time.Now()
	const query = `INSERT INTO week (hour,minute, day, month, year) VALUES ($1, $1, $1, $1, $1)`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, week.Hour, week.Minute, week.Day, week.Month, week.Year)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
