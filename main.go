package main

import (
	"encoding/json"
	"net/http"
	"swind/tenkijp/mydrive"
	"swind/tenkijp/tenkijp"
	"time"
)

func ParseTenkiJPData() {
	parser := tenkijp.TenkiJP{Client: &http.Client{}}

	data_list := parser.GetTodayData()

	day := time.Now()
	const layout = "2006-01-02"
	SaveToGoogleDrive(data_list, day.Format(layout)+".json")
}

func SaveToGoogleDrive(item interface{}, filename string) error {
	parents := []string{"0B0pY6eBbGeRad01xUkZ5VzR4cGM"}

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}

	svr, err := mydrive.GetMyService()
	mydrive.UploadWithParents(svr, parents, filename, string(b))
	return err
}

func main() {
	ParseTenkiJPData()
}
