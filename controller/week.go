package controller

import (
	"fmt"
	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"
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

	//var week e.CreateWeek
	//fmt.Println("start of addweek")
	//	var week e.CreateWeek
	//	err := c.Bind(&week)
	jsonInterface := GetJSONRawBody3(c)
	//	fmt.Printf("the stringigy: %s\n", test)

	//	jsonBytes, err := json.Marshal(jsonInterface)
	//	if err != nil {

	//	}
	//var monthDays = 0
	//	boot := json.Unmarshal(test, &e.CreateWeek)
	//fmt.Printf("the helloa: %s\n", jsonBytes)
	/*endDay := gjson.GetBytes(test2, "end.day")
	endMonth := gjson.GetBytes(test2, "end.month")
	endYear := gjson.GetBytes(test2, "end.year")
	startDay := gjson.GetBytes(test2, "start.day") */
	massTimeDay, err := GetDataFromJSONByKey(jsonInterface, "massTime.day")
	massTimeMonth, err := GetDataFromJSONByKey(jsonInterface, "massTime.month")
	massTimeYear, err := GetDataFromJSONByKey(jsonInterface, "massTime.year")
	massTimeHour, err := GetDataFromJSONByKey(jsonInterface, "massTime.hour")
	massTimeMinute, err := GetDataFromJSONByKey(jsonInterface, "massTime.minute")
	startYear, err := GetDataFromJSONByKey(jsonInterface, "start.year")
	startMonth, err := GetDataFromJSONByKey(jsonInterface, "start.month")

	//endDay, err := GetDataFromJSONByKey(jsonInterface, "end.day")
	startDay, err := GetDataFromJSONByKey(jsonInterface, "start.day")
	endDay, err := GetDataFromJSONByKey(jsonInterface, "end.day")

	endMonth, err := GetDataFromJSONByKey(jsonInterface, "end.month")

	//	startMonth := gjson.GetBytes(jsonBytes, "start.month")
	//	endMonth := gjson.GetBytes(jsonBytes, "end.month")
	/*startYear := gjson.GetBytes(test2, "start.year")
	massTimeDay := gjson.GetBytes(test2, "massTime.day")
	massTimeHour := gjson.GetBytes(test2, "massTime.hour")
	massTimeMinute := gjson.GetBytes(test2, "massTime.minute")
	massTimeMonth := gjson.GetBytes(test2, "massTime.month")
	massTimeYear := gjson.GetBytes(test2, "massTime.year") */
	//strconv.ParseFloat(startMonth, 10)
	/*raw := []byte(startMonth.Raw)
	fmt.Printf("booh: %s\n", raw)
	too, err := strconv.Atoi(string([]byte(startMonth.Raw)))
	fmt.Printf("booh: %d\n", too)
	fmt.Printf("json data start month: %s\n", startMonth.Num)
	fmt.Printf("json data start month: %s\n", startMonth.Type) */

	//dayOfWeekMassTime := daysOfWeek(massTimeMonth, massTimeDay, massTimeYear)
	//fmt.Printf("day of week: %s\n", dayOfWeekMassTime)

	daytoCheck := startDay
	var dayEndToCheck int = endDay

	for i := startMonth; i <= endMonth; i++ {
		if (daytoCheck != startDay) && (startMonth == massTimeMonth) {
			fmt.Printf("day going reset start\n")
			dayOfWeekMassTime := daysOfWeek(massTimeMonth, massTimeDay, massTimeYear)
			fmt.Printf("day of the week is : %s\n", dayOfWeekMassTime)

			daytoCheck = 1
			for b := 1; b <= 8; b++ {
				fmt.Printf("day going reset: %d\n", b)
				checkDay := daysOfWeek(startMonth, b, massTimeYear)
				fmt.Printf("day of the week checkday : %s\n", checkDay)
				if dayOfWeekMassTime == checkDay {
					daytoCheck = b
				}
			}

		}
		fmt.Printf("start of for month loop %d\n", i)
		for dateLoop := true; dateLoop; {
			fmt.Printf("begin of loop\n")
			//			fmt.Printf("start month: %d\n", startMonth)//
			//			fmt.Printf("massTimeMonth: %d\n", startMonth)//
			//		fmt.Printf("startYear: %d\n", startYear)
			//		fmt.Printf("massTimeYear: %d\n", massTimeYear)
			fmt.Printf("massTimeDay: %d\n", massTimeDay)

			fmt.Printf("startDay: %d\n", daytoCheck)
			fmt.Printf("endDay: %d\n", endDay)
			//	fmt.Printf("tart month: %s\n", startMonth)
			//	fmt.Printf("json data at string: %s\n", ApplyMarshal(months))
			//	fmt.Printf("json data start month: %d\n", i)

			//if startMonth == massTimeMonth && startYear == massTimeYear {

			fmt.Printf("next in loop\n")
			var totalDays int = daysInMonth(i, startYear)

			if startMonth < endMonth {
				dayEndToCheck = totalDays
			}
			if daytoCheck > dayEndToCheck && (startMonth == massTimeMonth) {
				fmt.Printf("dateLoop is false\n")
				dateLoop = false
				continue
			}
			fmt.Printf("after first check\n")
			fmt.Printf("massTimeDay: %d\n", massTimeDay)
			fmt.Printf("startDay: %d\n", daytoCheck)
			if (i == massTimeMonth) && (daytoCheck < massTimeDay) {
				fmt.Printf("reset daytocheck next in loop lalalala\n")
				daytoCheck = massTimeDay
			}

			week.Day = daytoCheck
			week.Hour = massTimeHour
			week.Minute = massTimeMinute
			week.Month = i
			week.Year = startYear
			fmt.Printf("befre entering model\n")
			m.AddWeek(&week)
			//	err := c.Bind(&week)

			if daytoCheck+7 > totalDays {
				dateLoop = false
				fmt.Printf("over month %d\n", daytoCheck)
			} else {
				fmt.Printf("before  7 %d\n", daytoCheck)
				daytoCheck = daytoCheck + 7
				fmt.Printf("adding 7 %d\n", daytoCheck)
				fmt.Printf("dateloop 7 %t\n", dateLoop)
			}

			//	} else {
			//		fmt.Printf("wtf is going on\n")
			//	}

			//			startDay

			//monthDays += daydInMonth

		}

		//	fmt.Printf("json data outside days: %d\n", massTimeDay)
		//	fmt.Printf("json data outside days cool: %d\n", startDay)

		//	fmt.Printf("json data type outside days: %s\n", reflect.TypeOf(massTimeDay))
		//	fmt.Printf("json data type outside days cool: %s\n", reflect.TypeOf(startDay))

		//	if startDay < massTimeDay {
		//var daysbetweenStartDayAndMassTime int = int(massTimeDay) - int(startDay)
		//
		//fmt.Printf("json data outside dddddays: %d\n", daysbetweenStartDayAndMassTime)
		//		daydInMonth := daysInMonth(i, startYear)
		///	day := daysIn(startDay, startYear)
		//		//if startDay < massTimeDay {

		//	massTimeDay := startDay
		//
		//		fmt.Printf("json data outside days33: %d\n", daydInMonth)
		//fmt.Printf("json data outside days cool: %d\n", daydInMonth)

		//		for dateLoop := true; dateLoop; dateLoop = !true {

		//monthDays += daydInMonth

	}

	//test, err := GetDataFromJSONByKey(months, ConvertIntToString(i))
	if err != nil {

	}
	//	fmt.Printf("json data: %d\n", test)
	//test, err := json.Marshal(months[i])
	//	monthDays += testing

	//fmt.Printf("json data: %d\n", monthDays)
	//	for i := startDay; i < 5; i++ {
	//		sum +=
	//	}

	//boot := c.Bind(&week)
	//	if err != nil {/
	//		return c.JSON(http.StatusBadRequest, "could not bind stuff")
	//	}

	//test5 := json.Unmarshal(test2, &week)

	//	fmt.Printf("json data: %s\n", end)

	//	jsonData := GetJSONRawBody(c)

	//challenge := json_map["end"]

	//jsonData, err := json.Marshal(c)
	//	fmt.Printf("json data: %s\n", jsonData)
	//fmt.Println(c.Get("start"))
	//var week e.Week
	//	err := c.Bind(&week)

	///if err != nil {
	//		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	//	}
	//fmt.Println(week)
	///	err = m.AddWeek(&week)

	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	//	}

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
