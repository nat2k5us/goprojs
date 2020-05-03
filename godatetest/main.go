package main

import (
	"flag"
	"fmt"
	"github.com/araddon/dateparse"
	"time"
)

var (
	timezone = ""
)

func main() {
	dateExample := "11/11/2019 01:00:01 PM"
	ret := GetRFCFormatDate(dateExample)
	fmt.Println("converted: ",ret)
}
// GetRFCFormatDate get the date
func GetRFCFormatDate(d string) (string){
	flag.StringVar(&timezone, "timezone", "UTC", "Timezone aka `America/Los_Angeles` formatted time-zone")
	flag.Parse()

	if timezone != "" {
		// NOTE:  This is very, very important to understand
		// time-parsing in go
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			panic(err.Error())
		}
		time.Local = loc
	}
	
	t, err := dateparse.ParseLocal(d)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("original: ",t)
	v := t.Format(time.RFC3339)
	
	return v
}
