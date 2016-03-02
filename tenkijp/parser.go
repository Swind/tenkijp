package tenkijp

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func get_cities(url string) []City {
	results := []City{}

	doc, err := goquery.NewDocument(url)

	if err != nil {
		fmt.Println(err)
	} else {
		doc.Find("ul.localList>li>a").Each(func(index int, s *goquery.Selection) {
			city := City{}
			city.Name = s.Text()

			city.Url, _ = s.Attr("href")

			results = append(results, city)
		})
	}

	return results
}

func get_areas_and_towns(city City) []Area {
	results := []Area{}

	doc, err := city.ReadHTML()

	if err != nil {
		fmt.Println(err)
	} else {
		doc.Find("div.contentsBox>div.wrap>h3.subTitle").Each(func(index int, s *goquery.Selection) {
			area_selector := s.Find("span.city_name>a")
			area := Area{}
			area.Name = area_selector.Text()
			area.Url, _ = area_selector.Attr("href")

			towns_selector := s.Next()
			towns_selector.Find("li>a").Each(func(index int, town_selector *goquery.Selection) {
				town := Town{}
				town.Name = town_selector.Text()
				town.Url, _ = town_selector.Attr("href")
				//area.Towns = append(area.Towns, town)
			})

			results = append(results, area)
		})
	}

	return results
}

func map_cities(cities []City, f func(City) City) []City {
	results := make([]City, len(cities))

	for i, elem := range cities {
		results[i] = f(elem)
	}

	return results
}

func GetCountry() Country {
	country := Country{}
	country.Name = "Japan"
	country.Url = DOMAIN

	country.Cities = get_cities(country.Url)
	country.Cities = map_cities(country.Cities, func(city City) City {
		city.Areas = get_areas_and_towns(city)

		return city
	})

	return country
}

func GetAreaDressIndex(area Area, doc *goquery.Document) DressIndex {
	index := DressIndex{}

	index.ToIndex, _ = strconv.Atoi(doc.Find("dl#exponentLargeLeft>dd>dl>dd").First().Text())
	index.ToAdvice = doc.Find("dl#exponentLargeLeft>dd>p").Last().Text()

	index.TmrIndex, _ = strconv.Atoi(doc.Find("dl#exponentLargeRight>dd>dl>dd").First().Text())
	index.TmrAdvice = doc.Find("dl#exponentLargeRight>dd>p").Last().Text()

	return index
}

func GetAreaTemperature(area Area, doc *goquery.Document) Temperature {
	temp := Temperature{}

	temp.ToHighTemp, _ = strconv.Atoi(doc.Find("div#townLeftOneBox tr.highTemp td.temp span.bold").First().Text())
	temp.ToLowTemp, _ = strconv.Atoi(doc.Find("div#townLeftOneBox tr.lowTemp td.temp span.bold").First().Text())

	temp.TmrHighTemp, _ = strconv.Atoi(doc.Find("div#townRightOneBox tr.highTemp td.temp span.bold").First().Text())
	temp.TmrLowTemp, _ = strconv.Atoi(doc.Find("div#townRightOneBox tr.lowTemp td.temp span.bold").First().Text())

	return temp
}
