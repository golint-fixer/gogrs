package main

import (
	"fmt"
	"time"
)

type DailyData struct {
	stock_no string
	date     time.Time
}

func (d DailyData) Url() (url string) {
	url = fmt.Sprintf(TWSECSV, d.date.Year(), d.date.Month(), d.stock_no, d.date.UnixNano())
	return
}

func main() {
	d := &DailyData{
		stock_no: "2618",
		date:     time.Now(),
	}

	fmt.Println(d.Url())

}
