package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	//var jsonMap map[string]interface{}
	var UsherGroup e.UsherGroupKV
	var dayToCheck int
	var days int

	jsonInterface := GetJSONRawBody3(c)

	usherGroups := GetStringArrayDataFromJSONByKey(jsonInterface, "UsherGroup")

	for i := 0; i < len(usherGroups); i++ {
		usherGroup := usherGroups[i]
		json.Unmarshal([]byte(usherGroup), &UsherGroup)
		usherGroupData, err := m.GetUsherGroupById(UsherGroup.Value)

		if err != nil {

		}
		//	massTimeHour, err := GetIntDataFromJSONByKey(jsonInterface, "massTime.hour")
		//	massTimeMinute, err := GetIntDataFromJSONByKey(jsonInterface, "massTime.minute")
		startYear, err := GetIntDataFromJSONByKey(jsonInterface, "range.start.year")
		startMonth, err := GetIntDataFromJSONByKey(jsonInterface, "range.start.month")
		startDay, err := GetIntDataFromJSONByKey(jsonInterface, "range.start.day")
		endDay, err := GetIntDataFromJSONByKey(jsonInterface, "range.end.day")
		endYear, err := GetIntDataFromJSONByKey(jsonInterface, "range.end.year")
		endMonth, err := GetIntDataFromJSONByKey(jsonInterface, "range.end.month")

		fmt.Printf(("about to start loop"))
		for y := startYear; y <= endYear; y++ {
			fmt.Printf(("about to start loop year"))
			for i := startMonth; i <= endMonth; i++ {
				fmt.Printf(("about to start loop month"))
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
					//		fmt.Printf("about to start loop day 1: %s\n", usherGroupData.Day)
					dayofWeekMass := daysOfWeek(i, d, y)
					//		fmt.Printf("about to start loop day 2: %s\n", strings.ToLower(dayofWeekMass.String()))
					if usherGroupData.Day == strings.ToLower(dayofWeekMass.String()) {
						//	fmt.Printf("I am IN\n\n\n")
						week.Day = d
						week.Hour = usherGroupData.Hour
						week.Minute = usherGroupData.Minute
						week.Month = i
						week.Year = y
						week.UsherGroup = usherGroupData.ID
						m.AddWeek(&week)

					}
				}
			}
		}
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//return c.JSON(http.StatusOK, gallery)
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
