package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"swind/tenkijp/resource"

	"github.com/PuerkitoBio/goquery"
)

func get_cities(url string) []resource.City {
	results := []resource.City{}

	doc, err := goquery.NewDocument(url)

	if err != nil {
		fmt.Println(err)
	} else {
		doc.Find("ul.localList>li>a").Each(func(index int, s *goquery.Selection) {
			city := resource.City{}
			city.Name = s.Text()

			city.Url, _ = s.Attr("href")

			results = append(results, city)
		})
	}

	return results
}

func get_areas_and_towns(city resource.City) []resource.Area {
	results := []resource.Area{}

	doc, err := city.ReadHTML()

	if err != nil {
		fmt.Println(err)
	} else {
		doc.Find("div.contentsBox>div.wrap>h3.subTitle").Each(func(index int, s *goquery.Selection) {
			area_selector := s.Find("span.city_name>a")
			area := resource.Area{}
			area.Name = area_selector.Text()
			area.Url, _ = area_selector.Attr("href")

			towns_selector := s.Next()
			towns_selector.Find("li>a").Each(func(index int, town_selector *goquery.Selection) {
				town := resource.Town{}
				town.Name = town_selector.Text()
				town.Url, _ = town_selector.Attr("href")
				area.Towns = append(area.Towns, town)
			})

			results = append(results, area)
		})
	}

	return results
}

func map_cities(cities []resource.City, f func(resource.City) resource.City) []resource.City {
	results := make([]resource.City, len(cities))

	for i, elem := range cities {
		results[i] = f(elem)
	}

	return results
}

func GetCountry() resource.Country {
	country := resource.Country{}
	country.Name = "Japan"
	country.Url = resource.DOMAIN

	country.Cities = get_cities(country.Url)
	country.Cities = map_cities(country.Cities, func(city resource.City) resource.City {
		city.Areas = get_areas_and_towns(city)

		return city
	})

	return country
}

func main() {
	country := GetCountry()
	b, _ := json.Marshal(country)
	ioutil.WriteFile("./urls.json", b, 0644)
}
