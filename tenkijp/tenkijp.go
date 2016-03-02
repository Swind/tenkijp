package tenkijp

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type TenkiJP struct {
	Client *http.Client
}

func (self *TenkiJP) GetDocument(url string) (*goquery.Document, error) {
	res, e := self.Client.Get(url)
	if e != nil {
		return nil, e
	}

	return goquery.NewDocumentFromResponse(res)
}

func (self *TenkiJP) Task(c chan AreaData, area Area) {
	dress_doc, _ := self.GetDocument(area.GetFullPath())
	index := GetAreaDressIndex(area, dress_doc)

	temp_doc, _ := self.GetDocument(area.GetFullDressPath())
	temp := GetAreaTemperature(area, temp_doc)

	data := AreaData{time.Now(), area, index, temp}

	c <- data
}

func (self *TenkiJP) GetTodayData() []AreaData {
	// Load all urls
	country, _ := Load("./country.json")
	data_list := []AreaData{}

	c := make(chan AreaData)

	// Parse area data and save to datastore
	area_count := 0
	for _, city := range country.Cities {
		for _, area := range city.Areas {
			go self.Task(c, area)
			area_count++
		}
	}

	// Wait all task finished
	for i := 0; i < area_count; i++ {
		data_list = append(data_list, <-c)
	}

	return data_list
}
