package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"mini-test/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func WeatherReq(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	var response models.WeatherResponse

	link := "http://api.openweathermap.org/data/2.5/forecast?q=Jakarta&appid=b5738d44225ec8bfa816728ad855a994"

	req, err := http.NewRequest("GET", link, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res, errRes := client.Do(req)

	if errRes != nil {
		fmt.Println(res.StatusCode)
		return
	}

	body, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Weather Forecast:")
	//  "dt_txt": "2022-08-30 21:00:00"

	for index, value := range response.List {
		date := strings.Split(value.DtTxt, " ")
		yyyy_mm_dd := strings.Split(date[0], "-")

		dayInt, _ := strconv.Atoi(yyyy_mm_dd[2])
		monthInt, _ := strconv.Atoi(yyyy_mm_dd[1])

		dayStr := time.Date(time.Now().Year(), time.Now().Month(), dayInt, 0, 0, 0, 0, time.UTC).Weekday().String()
		monthstr := time.Date(time.Now().Year(), time.Month(monthInt), 0, 0, 0, 0, 0, time.UTC).Month().String()
		if index%8 == 0 {
			fmt.Printf("%s, %d %s %s: %.2f°C\n", dayStr, dayInt, monthstr, date[0], value.Main.Temp-273.15)
		}
	}
}
