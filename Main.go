package main

import (
	"fmt"
	"swind/tenkijp/parser"
	"swind/tenkijp/resource"
	"time"
)

type AreaData struct {
	Date        time.Time
	Area        resource.Area
	Index       resource.DressIndex
	Temperature resource.Temperature
}

func main() {
	country, _ := resource.Load("./country.json")

	data_list := []AreaData{}
	for _, city := range country.Cities {
		for _, area := range city.Areas {

			index := parser.GetAreaDressIndex(area)
			temp := parser.GetAreaTemperature(area)

			data := AreaData{time.Now(), area, index, temp}
			data_list = append(data_list, data)
		}
	}

	fmt.Println(data_list)
}
