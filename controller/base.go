package controller

import (
	"encoding/json"
	"fmt"
	"io"
	e "scheduler-api/entity"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
)

var (
	EmptyValue = make([]int, 0)
)

func GetJSONRawBody3(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	return jsonBody
}

func ApplyMarshal(jsonInterface map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(jsonInterface)
	if err != nil {
		return nil
	}

	return jsonBytes
}

func ConvertStringToInt(stringData gjson.Result) (int, error) {
	//	fmt.Printf("json stirng data marsh string hwerw: %s\n", stringData)
	//	fmt.Printf("json stirng data marsh string raw: %s\n", stringData.Raw)
	intData, err := strconv.Atoi(string([]byte(stringData.Raw)))
	//	intData, err := strconv.Atoi(stringData)
	//intNumberMonth, err := strconv.Atoi(startMonth)
	//	fmt.Printf("json stirng data marsh int: %d\n", intData)
	return intData, err
}

func ConvertStructToJSON(bytesData []byte) map[string]interface{} {
	var jsonMap map[string]interface{}

	stringData := string(bytesData[:])

	stringDataReady := fmt.Sprintf("{\"data\": %s}", stringData)

	json.Unmarshal([]byte(stringDataReady), &jsonMap)

	return jsonMap

}

func ConvertIntToString(number int) string {
	intNumberMonth := strconv.Itoa(number)

	return intNumberMonth
}

func GetDataFromJSONByKey(jsonInterface map[string]interface{}, key string) (int, error) {
	//	fmt.Printf("json stirng data interface: %s\n", jsonInterface)
	//	fmt.Printf("json stirng key: %s\n", key)

	jsonBytes := ApplyMarshal(jsonInterface)
	//	fmt.Printf("json stirng data marsh: %s\n", jsonBytes)

	valueBytes := gjson.GetBytes(jsonBytes, key)
	//	fmt.Printf("number starting month: %s\n", valueBytes)
	valueInt, err := ConvertStringToInt(valueBytes)
	//	fmt.Printf("json stirng data marsh int: %d\n", valueInt)
	return valueInt, err
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	//fmt.Printf("error cehck: %s\n", err)
	if err != nil {

		//fmt.Printf("not nil jsonDatajson data: %s\n", err)
		//log.Error("empty json body")
		return nil
	} else {
		//	fmt.Printf("nil here jsonDatajson dataaaaa: %s\n", err)
		//	fmt.Printf("end")
		return jsonBody
	}
}

func GetJSONRawBody2(c echo.Context) {
	fmt.Println("start of addweek2")
	my_data := echo.Map{}

	if err := c.Bind(&my_data); err != nil {
		//	return err
	} else {

		start := fmt.Sprintf("%v", my_data["start"])
		fmt.Printf("jsonDatajson data: %s\n", start)
		//		useremail := fmt.Sprintf("%v", my_data["end"])
	}

	//	fmt.Printf("jsonDatajson data: %s\n", start)
	/*	var week e.CreateWeek
		//jsonBody := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&week)
		if err != nil {

			log.Error("empty json body")
			return nil
		}
		fmt.Printf("jsonDatajson data: %s\n", week)
		//fmt.Printf("end") */
	//return err
}

func UnwantedJSONHandler(c echo.Context) error {
	b, _ := io.ReadAll(c.Request().Body)
	fmt.Printf("dispay b: %v\n", b)
	var week e.CreateWeek
	answer := json.Unmarshal(b, &week)

	if answer != nil {
		fmt.Printf("jsonData to display: %s\n", &answer)
		return answer
	}
	//	log.Println(err.Error())
	return answer
}

//	fmt.Printf("jsonDatajson data: %s\n", week)

//	return c.JSON(week)
//	return week
//
// return week
// log.Println(r)
