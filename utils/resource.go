package tenkijp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const DOMAIN = "http://www.tenki.jp"

type AreaData struct {
	Date        time.Time   `json:"date"`
	Area        Area        `json:"area"`
	Index       DressIndex  `json:"index"`
	Temperature Temperature `json:"temperature"`
}

/* ==========================================
 *
 * Address
 *
 * ===========================================*/

type Address struct {
	Url string `json:"url"`
}

func (address *Address) ReadHTML() (*goquery.Document, error) {
	return goquery.NewDocument(DOMAIN + address.Url)
}

func (address *Address) GetFullPath() string {
	return DOMAIN + address.Url
}

/* ==========================================
 *
 * Country
 *
 * ===========================================*/
type Country struct {
	Name   string `json:"name"`
	Cities []City `json:"cities"`
	Address
}

/* ==========================================
 *
 *  City
 *
 * ===========================================*/
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

/* ==========================================
 *
 *  Area
 *
 * ===========================================*/
type Area struct {
	Name string `json:"name"`
	//Towns []Town `json:"towns"`
	Address
}

func (area *Area) DressUrl() string {
	return strings.Replace(area.Url, "forecast", "indexes/dress", -1)
}

func (area *Area) ReadDressHTML() (*goquery.Document, error) {
	return goquery.NewDocument(DOMAIN + area.DressUrl())
}

func (area *Area) GetFullDressPath() string {
	return DOMAIN + area.DressUrl()
}

func (area *Area) Id() int {
	s := strings.Split(area.Url, "/")

	id_str := strings.Replace(s[len(s)-1], ".html", "", -1)
	id, _ := strconv.Atoi(id_str)

	return id
}

/* ==========================================
 *
 *  Town
 *
 * ===========================================*/
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

/* ==========================================
 *
 *  TenkiJP data
 *
 * ===========================================*/
type DressIndex struct {
	ToIndex  int    `json:"today_index"`
	ToAdvice string `json:"today_advice"`

	TmrIndex  int    `json:"tomorrow_index"`
	TmrAdvice string `json:"tomorrow_advice"`
}

type Temperature struct {
	ToHighTemp int `json:"today_high_temperature"`
	ToLowTemp  int `json:"today_low_temperature"`

	TmrHighTemp int `json:"tomorrow_high_temperature"`
	TmrLowTemp  int `json:"tomorrow_low_temperature"`
}

/* ==========================================
 *
 *  Country load and save
 *
 * ===========================================*/
func Load(path string) (Country, error) {
	country := Country{}

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

func Save(country Country, path string) error {
	b, err := json.Marshal(country)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, b, 677)
	return err
}
