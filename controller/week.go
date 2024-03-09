package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func daysInMonth(m int, year int) int {

	return time.Date(year, time.Month(m)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func daysOfWeek(m int, d int, year int) time.Weekday {

	return time.Date(year, time.Month(m)+1, d, 0, 0, 0, 0, time.UTC).Weekday()
}

func AddWeek(c echo.Context) error {
	var week e.Week
	var schedule e.Schedule
	var unavailable e.Unavaialble
	var UsherGroup e.UsherGroupKV
	var dayToCheck int
	var days int
	var userUsherGroup e.UserUsherGroup
	var peopleAmount int

	jsonInterface := GetJSONRawBody3(c)

	usherGroups := GetStringArrayDataFromJSONByKey(jsonInterface, "UsherGroup")

	for i := 0; i < len(usherGroups); i++ {
		usherGroup := usherGroups[i]
		json.Unmarshal([]byte(usherGroup), &UsherGroup)
		usherGroupData, err := m.GetUsherGroupById(UsherGroup.Value)

		if err != nil {

		}
		startYear, err := GetIntDataFromJSONByKey(jsonInterface, "range.start.year")
		startMonth, err := GetIntDataFromJSONByKey(jsonInterface, "range.start.month")
		startDay, err := GetIntDataFromJSONByKey(jsonInterface, "range.start.day")
		endDay, err := GetIntDataFromJSONByKey(jsonInterface, "range.end.day")
		endYear, err := GetIntDataFromJSONByKey(jsonInterface, "range.end.year")
		endMonth, err := GetIntDataFromJSONByKey(jsonInterface, "range.end.month")

		for y := startYear; y <= endYear; y++ {
			for i := startMonth; i <= endMonth; i++ {
				if i == startMonth {
					dayToCheck = startDay
				} else {
					dayToCheck = 1
				}
				if i == endMonth {
					days = endDay
				} else {
					days = daysInMonth(i, y)
				}

				for d := dayToCheck; d <= days; d++ {
					peopleAmount = usherGroupData.UsherAmount
					dayofWeekMass := daysOfWeek(i, d, y)
					if usherGroupData.Day == strings.ToLower(dayofWeekMass.String()) {
						fmt.Printf("\n\n\nstart day\n\n\n")
						week.Day = d
						week.Hour = usherGroupData.Hour
						week.Minute = usherGroupData.Minute
						week.Month = i
						week.Year = y
						week.UsherGroup = usherGroupData.ID
						weekId, err := m.AddWeek(&week)
						userUsherGroup.UsherGroup = usherGroupData.ID
						userUsherGroup.Number = 0
						usersInUsherGroup, err := m.GetUserUsherGroupByUsherGroup(userUsherGroup)

						if err != nil {

						}

						if len(usersInUsherGroup) < peopleAmount {
							fmt.Printf("\n\n\nless then user amount\n\n\n")

							peopleAmount = usherGroupData.UsherAmount - len(usersInUsherGroup)

							unavailable.UsherGroup = usherGroupData.ID
							m.RemoveUnAvailable(&unavailable)

							for last := 0; last < len(usersInUsherGroup); last++ {
								schedule.UserUsherGroup = usersInUsherGroup[last].ID
								schedule.Week = weekId
								m.AddSchedule(&schedule)
								unavailable.UserUsherGroup = usersInUsherGroup[last].ID
								unavailable.UsherGroup = usherGroupData.ID
								m.AddUnAvailable(&unavailable)
							}

							usersInUsherGroup, err = m.GetUserUsherGroupByUsherGroup(userUsherGroup)
							if err != nil {

							}
							/*
								unavailable.UsherGroup = usherGroupData.ID
								m.RemoveUnAvailable(&unavailable)
								number := uint64(usherGroupData.UsherAmount - len(usersInUsherGroup))
								fmt.Printf("\n\n\nthe number: %d", number)
								userUsherGroup.Number = number
								userUsherGroup.UsherGroup = usherGroupData.ID
								usersInUsherGroupTemp, err := m.GetUserUsherGroupByUsherGroup(userUsherGroup)
								if err != nil {
									panic(err)
								}
								usersInUsherGroup = append(usersInUsherGroup, usersInUsherGroupTemp...)
								//							fmt.Printf("made it here")

								//							os.Exit(3)
							*/
						}

						for people := 0; people < peopleAmount; people++ {
							fmt.Printf("\n\n\nstart people\n\n\n")

							fmt.Printf("the list %v", usersInUsherGroup)
							rand.Seed(time.Now().Unix())
							n := rand.Intn(len(usersInUsherGroup))
							schedule.UserUsherGroup = usersInUsherGroup[n].ID
							schedule.Week = weekId
							err := m.AddSchedule(&schedule)
							unavailable.UserUsherGroup = usersInUsherGroup[n].ID
							unavailable.UsherGroup = usherGroupData.ID

							m.AddUnAvailable(&unavailable)
							if err != nil {
								fmt.Printf("error add achedule: %s:", err)
							}

							fmt.Printf("\n\n\n before slice: %v\n\n\n", usersInUsherGroup)
							//	slice := []int{1, 2, 3, 4}
							fmt.Printf("\n\n\nthe number: %d\n\n", n)
							fmt.Printf("\n\n\nthe number remaining: %d\n\n", len(usersInUsherGroup))
							usersInUsherGroup, err = remove(usersInUsherGroup, n)

							fmt.Printf("\n\n\n after slice: %v\n\n\n", usersInUsherGroup)

							//	fmt.Println(slice) // [1 3 4]
							//	RemoveIndex(usersInUsherGroup, n)
							fmt.Printf("\n\n\n end people\n\n\n")
						}
					}
				}
				fmt.Printf("made it here")
				os.Exit(3)

			}
		}
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//return c.JSON(http.StatusOK, gallery)
}

func remove(s []e.UserUsherGroup, index int) ([]e.UserUsherGroup, error) {
	if index >= len(s) {
		return nil, errors.New("Out of Range Error")
	}
	return append(s[:index], s[index+1:]...), nil
}

func GetWeeks(c echo.Context) error {
	gallery, err := m.GetMainGallery()

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	}

	return c.JSON(http.StatusOK, gallery)
}

func GetWeek(c echo.Context) error {
	category := c.Param("category")

	gallery, err := m.GetGalleryByCategory(category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find category")
	}

	return c.JSON(http.StatusOK, gallery)
}
