package main

import (
	"fmt"
	"mini-test/weather"
	"net/http"
)

func main() {
	http.HandleFunc("/weather", weather.WeatherReq)
	fmt.Println("runing at port 5000")
	fmt.Println(http.ListenAndServe(":5000", nil))
}
