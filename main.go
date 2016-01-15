package main

import (
	"fmt"
	. "swind/tenkijp/utils"
	"time"
)

func Task(c chan AreaData, area Area) {
	dress_doc, _ := area.ReadDressHTML()
	index := GetAreaDressIndex(area, dress_doc)

	temp_doc, _ := area.ReadHTML()
	temp := GetAreaTemperature(area, temp_doc)

	data := AreaData{time.Now(), area, index, temp}

	c <- data
}

func GetTodayData() []AreaData {
	// Load all urls
	country, _ := Load("./country.json")
	data_list := []AreaData{}

	c := make(chan AreaData)

	// Parse area data and save to datastore
	area_count := 0
	for _, city := range country.Cities {
		for _, area := range city.Areas {
			go Task(c, area)
			area_count++
		}
	}

	// Wait all task finished
	for i := 0; i < area_count; i++ {
		data_list = append(data_list, <-c)
	}

	return data_list
}

func main() {
	data_list := GetTodayData()
	fmt.Println(data_list)
}
