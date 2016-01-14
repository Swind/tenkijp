package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"swind/tenkijp/resource"
	"time"
)

func read_data(path string) (resource.Country, error) {
	country := resource.Country{}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("With error:", err, "when read ", path)
		return country, err
	}

	err = json.Unmarshal(b, &country)
	if err != nil {
		fmt.Println("With error:", err, "when unmarshal JSON")
		return country, err
	}

	return country, nil
}

func get_area_dress_finger(country resource.Country) {
	for _, city := range country.Cities {
		for _, area := range city.Areas {
			doc, _ := area.ReadDressHTML()

			index := resource.DressIndex{}
			index.Date = time.Now()
			index.ToIndex, _ = strconv.Atoi(doc.Find("dl#exponentLargeLeft>dd>dl>dd").First().Text())
			index.ToAdvice = doc.Find("dl#exponentLargeLeft>dd>p").Last().Text()

			index.TmrIndex, _ = strconv.Atoi(doc.Find("dl#exponentLargeRight>dd>dl>dd").First().Text())
			index.TmrAdvice = doc.Find("dl#exponentLargeRight>dd>p").Last().Text()
		}
	}
}

/*
func main() {
	fmt.Println("Start to get the dress finger...")
	country := resource.Country{}
	b, _ := ioutil.ReadFile("./urls.json")
	json.Unmarshal(b, &country)
	fmt.Println(country)
}
*/
