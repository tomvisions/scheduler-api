package model

import (
	"fmt"
	"scheduler-api/db"
	e "scheduler-api/entity"
)

func AddWeek(week *e.Week) error {
	//now := time.Now()
	//week.ID = "sdsdfsdfsdf"
	fmt.Printf("about to insert into db\n")
	const query = `INSERT INTO week (id, hour,minute, day, month, year) VALUES (?, ?, ?, ?, ?, ?)`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d-%d-%d-%d-%d", week.Hour, week.Minute, week.Day, week.Month, week.Year)

	_, err = tx.Exec(query, id, week.Hour, week.Minute, week.Day, week.Month, week.Year)

	if err != nil {
		fmt.Printf("Error in SQL: %s\n", err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
