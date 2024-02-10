package controller

import (
	"fmt"
	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	"time"

	"github.com/labstack/echo/v4"
)

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func AddWeek(c echo.Context) error {

	months := map[string]interface{}{
		"1":  31,
		"2":  28,
		"3":  31,
		"4":  30,
		"5":  31,
		"6":  30,
		"7":  31,
		"8":  31,
		"9":  30,
		"10": 31,
		"11": 30,
		"12": 31,
	}

	//var week e.CreateWeek
	fmt.Println("start of addweek")
	//	var week e.CreateWeek
	//	err := c.Bind(&week)
	jsonInterface := GetJSONRawBody3(c)
	//	fmt.Printf("the stringigy: %s\n", test)

	//	jsonBytes, err := json.Marshal(jsonInterface)
	//	if err != nil {

	//	}
	var monthDays = 0
	//	boot := json.Unmarshal(test, &e.CreateWeek)
	//fmt.Printf("the helloa: %s\n", jsonBytes)
	/*endDay := gjson.GetBytes(test2, "end.day")
	endMonth := gjson.GetBytes(test2, "end.month")
	endYear := gjson.GetBytes(test2, "end.year")
	startDay := gjson.GetBytes(test2, "start.day") */

	startMonth, err := GetDataFromJSONByKey(jsonInterface, "start.month")

	if err != nil {

	}
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
	for i := startMonth; i <= endMonth; i++ {
		fmt.Printf("tart month: %s\n", startMonth)
		fmt.Printf("json data at string: %s\n", ApplyMarshal(months))
		fmt.Printf("json data start month: %d\n", i)

		test, err := GetDataFromJSONByKey(months, ConvertIntToString(i))
		if err != nil {

		}
		fmt.Printf("json data: %d\n", test)
		//test, err := json.Marshal(months[i])
		monthDays += test
	}
	fmt.Printf("json data: %d\n", monthDays)
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
