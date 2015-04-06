package twse

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/toomore/gogrs/utils"
)

func assertType(t *testing.T, t1 interface{}, t2 interface{}) {
	if reflect.TypeOf(t1) != reflect.TypeOf(t2) {
		t.Errorf("Diff type t1(%s), t2(%s)", reflect.TypeOf(t1), reflect.TypeOf(t2))
	}
}

func TestURL(t *testing.T) {
	d := &DailyData{
		No:   "2618",
		Date: time.Now(),
	}
	assertType(t, d, &DailyData{})
}

func ExampleDailyData() {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(d)

	stockData, _ := d.Get()
	fmt.Println(stockData)
}

func TestDailyData_Get(*testing.T) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(d)

	stockData, _ := d.Get()
	d.Get() // Test Cache.
	fmt.Println(stockData)
}

func TestGetList(*testing.T) {
	var stock = &DailyData{
		No:   "2329",
		Date: time.Date(2015, 03, 20, 0, 0, 0, 0, time.Local),
	}

	stock.Get()
	fmt.Println(stock.RawData)
	fmt.Println(stock.MA(6))
	fmt.Println(stock.MAV(6))
	fmt.Println(stock.GetPriceList())
	fmt.Println(utils.ThanPastFloat64(stock.GetPriceList(), 3, true))
	fmt.Println(utils.ThanPastFloat64(stock.GetPriceList(), 3, false))
	fmt.Println(stock.GetVolumeList())
	fmt.Println(utils.ThanPastUint64(stock.GetVolumeList(), 3, true))
	fmt.Println(utils.ThanPastUint64(stock.GetVolumeList(), 3, false))
	fmt.Println(stock.GetRangeList())
	fmt.Println(stock.GetOpenList())
	fmt.Println(stock.IsRed())
}

func BenchmarkGet(b *testing.B) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	for i := 0; i <= b.N; i++ {
		d.Get()
	}
}

func BenchmarkGetVolumeList(b *testing.B) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local),
	}
	d.Get()
	for i := 0; i <= b.N; i++ {
		d.GetVolumeList()
	}
}

func BenchmarkGetPriceList(b *testing.B) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local),
	}
	d.Get()
	for i := 0; i <= b.N; i++ {
		d.GetPriceList()
	}
}

func ExampleDailyData_Get_notEnoughData() {
	year, month, _ := time.Now().Date()
	d := &DailyData{
		No:   "2618",
		Date: time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local),
	}

	stockData, err := d.Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stockData)
	}
}

func ExampleDailyData_Round() {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(d)

	stockData, _ := d.Get()
	fmt.Println(stockData)

	fmt.Println(d.URL()) // 2014/12

	d.Round()
	fmt.Println(d.URL()) // 2014/11
	stockData, _ = d.Get()
	fmt.Println(stockData)

	d.Round()
	fmt.Println(d.URL()) // 2014/10
	stockData, _ = d.Get()
	fmt.Println(stockData)
}

func TestDailyData_Round(*testing.T) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	d.Round()
	fmt.Println(d.URL())
}

func TestDailyData_GetByTimeMap(*testing.T) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(d.GetByTimeMap())
}

func TestDailyData_FormatDailyData(*testing.T) {
	d := &DailyData{
		No:   "2618",
		Date: time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
	}
	d.Get()
	fmt.Println(d.FormatDailyData())
}
