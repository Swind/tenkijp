package resource

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const DOMAIN = "http://www.tenki.jp"

type Address struct {
	Url string `json:"url"`
}

func (address *Address) ReadHTML() (*goquery.Document, error) {
	return goquery.NewDocument(DOMAIN + address.Url)
}

type Country struct {
	Name   string `json:"name"`
	Cities []City `json:"cities"`
	Address
}

type City struct {
	Name  string `json:"name"`
	Areas []Area `json:"areas"`
	Address
}

func (city *City) Id() int {
	s := strings.Split(city.Url, "/")

	id1, _ := strconv.Atoi(s[len(s)-1])
	id2, _ := strconv.Atoi(s[len(s)-2])

	return id1*100 + id2
}

type Area struct {
	Name  string `json:"name"`
	Towns []Town `json:"towns"`
	Address
}

func (area *Area) DressUrl() string {
	return strings.Replace(area.Url, "forecast", "indexes/dress", -1)
}

func (area *Area) ReadDressHTML() (*goquery.Document, error) {
	return goquery.NewDocument(area.DressUrl())
}

func (area *Area) Id() int {
	s := strings.Split(area.Url, "/")

	id_str := strings.Replace(s[len(s)-1], ".html", "", -1)
	id, _ := strconv.Atoi(id_str)

	return id
}

type Town struct {
	Name string `json:"name"`
	Address
}

func (town *Town) Id() int {
	s := strings.Split(town.Url, "/")

	id_str := strings.Replace(s[len(s)-1], ".html", "", -1)
	id_str = strings.Replace(id_str, "_", "", -1)
	id, _ := strconv.Atoi(id_str)

	return id
}

type DressIndex struct {
	Date      time.Time `json:"date"`
	ToIndex   int       `json:"today_index"`
	ToAdvice  string    `json:"today_advice"`
	TmrIndex  int       `json:"tomorrow_index"`
	TmrAdvice string    `json:"tomorrow_advice"`
}
